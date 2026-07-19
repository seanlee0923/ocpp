package csms

import (
	"context"
	"crypto/tls"
	"crypto/x509"
	"encoding/base64"
	"net/http"
	"net/http/httptest"
	"strings"
	"sync/atomic"
	"testing"
	"time"

	"github.com/gorilla/websocket"
	"github.com/seanlee0923/ocpp/protocol"
)

func TestNewRequiresAuthenticatorForSecuredProfiles(t *testing.T) {
	for _, profile := range []SecurityProfile{
		SecurityProfileBasicAuth,
		SecurityProfileTLSBasicAuth,
		SecurityProfileTLSClientCertificate,
	} {
		if _, err := New(Config{Security: SecurityConfig{Profile: profile}}); err == nil {
			t.Errorf("New succeeded without authenticator for profile %d", profile)
		}
	}
}

func TestBasicAuthenticationStoresPrincipalWithoutCredentials(t *testing.T) {
	server, err := New(Config{
		Versions: []protocol.Version{protocol.OCPP16},
		Security: SecurityConfig{
			Profile: SecurityProfileBasicAuth,
			Authenticator: AuthenticatorFunc(func(_ context.Context, request AuthenticationRequest) (Principal, error) {
				if request.Basic == nil || request.Basic.Username != "CP-001" || string(request.Basic.Password) != "secret" {
					return Principal{}, ErrAuthenticationFailed
				}
				return Principal{ID: "station:CP-001", Attributes: map[string]string{"tenant": "one"}}, nil
			}),
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()
	header := basicHeader("CP-001", "secret")
	conn := dialStationWithHeader(t, httpServer.URL, protocol.OCPP16, header)
	defer conn.Close()
	session := waitForSession(t, server, "CP-001")
	principal := session.Principal()
	if principal.ID != "station:CP-001" || principal.Attributes["tenant"] != "one" {
		t.Fatalf("principal = %#v", principal)
	}
	principal.Attributes["tenant"] = "mutated"
	if session.Principal().Attributes["tenant"] != "one" {
		t.Fatal("session principal attributes were externally mutated")
	}
}

func TestBasicAuthenticationRejectsMissingCredentials(t *testing.T) {
	server, err := New(Config{
		Versions: []protocol.Version{protocol.OCPP16},
		Security: SecurityConfig{
			Profile:       SecurityProfileBasicAuth,
			Authenticator: AuthenticatorFunc(func(context.Context, AuthenticationRequest) (Principal, error) { return Principal{}, nil }),
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()
	_, response, err := (&websocket.Dialer{Subprotocols: []string{string(protocol.OCPP16)}}).Dial(
		"ws"+strings.TrimPrefix(httpServer.URL, "http")+"/CP-001", nil,
	)
	if err == nil || response == nil || response.StatusCode != http.StatusUnauthorized {
		t.Fatalf("error = %v, response = %#v, want HTTP 401", err, response)
	}
	if response.Header.Get("WWW-Authenticate") == "" {
		t.Fatal("WWW-Authenticate header is missing")
	}
}

func TestTLSProfileRejectsPlainWebSocket(t *testing.T) {
	server, err := New(Config{
		Versions: []protocol.Version{protocol.OCPP201},
		Security: SecurityConfig{
			Profile:       SecurityProfileTLSBasicAuth,
			Authenticator: AuthenticatorFunc(func(context.Context, AuthenticationRequest) (Principal, error) { return Principal{ID: "x"}, nil }),
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()
	_, response, err := (&websocket.Dialer{Subprotocols: []string{string(protocol.OCPP201)}}).Dial(
		"ws"+strings.TrimPrefix(httpServer.URL, "http")+"/CP-001", basicHeader("CP-001", "secret"),
	)
	if err == nil || response == nil || response.StatusCode != http.StatusForbidden {
		t.Fatalf("error = %v, response = %#v, want HTTP 403", err, response)
	}
}

func TestClientCertificateProfileRequiresVerifiedCertificate(t *testing.T) {
	server, err := New(Config{Security: SecurityConfig{
		Profile: SecurityProfileTLSClientCertificate,
		Authenticator: AuthenticatorFunc(func(context.Context, AuthenticationRequest) (Principal, error) {
			return Principal{ID: "certificate-principal"}, nil
		}),
	}})
	if err != nil {
		t.Fatal(err)
	}
	request := httptest.NewRequest(http.MethodGet, "https://example.test/CP-001", nil)
	request.TLS = &tls.ConnectionState{Version: tls.VersionTLS13}
	if _, status, err := server.authenticate(request, "CP-001", protocol.OCPP201); err == nil || status != http.StatusForbidden {
		t.Fatalf("status = %d, error = %v, want unverified certificate rejection", status, err)
	}
	certificate := &x509.Certificate{}
	request.TLS.PeerCertificates = []*x509.Certificate{certificate}
	request.TLS.VerifiedChains = [][]*x509.Certificate{{certificate}}
	principal, status, err := server.authenticate(request, "CP-001", protocol.OCPP201)
	if err != nil || status != 0 || principal.ID != "certificate-principal" {
		t.Fatalf("principal = %#v, status = %d, error = %v", principal, status, err)
	}
}

func TestAuthenticationRunsBeforeDuplicatePolicy(t *testing.T) {
	var duplicateCalls atomic.Int32
	server, err := New(Config{
		Versions: []protocol.Version{protocol.OCPP16},
		Security: SecurityConfig{
			Profile: SecurityProfileBasicAuth,
			Authenticator: AuthenticatorFunc(func(_ context.Context, request AuthenticationRequest) (Principal, error) {
				if request.Basic == nil || string(request.Basic.Password) != "valid" {
					return Principal{}, ErrAuthenticationFailed
				}
				return Principal{ID: request.Identity}, nil
			}),
		},
		OnDuplicateSession: func(context.Context, DuplicateSessionAttempt) (DuplicateSessionDecision, error) {
			duplicateCalls.Add(1)
			return ReplaceExistingSession, nil
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()
	first := dialStationWithHeader(t, httpServer.URL, protocol.OCPP16, basicHeader("CP-001", "valid"))
	defer first.Close()
	_, response, err := (&websocket.Dialer{Subprotocols: []string{string(protocol.OCPP16)}}).Dial(
		"ws"+strings.TrimPrefix(httpServer.URL, "http")+"/CP-001", basicHeader("CP-001", "invalid"),
	)
	if err == nil || response == nil || response.StatusCode != http.StatusUnauthorized {
		t.Fatalf("error = %v, response = %#v", err, response)
	}
	if duplicateCalls.Load() != 0 {
		t.Fatal("duplicate policy was called before authentication succeeded")
	}
}

func TestHandshakeLimiterAndSecurityEvents(t *testing.T) {
	events := make(chan SecurityEvent, 1)
	server, err := New(Config{
		Versions: []protocol.Version{protocol.OCPP16},
		Security: SecurityConfig{
			HandshakeLimiter: HandshakeLimiterFunc(func(context.Context, HandshakeAttempt) bool { return false }),
			OnEvent:          func(_ context.Context, event SecurityEvent) { events <- event },
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	httpServer := httptest.NewServer(server)
	defer httpServer.Close()
	_, response, err := (&websocket.Dialer{Subprotocols: []string{string(protocol.OCPP16)}}).Dial(
		"ws"+strings.TrimPrefix(httpServer.URL, "http")+"/CP-001", nil,
	)
	if err == nil || response == nil || response.StatusCode != http.StatusTooManyRequests {
		t.Fatalf("error = %v, response = %#v, want HTTP 429", err, response)
	}
	if event := <-events; event.Type != SecurityEventHandshakeRateLimited {
		t.Fatalf("event = %#v", event)
	}
}

func TestDefaultIdentityValidation(t *testing.T) {
	if err := ValidateIdentity("CP-001_ok.~"); err != nil {
		t.Fatal(err)
	}
	for _, identity := range []string{"", "bad identity", "bad/slash", strings.Repeat("x", 256)} {
		if err := ValidateIdentity(identity); err == nil {
			t.Errorf("ValidateIdentity(%q) succeeded", identity)
		}
	}
}

func TestIPRateLimiter(t *testing.T) {
	limiter, err := NewIPRateLimiter(2, 20*time.Millisecond)
	if err != nil {
		t.Fatal(err)
	}
	attempt := HandshakeAttempt{RemoteAddr: "192.0.2.1:1234"}
	if !limiter.Allow(context.Background(), attempt) {
		t.Fatal("first attempt was rejected")
	}
	if !limiter.Allow(context.Background(), attempt) {
		t.Fatal("second attempt was rejected")
	}
	if limiter.Allow(context.Background(), attempt) {
		t.Fatal("rate limit was not enforced")
	}
	if !limiter.Allow(context.Background(), HandshakeAttempt{RemoteAddr: "192.0.2.2:1234"}) {
		t.Fatal("a different IP shared the rate limit")
	}
	time.Sleep(25 * time.Millisecond)
	if !limiter.Allow(context.Background(), attempt) {
		t.Fatal("rate limit window did not reset")
	}
}

func basicHeader(username, password string) http.Header {
	header := make(http.Header)
	token := base64.StdEncoding.EncodeToString([]byte(username + ":" + password))
	header.Set("Authorization", "Basic "+token)
	return header
}

func dialStationWithHeader(t *testing.T, serverURL string, version protocol.Version, header http.Header) *websocket.Conn {
	t.Helper()
	dialer := websocket.Dialer{Subprotocols: []string{string(version)}}
	conn, _, err := dialer.Dial("ws"+strings.TrimPrefix(serverURL, "http")+"/CP-001", header)
	if err != nil {
		t.Fatal(err)
	}
	return conn
}
