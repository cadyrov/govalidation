package validation

// In returns a validation rule that checks if a value can be found in the given list of values.
// Note that the value being checked and the possible range of values must be of the same type.
// An empty value is considered valid. Use the Required rule to make sure a value is not empty.
func In(values ...interface{}) *InRule {
	return &InRule{
		elements: values,
		code:     1101,
	}
}

type InRule struct {
	elements []interface{}
	message  string
	code     int
}

// Validate checks if the given value is valid or not.
func (r *InRule) Validate(value interface{}) (code int, args []interface{}) {
	value, isNil := Indirect(value)
	if isNil || IsEmpty(value) {
		return
	}

	for _, e := range r.elements {
		if e == value {
			return
		}
	}
	code = r.code
	return
}
