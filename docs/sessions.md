# 세션과 연결 생명주기

`Server`는 identity마다 하나의 활성 `Session`을 관리한다. identity는 기본적으로 RFC 3986
unreserved 문자와 최대 255 UTF-8 byte로 제한된다.

```go
session, ok := server.Session("CP-001")
sessions := server.Sessions()
info := session.Info()
```

`SessionInfo`에는 version, principal, remote address, 연결·최근 송수신·pong 시간과 상태가
있다. 반환된 principal attribute와 session slice는 복사본이므로 내부 registry를 직접
수정하지 않는다.

중복 identity는 기본적으로 HTTP 409로 거부한다. 필요한 경우에만 정책 callback으로 기존
연결 교체를 승인한다.

```go
OnDuplicateSession: func(ctx context.Context, attempt csms.DuplicateSessionAttempt) (
    csms.DuplicateSessionDecision, error,
) {
    return csms.ReplaceExistingSession, nil
},
```

새 WebSocket upgrade가 성공하기 전에는 기존 세션을 종료하지 않는다. 교체된 세션의 오류는
`ErrSessionReplaced`다.

`PingInterval`, `PongTimeout`, `IdleTimeout`으로 연결 상태를 관리한다. `Session.Done()`은
종료 시 닫히고 `Session.Err()`에 원인이 남는다. `Session.Context()`는 handler와 background
작업의 취소 기준으로 사용할 수 있다.

`Server.Shutdown(ctx)`은 새 연결을 거부하고 활성 세션과 handler가 끝날 때까지 기다린다.
반복 호출은 안전하지만 종료한 Server는 재사용할 수 없다.
