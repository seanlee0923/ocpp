# Validation and error handling

[한국어](../validation-and-errors.md) | **English**

Generated payloads provide `Validate` and `ValidateJSON`. Network input is
validated as raw JSON before it is decoded into a struct, blocking unknown
properties and violations of required, enum, type, length, range, and
format constraints. A value produced by calling `json.Unmarshal` directly is
not automatically validated, so applications that use generated types on
their own should call `Validate()` explicitly.

Inbound request violations are converted into an OCPP-J CALLERROR.

| Constraint | Result |
|---|---|
| Unknown property, enum, length/range/format | `PropertyConstraintViolation` |
| Required field or array item count | Version-specific occurrence violation |
| JSON field type | `TypeConstraintViolation` |
| Payload syntax | 1.6 `FormationViolation`, 2.x `FormatViolation` |

The OCPP 1.6 wire spelling is `OccurenceConstraintViolation`; OCPP 2.x uses
`OccurrenceConstraintViolation`. The library preserves this spec
inconsistency as-is.

Errors that can be branched on include configuration, handler registration,
unique ID, pending limit, and session closure sentinels. Use `errors.Is`
instead of comparing strings. A remote CALLERROR and a generated validation
error can be checked with `errors.As` and `protocol.IsValidationError`
respectively.

A plain Go error returned by a handler becomes an `InternalError` that does
not expose its internal content. Public close reasons and structured logs
likewise never include credentials, payloads, idTokens, or handler error
strings.
