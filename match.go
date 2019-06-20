package validation

import (
	"errors"
	"regexp"
)

// Match returns a validation rule that checks if a value matches the specified regular expression.
// This rule should only be used for validating strings and byte slices, or a validation error will be reported.
// An empty value is considered valid. Use the Required rule to make sure a value is not empty.
func Match(re *regexp.Regexp) *MatchRule {
	return &MatchRule{
		re:      re,
		message: MsgByCode(1105),
		code:    1105,
	}
}

type MatchRule struct {
	re      *regexp.Regexp
	message string
	code    int
}

// Validate checks if the given value is valid or not.
func (v *MatchRule) Validate(value interface{}) ExternalError {
	value, isNil := Indirect(value)
	if isNil {
		return nil
	}

	isString, str, isBytes, bs := StringOrBytes(value)
	if isString && (str == "" || v.re.MatchString(str)) {
		return nil
	} else if isBytes && (len(bs) == 0 || v.re.Match(bs)) {
		return nil
	}
	return NewExternalError(errors.New(v.message), v.code)
}

// Error sets the error message for the rule.
func (v *MatchRule) Error(message string) *MatchRule {
	v.message = message
	return v
}
