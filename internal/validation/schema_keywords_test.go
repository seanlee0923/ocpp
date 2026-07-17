package validation_test

import (
	"testing"

	"github.com/seanlee0923/ocpp/internal/validation"
)

// None of the 365 official OCPP 1.6/2.0.1/2.1 schemas use pattern, anyOf,
// oneOf, or allOf (see the Schema struct doc comment in validation.go), so
// TestAllGeneratedTypesValidateAndRoundTrip never exercises this logic.
// These synthetic schemas confirm the keywords behave correctly on their own
// so a future OCPP revision that does use them starts from working code.

func TestSchemaPattern(t *testing.T) {
	schema := &validation.Schema{Type: "string", Pattern: "^[A-Z]{2}[0-9]{4}$"}
	tests := []struct {
		name  string
		data  string
		valid bool
	}{
		{"matches", `"AB1234"`, true},
		{"wrong shape", `"ab1234"`, false},
		{"too short", `"AB12"`, false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := validation.ValidateJSON("Pattern.json", schema, []byte(test.data))
			if test.valid && err != nil {
				t.Fatalf("ValidateJSON(%s) = %v, want nil", test.data, err)
			}
			if !test.valid && err == nil {
				t.Fatalf("ValidateJSON(%s) = nil, want an error", test.data)
			}
		})
	}
}

func TestSchemaPatternRejectsInvalidRegex(t *testing.T) {
	schema := &validation.Schema{Type: "string", Pattern: "("}
	if err := validation.ValidateJSON("BadPattern.json", schema, []byte(`"anything"`)); err == nil {
		t.Fatal("expected an error for an uncompilable pattern")
	}
}

func TestSchemaAnyOf(t *testing.T) {
	schema := &validation.Schema{AnyOf: []*validation.Schema{
		{Type: "string"},
		{Type: "integer"},
	}}
	tests := []struct {
		name  string
		data  string
		valid bool
	}{
		{"matches first branch", `"text"`, true},
		{"matches second branch", `42`, true},
		{"matches neither branch", `true`, false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := validation.ValidateJSON("AnyOf.json", schema, []byte(test.data))
			if test.valid && err != nil {
				t.Fatalf("ValidateJSON(%s) = %v, want nil", test.data, err)
			}
			if !test.valid && err == nil {
				t.Fatalf("ValidateJSON(%s) = nil, want an error", test.data)
			}
		})
	}
}

func TestSchemaOneOfRequiresExactlyOneMatch(t *testing.T) {
	schema := &validation.Schema{OneOf: []*validation.Schema{
		{Type: "integer"},
		{Type: "integer", HasMaximum: true, Maximum: 100},
	}}
	tests := []struct {
		name  string
		data  string
		valid bool
	}{
		{"no branch matches", `"text"`, false},
		// 50 satisfies both "any integer" and "integer <= 100", so oneOf's
		// exactly-one requirement rejects it.
		{"both branches match", `50`, false},
		{"exactly one branch matches", `500`, true},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := validation.ValidateJSON("OneOf.json", schema, []byte(test.data))
			if test.valid && err != nil {
				t.Fatalf("ValidateJSON(%s) = %v, want nil", test.data, err)
			}
			if !test.valid && err == nil {
				t.Fatalf("ValidateJSON(%s) = nil, want an error", test.data)
			}
		})
	}
}

func TestSchemaAllOfRequiresEveryBranchToMatch(t *testing.T) {
	schema := &validation.Schema{AllOf: []*validation.Schema{
		{Type: "integer"},
		{Type: "integer", HasMinimum: true, Minimum: 10},
	}}
	tests := []struct {
		name  string
		data  string
		valid bool
	}{
		{"satisfies every branch", `20`, true},
		{"fails the minimum branch", `5`, false},
		{"fails the type branch", `"20"`, false},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			err := validation.ValidateJSON("AllOf.json", schema, []byte(test.data))
			if test.valid && err != nil {
				t.Fatalf("ValidateJSON(%s) = %v, want nil", test.data, err)
			}
			if !test.valid && err == nil {
				t.Fatalf("ValidateJSON(%s) = nil, want an error", test.data)
			}
		})
	}
}
