package validation

// NotNil is a validation rule that checks if a value is not nil.
// NotNil only handles types including interface, pointer, slice, and map.
// All other types are considered valid.
var NotNil = &notNilRule{code: 1201}

type notNilRule struct {
	code int
}

// Validate checks if the given value is valid or not.
func (r *notNilRule) Validate(value interface{}) (code int, args []interface{}) {
	_, isNil := Indirect(value)
	if isNil {
		code = r.code
	}
	return
}
