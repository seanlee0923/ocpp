# Authentication and TLS

[한국어](../security.md) | **English**

`csms.SecurityConfig` configures the WebSocket handshake policy.

| Value | Policy |
|---|---|
| `SecurityProfileUnsecured` | No TLS/Basic Auth enforced |
| `SecurityProfileBasicAuth` | Basic Auth required |
| `SecurityProfileTLSBasicAuth` | TLS and Basic Auth required |
| `SecurityProfileTLSClientCertificate` | TLS and a verified client certificate required |

An `Authenticator` is required for Profile 1 and above.

```go
Security: csms.SecurityConfig{
    Profile: csms.SecurityProfileBasicAuth,
    Authenticator: csms.AuthenticatorFunc(func(
        ctx context.Context,
        request csms.AuthenticationRequest,
    ) (csms.Principal, error) {
        if !validCredential(request.Identity, request.Basic) {
            return csms.Principal{}, csms.ErrAuthenticationFailed
        }
        return csms.Principal{ID: request.Identity}, nil
    }),
},
```

By default the Basic Auth username must match the URL identity. Use
`AllowBasicUsernameMismatch` only for special deployments. The password
byte slice is cleared after the callback returns, so it should not be
retained.

TLS for Profile 2/3 must be provided by `http.Server` or a front-end TLS
listener. A client certificate requires `VerifiedChains` as verified by Go's
TLS stack. Arbitrary headers forwarded by a reverse proxy are never trusted
as verified TLS information.

`HandshakeLimiter`, `ValidateIdentity`, and `OnEvent` let you wire in rate
limiting, identity policy, and security events. The library provides the
transport type for OCPP certificate Actions, but the CA, private keys,
OCSP/CRL, and certificate lifecycle are the application infrastructure's
responsibility.

`Authenticator.Authenticate` and `HandshakeLimiter.Allow` run synchronously
on the HTTP goroutine handling that one handshake, so a slow or hung
implementation only stalls that single Charging Station's connection
attempt rather than the whole server — but if it never returns, that
goroutine and the pending upgrade leak for as long as the client keeps the
underlying TCP connection open. Both should return promptly.

Runnable configuration examples:

- [Security Profile 1 Basic Auth](../../examples/basic-auth)
- [Security Profile 2 TLS with Basic Auth](../../examples/tls-basic-auth)
- [Security Profile 3 mutual TLS](../../examples/mtls)

The TLS examples do not include real certificates in the repository; they
read the certificate, private key, and client CA paths from environment
variables.
