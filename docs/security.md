# 인증과 TLS

`csms.SecurityConfig`는 WebSocket handshake 정책을 구성한다.

| 값 | 정책 |
|---|---|
| `SecurityProfileUnsecured` | TLS/Basic Auth를 강제하지 않음 |
| `SecurityProfileBasicAuth` | Basic Auth 요구 |
| `SecurityProfileTLSBasicAuth` | TLS와 Basic Auth 요구 |
| `SecurityProfileTLSClientCertificate` | TLS와 검증된 client certificate 요구 |

Profile 1 이상에서는 `Authenticator`가 필요하다.

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

Basic Auth username은 기본적으로 URL identity와 같아야 한다. 특별한 배포 환경에서만
`AllowBasicUsernameMismatch`를 사용한다. password byte slice는 callback이 끝난 뒤 지워지므로
보관하지 않는다.

Profile 2/3의 TLS는 `http.Server` 또는 앞단 TLS listener가 제공해야 한다. client certificate는
Go TLS가 검증한 `VerifiedChains`가 있어야 한다. reverse proxy에서 전달한 임의 header를
인증된 TLS 정보로 신뢰하지 않는다.

`HandshakeLimiter`, `ValidateIdentity`, `OnEvent`로 rate limit, identity 정책과 보안 이벤트를
연결할 수 있다. OCPP 인증서 Action의 transport 타입은 제공하지만 CA, private key, OCSP/CRL과
인증서 lifecycle은 애플리케이션 인프라의 책임이다.
