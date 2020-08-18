package validation

// Required is a validation rule that checks if a value is not empty.
// A value is considered not empty if
// - integer, float: not zero
// - bool: true
// - string, array, slice, map: len() > 0
// - interface, pointer: not nil and the referenced value is not empty
// - any other types
var Required = &requiredRule{skipNil: false, code: 1202}

// NilOrNotEmpty checks if a value is a nil pointer or a value that is not empty.
// NilOrNotEmpty differs from Required in that it treats a nil pointer as valid.
var NilOrNotEmpty = &requiredRule{skipNil: true, code: 1202}

type requiredRule struct {
	skipNil bool
	code    int
}

// Validate checks if the given value is valid or not.
func (v *requiredRule) Validate(value interface{}) (code int, args []interface{}) {
	value, isNil := Indirect(value)
	if v.skipNil && !isNil && IsEmpty(value) || !v.skipNil && (isNil || IsEmpty(value)) {
		code = v.code
		return
	}
	return
}
