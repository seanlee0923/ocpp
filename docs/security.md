# 인증과 TLS

**한국어** | [English](en/security.md)

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

`Authenticator.Authenticate`와 `HandshakeLimiter.Allow`는 해당 handshake 하나를 처리하는
HTTP goroutine에서 동기 실행되므로, 느리거나 멈춘 구현체는 그 Charging Station 한 대의
연결 시도만 지연시키고 서버 전체를 막지는 않는다 — 다만 영원히 반환하지 않으면 그 goroutine과
대기 중인 upgrade는 클라이언트가 TCP 연결을 유지하는 한 계속 남는다. 두 콜백 모두 빠르게
반환해야 한다.

실행 가능한 설정 예제:

- [Security Profile 1 Basic Auth](../examples/basic-auth)
- [Security Profile 2 TLS와 Basic Auth](../examples/tls-basic-auth)
- [Security Profile 3 mutual TLS](../examples/mtls)

TLS 예제는 실제 인증서를 저장소에 포함하지 않고 환경변수로 certificate, private key와
client CA 경로를 받는다.
