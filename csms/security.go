package csms

import (
	"context"
	"crypto/tls"
	"errors"
	"fmt"
	"net"
	"net/http"
	"strings"
	"sync"
	"time"
	"unicode/utf8"

	"github.com/seanlee0923/ocpp/protocol"
)

type SecurityProfile uint8

const (
	SecurityProfileUnsecured SecurityProfile = iota
	SecurityProfileBasicAuth
	SecurityProfileTLSBasicAuth
	SecurityProfileTLSClientCertificate
)

var (
	ErrAuthenticationFailed      = errors.New("authentication failed")
	ErrAuthenticationUnavailable = errors.New("authentication service unavailable")
)

type BasicCredentials struct {
	Username string
	Password []byte
}

type AuthenticationRequest struct {
	Identity   string
	Version    protocol.Version
	RemoteAddr string
	TLS        *tls.ConnectionState
	Basic      *BasicCredentials
}

type Principal struct {
	ID         string
	Attributes map[string]string
}

type Authenticator interface {
	Authenticate(context.Context, AuthenticationRequest) (Principal, error)
}

type AuthenticatorFunc func(context.Context, AuthenticationRequest) (Principal, error)

func (function AuthenticatorFunc) Authenticate(ctx context.Context, request AuthenticationRequest) (Principal, error) {
	return function(ctx, request)
}

type IdentityValidator func(string) error

type HandshakeAttempt struct {
	Identity   string
	Version    protocol.Version
	RemoteAddr string
}

type HandshakeLimiter interface {
	Allow(context.Context, HandshakeAttempt) bool
}

type HandshakeLimiterFunc func(context.Context, HandshakeAttempt) bool

func (function HandshakeLimiterFunc) Allow(ctx context.Context, attempt HandshakeAttempt) bool {
	return function(ctx, attempt)
}

type ipRateLimitWindow struct {
	count int
	reset time.Time
}

// IPRateLimiter limits WebSocket handshake attempts per remote IP using fixed
// windows. It is safe for concurrent use.
type IPRateLimiter struct {
	mu      sync.Mutex
	limit   int
	window  time.Duration
	clients map[string]ipRateLimitWindow
}

func NewIPRateLimiter(limit int, window time.Duration) (*IPRateLimiter, error) {
	if limit <= 0 || window <= 0 {
		return nil, errors.New("IP rate limit and window must be positive")
	}
	return &IPRateLimiter{limit: limit, window: window, clients: make(map[string]ipRateLimitWindow)}, nil
}

// maxRateLimiterClients hard-caps how many distinct remote hosts
// IPRateLimiter tracks at once. Without a hard cap, an attacker making one
// handshake attempt per distinct source IP (e.g. from a large subnet) can
// grow limiter.clients without bound: the old cleanup only removed entries
// whose window had already elapsed, and every entry from an ongoing,
// rotating-IP attack is by definition still within its window — this is
// the server's own rate limiter becoming an unbounded-memory DoS vector.
const maxRateLimiterClients = 1024

func (limiter *IPRateLimiter) Allow(_ context.Context, attempt HandshakeAttempt) bool {
	host, _, err := net.SplitHostPort(attempt.RemoteAddr)
	if err != nil {
		host = attempt.RemoteAddr
	}
	now := time.Now()
	limiter.mu.Lock()
	defer limiter.mu.Unlock()
	entry, tracked := limiter.clients[host]
	if entry.reset.IsZero() || !now.Before(entry.reset) {
		entry = ipRateLimitWindow{reset: now.Add(limiter.window)}
	}
	entry.count++
	if !tracked && len(limiter.clients) >= maxRateLimiterClients {
		limiter.evictOne(now)
	}
	limiter.clients[host] = entry
	return entry.count <= limiter.limit
}

// evictOne makes room for a new client once limiter.clients is at capacity:
// it deletes the first already-expired entry it finds, or, if none of the
// tracked entries have expired yet, the one closest to expiring.
func (limiter *IPRateLimiter) evictOne(now time.Time) {
	var oldestHost string
	var oldestReset time.Time
	for host, candidate := range limiter.clients {
		if !now.Before(candidate.reset) {
			delete(limiter.clients, host)
			return
		}
		if oldestHost == "" || candidate.reset.Before(oldestReset) {
			oldestHost, oldestReset = host, candidate.reset
		}
	}
	if oldestHost != "" {
		delete(limiter.clients, oldestHost)
	}
}

type SecurityEventType string

const (
	SecurityEventAuthenticationSucceeded SecurityEventType = "AuthenticationSucceeded"
	SecurityEventAuthenticationFailed    SecurityEventType = "AuthenticationFailed"
	SecurityEventHandshakeRateLimited    SecurityEventType = "HandshakeRateLimited"
	SecurityEventDuplicateRejected       SecurityEventType = "DuplicateSessionRejected"
	SecurityEventSessionReplaced         SecurityEventType = "SessionReplaced"
)

type SecurityEvent struct {
	Type       SecurityEventType
	Identity   string
	Version    protocol.Version
	RemoteAddr string
	Reason     string
}

type SecurityEventHandler func(context.Context, SecurityEvent)

type SecurityConfig struct {
	Profile                    SecurityProfile
	Authenticator              Authenticator
	ValidateIdentity           IdentityValidator
	AllowBasicUsernameMismatch bool
	MinTLSVersion              uint16
	HandshakeLimiter           HandshakeLimiter
	OnEvent                    SecurityEventHandler
}

func prepareSecurity(config *SecurityConfig) error {
	if config.Profile > SecurityProfileTLSClientCertificate {
		return fmt.Errorf("unsupported security profile %d", config.Profile)
	}
	if config.ValidateIdentity == nil {
		config.ValidateIdentity = ValidateIdentity
	}
	if config.MinTLSVersion == 0 {
		config.MinTLSVersion = tls.VersionTLS12
	}
	if config.Profile != SecurityProfileUnsecured && config.Authenticator == nil {
		return errors.New("authenticator is required for the configured security profile")
	}
	return nil
}

// ValidateIdentity accepts RFC 3986 unreserved characters and limits an OCPP
// identity to 255 bytes. Applications may replace it in SecurityConfig.
func ValidateIdentity(identity string) error {
	if identity == "" || len(identity) > 255 || !utf8.ValidString(identity) {
		return errors.New("identity must contain between 1 and 255 valid UTF-8 bytes")
	}
	for _, character := range identity {
		if (character >= 'a' && character <= 'z') ||
			(character >= 'A' && character <= 'Z') ||
			(character >= '0' && character <= '9') ||
			strings.ContainsRune("-._~", character) {
			continue
		}
		return fmt.Errorf("identity contains unsupported character %q", character)
	}
	return nil
}

func (s *Server) authenticate(r *http.Request, identity string, version protocol.Version) (Principal, int, error) {
	security := s.config.Security
	profile := security.Profile
	if profile == SecurityProfileTLSBasicAuth || profile == SecurityProfileTLSClientCertificate {
		if r.TLS == nil || r.TLS.Version < security.MinTLSVersion {
			return Principal{}, http.StatusForbidden, errors.New("required TLS policy was not satisfied")
		}
	}
	if profile == SecurityProfileTLSClientCertificate {
		if len(r.TLS.PeerCertificates) == 0 || len(r.TLS.VerifiedChains) == 0 {
			return Principal{}, http.StatusForbidden, errors.New("verified client certificate is required")
		}
	}

	var basic *BasicCredentials
	if username, password, ok := r.BasicAuth(); ok {
		basic = &BasicCredentials{Username: username, Password: []byte(password)}
		defer clear(basic.Password)
	}
	if profile == SecurityProfileBasicAuth || profile == SecurityProfileTLSBasicAuth {
		if basic == nil {
			return Principal{}, http.StatusUnauthorized, ErrAuthenticationFailed
		}
		if !security.AllowBasicUsernameMismatch && basic.Username != identity {
			return Principal{}, http.StatusUnauthorized, ErrAuthenticationFailed
		}
	}

	if security.Authenticator == nil {
		return Principal{ID: identity}, 0, nil
	}
	principal, err := security.Authenticator.Authenticate(r.Context(), AuthenticationRequest{
		Identity: identity, Version: version, RemoteAddr: r.RemoteAddr, TLS: r.TLS, Basic: basic,
	})
	if err != nil {
		if errors.Is(err, ErrAuthenticationUnavailable) {
			return Principal{}, http.StatusServiceUnavailable, err
		}
		if profile == SecurityProfileTLSClientCertificate {
			return Principal{}, http.StatusForbidden, ErrAuthenticationFailed
		}
		return Principal{}, http.StatusUnauthorized, ErrAuthenticationFailed
	}
	if principal.ID == "" {
		return Principal{}, http.StatusUnauthorized, errors.New("authenticator returned an empty principal ID")
	}
	principal.Attributes = cloneAttributes(principal.Attributes)
	return principal, 0, nil
}

func cloneAttributes(attributes map[string]string) map[string]string {
	if attributes == nil {
		return nil
	}
	clone := make(map[string]string, len(attributes))
	for key, value := range attributes {
		clone[key] = value
	}
	return clone
}

func (s *Server) securityEvent(ctx context.Context, event SecurityEvent) {
	if handler := s.config.Security.OnEvent; handler != nil {
		// A diagnostic hook must not be able to take down the protocol
		// server (matching s.log()) — and here specifically, must not skip
		// releasing an identity's reservation: reserveIdentity's duplicate-
		// session-replacement path calls this while holding a reservation
		// token, and an unrecovered panic here would leave that identity
		// permanently rejected with "connection is already pending".
		defer func() { _ = recover() }()
		handler(ctx, event)
	}
}
