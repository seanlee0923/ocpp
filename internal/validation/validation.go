package validation

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"math"
	"net/url"
	"regexp"
	"strconv"
	"sync"
	"time"
	"unicode/utf8"
)

type Schema struct {
	Type                 string
	Properties           map[string]*Schema
	Required             []string
	AllowAdditional      bool
	AdditionalProperties *Schema
	Items                *Schema
	Enum                 []string
	MinLength            int
	HasMinLength         bool
	MaxLength            int
	HasMaxLength         bool
	Minimum              float64
	HasMinimum           bool
	Maximum              float64
	HasMaximum           bool
	ExclusiveMinimum     bool
	ExclusiveMaximum     bool
	MultipleOf           float64
	MinItems             int
	HasMinItems          bool
	MaxItems             int
	HasMaxItems          bool
	Pattern              string
	Format               string
	AnyOf                []*Schema
	OneOf                []*Schema
	AllOf                []*Schema
}

type Error struct {
	Schema string
	Path   string
	Reason string
}

func (e *Error) Error() string {
	if e.Path == "" {
		return fmt.Sprintf("payload does not conform to %s: %s", e.Schema, e.Reason)
	}
	return fmt.Sprintf("payload does not conform to %s at %s: %s", e.Schema, e.Path, e.Reason)
}

func IsError(err error) bool {
	var target *Error
	return errors.As(err, &target)
}

func Validate(name string, schema *Schema, value any) error {
	data, err := json.Marshal(value)
	if err != nil {
		return fmt.Errorf("marshal %s payload: %w", name, err)
	}
	return ValidateJSON(name, schema, data)
}

func ValidateJSON(name string, schema *Schema, data []byte) error {
	decoder := json.NewDecoder(bytes.NewReader(data))
	decoder.UseNumber()
	var value any
	if err := decoder.Decode(&value); err != nil {
		return fmt.Errorf("decode %s payload: %w", name, err)
	}
	var trailing any
	if err := decoder.Decode(&trailing); err != io.EOF {
		return &Error{Schema: name, Reason: "trailing JSON value"}
	}
	if err := validate(schema, value, "$", name); err != nil {
		return err
	}
	return nil
}

func validate(schema *Schema, value any, path, name string) error {
	if schema == nil {
		return nil
	}
	for _, candidate := range schema.AllOf {
		if err := validate(candidate, value, path, name); err != nil {
			return err
		}
	}
	if len(schema.AnyOf) > 0 && matching(schema.AnyOf, value, path, name) == 0 {
		return invalid(name, path, "does not match any allowed schema")
	}
	if len(schema.OneOf) > 0 && matching(schema.OneOf, value, path, name) != 1 {
		return invalid(name, path, "does not match exactly one allowed schema")
	}

	switch schema.Type {
	case "object":
		object, ok := value.(map[string]any)
		if !ok {
			return invalid(name, path, "must be an object")
		}
		for _, required := range schema.Required {
			if _, ok := object[required]; !ok {
				return invalid(name, path, fmt.Sprintf("required property %q is missing", required))
			}
		}
		for key, child := range object {
			if property, ok := schema.Properties[key]; ok {
				if err := validate(property, child, childPath(path, key), name); err != nil {
					return err
				}
				continue
			}
			if schema.AdditionalProperties != nil {
				if err := validate(schema.AdditionalProperties, child, childPath(path, key), name); err != nil {
					return err
				}
				continue
			}
			if !schema.AllowAdditional {
				return invalid(name, childPath(path, key), "unknown property")
			}
		}
	case "array":
		array, ok := value.([]any)
		if !ok {
			return invalid(name, path, "must be an array")
		}
		if schema.HasMinItems && len(array) < schema.MinItems {
			return invalid(name, path, fmt.Sprintf("must contain at least %d items", schema.MinItems))
		}
		if schema.HasMaxItems && len(array) > schema.MaxItems {
			return invalid(name, path, fmt.Sprintf("must contain at most %d items", schema.MaxItems))
		}
		for index, item := range array {
			if err := validate(schema.Items, item, fmt.Sprintf("%s[%d]", path, index), name); err != nil {
				return err
			}
		}
	case "string":
		text, ok := value.(string)
		if !ok {
			return invalid(name, path, "must be a string")
		}
		length := utf8.RuneCountInString(text)
		if schema.HasMinLength && length < schema.MinLength {
			return invalid(name, path, fmt.Sprintf("must contain at least %d characters", schema.MinLength))
		}
		if schema.HasMaxLength && length > schema.MaxLength {
			return invalid(name, path, fmt.Sprintf("must contain at most %d characters", schema.MaxLength))
		}
		if len(schema.Enum) > 0 && !contains(schema.Enum, text) {
			return invalid(name, path, fmt.Sprintf("value %q is not allowed", text))
		}
		if schema.Pattern != "" {
			pattern, err := compiledPattern(schema.Pattern)
			if err != nil || !pattern.MatchString(text) {
				return invalid(name, path, "does not match the required pattern")
			}
		}
		if schema.Format == "date-time" {
			if _, err := time.Parse(time.RFC3339, text); err != nil {
				return invalid(name, path, "must be an RFC 3339 date-time")
			}
		}
		if schema.Format == "uri" {
			parsed, err := url.Parse(text)
			if err != nil || parsed.Scheme == "" {
				return invalid(name, path, "must be an absolute URI")
			}
		}
	case "integer":
		number, ok := value.(json.Number)
		if !ok {
			return invalid(name, path, "must be an integer")
		}
		parsed, err := strconv.ParseFloat(number.String(), 64)
		if err != nil || math.Trunc(parsed) != parsed {
			return invalid(name, path, "must be an integer")
		}
		return validateNumber(schema, parsed, path, name)
	case "number":
		number, ok := value.(json.Number)
		if !ok {
			return invalid(name, path, "must be a number")
		}
		parsed, err := number.Float64()
		if err != nil {
			return invalid(name, path, "must be a number")
		}
		return validateNumber(schema, parsed, path, name)
	case "boolean":
		if _, ok := value.(bool); !ok {
			return invalid(name, path, "must be a boolean")
		}
	}
	return nil
}

func validateNumber(schema *Schema, value float64, path, name string) error {
	if schema.HasMinimum && (value < schema.Minimum || schema.ExclusiveMinimum && value == schema.Minimum) {
		return invalid(name, path, "is below the allowed minimum")
	}
	if schema.HasMaximum && (value > schema.Maximum || schema.ExclusiveMaximum && value == schema.Maximum) {
		return invalid(name, path, "is above the allowed maximum")
	}
	if schema.MultipleOf != 0 {
		quotient := value / schema.MultipleOf
		if math.Abs(quotient-math.Round(quotient)) > 1e-9 {
			return invalid(name, path, fmt.Sprintf("must be a multiple of %v", schema.MultipleOf))
		}
	}
	return nil
}

func matching(candidates []*Schema, value any, path, name string) int {
	count := 0
	for _, candidate := range candidates {
		if validate(candidate, value, path, name) == nil {
			count++
		}
	}
	return count
}

func invalid(name, path, reason string) error {
	return &Error{Schema: name, Path: path, Reason: reason}
}

func childPath(parent, child string) string {
	if parent == "$" {
		return parent + "." + child
	}
	return parent + "." + child
}

func contains(values []string, value string) bool {
	for _, candidate := range values {
		if candidate == value {
			return true
		}
	}
	return false
}

var patterns sync.Map

func compiledPattern(source string) (*regexp.Regexp, error) {
	if value, ok := patterns.Load(source); ok {
		return value.(*regexp.Regexp), nil
	}
	pattern, err := regexp.Compile(source)
	if err != nil {
		return nil, err
	}
	actual, _ := patterns.LoadOrStore(source, pattern)
	return actual.(*regexp.Regexp), nil
}
