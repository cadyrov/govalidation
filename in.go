package validation

import "errors"

// In returns a validation rule that checks if a value can be found in the given list of values.
// Note that the value being checked and the possible range of values must be of the same type.
// An empty value is considered valid. Use the Required rule to make sure a value is not empty.
func In(values ...interface{}) *InRule {
	return &InRule{
		elements: values,
		message:  MsgByCode(1101),
		code:     1101,
	}
}

type InRule struct {
	elements []interface{}
	message  string
	code     int
}

// Validate checks if the given value is valid or not.
func (r *InRule) Validate(value interface{}) ExternalError {
	value, isNil := Indirect(value)
	if isNil || IsEmpty(value) {
		return nil
	}

	for _, e := range r.elements {
		if e == value {
			return nil
		}
	}
	return NewExternalError(errors.New(r.message), r.code)
}

// Error sets the error message for the rule.
func (r *InRule) Error(message string) *InRule {
	r.message = message
	return r
}
