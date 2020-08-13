package validation

import (
	"unicode/utf8"
)

// Length returns a validation rule that checks if a value's length is within the specified range.
// If max is 0, it means there is no upper bound for the length.
// This rule should only be used for validating strings, slices, maps, and arrays.
// An empty value is considered valid. Use the Required rule to make sure a value is not empty.
func Length(min, max int) *LengthRule {
	code := 1104
	if min == 0 && max > 0 {
		code = 1301
	} else if min > 0 && max == 0 {
		code = 1302
	} else if min > 0 && max > 0 {
		if min == max {
			code = 1303
		} else {
			code = 1304
		}
	}
	return &LengthRule{
		min:  min,
		max:  max,
		code: code,
	}
}

// RuneLength returns a validation rule that checks if a string's rune length is within the specified range.
// If max is 0, it means there is no upper bound for the length.
// This rule should only be used for validating strings, slices, maps, and arrays.
// An empty value is considered valid. Use the Required rule to make sure a value is not empty.
// If the value being validated is not a string, the rule works the same as Length.
func RuneLength(min, max int) *LengthRule {
	r := Length(min, max)
	r.rune = true
	return r
}

type LengthRule struct {
	min, max int
	message  string
	rune     bool
	code     int
}

// Validate checks if the given value is valid or not.
func (v *LengthRule) Validate(value interface{}) (code int, args []interface{}) {
	value, isNil := Indirect(value)
	if isNil || IsEmpty(value) {
		return
	}

	var l int
	if s, ok := value.(string); ok && v.rune {
		l = utf8.RuneCountInString(s)
	} else if l, code = LengthOfValue(value); code != 0 {
		return
	}

	if v.min > 0 && l < v.min || v.max > 0 && l > v.max {
		code = v.code
	}
	return
}

// Error sets the error message for the rule.
func (v *LengthRule) Error(message string) *LengthRule {
	v.message = message
	return v
}
