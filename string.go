package validation

type stringValidator func(string) bool

// StringRule is a rule that checks a string variable using a specified stringValidator.
type StringRule struct {
	validate stringValidator
	code     int
}

// NewStringRule creates a new validation rule using a function that takes a string value and returns a bool.
// The rule returned will use the function to check if a given string or byte slice is valid or not.
// An empty value is considered to be valid. Please use the Required rule to make sure a value is not empty.
func NewStringRule(validator stringValidator, code int) *StringRule {
	return &StringRule{
		validate: validator,
		code:     code,
	}
}

// Error sets the error message for the rule.
func (v *StringRule) Error(message string, code int) *StringRule {
	return NewStringRule(v.validate, code)
}

// Validate checks if the given value is valid or not.
func (v *StringRule) Validate(value interface{}) (code int, args []interface{}) {
	value, isNil := Indirect(value)
	if isNil || IsEmpty(value) {
		return
	}
	str, code := EnsureString(value)
	if code != 0 {
		return
	}

	if v.validate(str) {
		return
	}
	code = v.code
	return
}
