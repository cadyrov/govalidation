package validation

// NotIn returns a validation rule that checks if a value os absent from, the given list of values.
// Note that the value being checked and the possible range of values must be of the same type.
// An empty value is considered valid. Use the Required rule to make sure a value is not empty.
func NotIn(values ...interface{}) *NotInRule {
	return &NotInRule{
		elements: values,
		code:     1107,
	}
}

type NotInRule struct {
	elements []interface{}
	code     int
}

// Validate checks if the given value is valid or not.
func (r *NotInRule) Validate(value interface{}) (code int, args []interface{}) {
	value, isNil := Indirect(value)
	if isNil || IsEmpty(value) {
		return
	}
	for _, e := range r.elements {
		if e == value {
			code = r.code
			return
		}
	}
	return
}
