# 검증과 오류 처리

생성 payload는 `Validate`와 `ValidateJSON`을 제공한다. 네트워크 입력은 struct decode 전에 raw
JSON으로 검증되어 unknown property, required, enum, 타입, 길이, 범위와 format 위반을 막는다.
직접 `json.Unmarshal`한 값은 자동 검증되지 않으므로 애플리케이션에서 생성 타입을 별도로
사용할 때는 `Validate()`를 호출한다.

Inbound request 위반은 OCPP-J CALLERROR로 변환된다.

| 제약 | 결과 |
|---|---|
| unknown property, enum, 길이·범위·format | `PropertyConstraintViolation` |
| required 또는 배열 항목 수 | version별 occurrence violation |
| JSON field 타입 | `TypeConstraintViolation` |
| payload 문법 | 1.6 `FormationViolation`, 2.x `FormatViolation` |

OCPP 1.6 wire spelling은 `OccurenceConstraintViolation`, OCPP 2.x는
`OccurrenceConstraintViolation`이다. 라이브러리는 이 규격 차이를 그대로 유지한다.

분기 가능한 주요 오류에는 설정, handler 등록, unique ID, pending limit와 session 종료
sentinel이 있다. 문자열을 비교하지 말고 `errors.Is`를 사용한다. 원격 CALLERROR와 생성 검증
오류는 각각 `errors.As`와 `protocol.IsValidationError`로 확인한다.

handler의 일반 Go error는 내부 내용을 노출하지 않는 `InternalError`가 된다. 공개 close reason과
구조화 로그에도 credential, payload, idToken과 handler error 문자열을 포함하지 않는다.
