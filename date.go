package validation

import (
	"time"
)

type DateRule struct {
	layout       string
	min, max     time.Time
	message      string
	rangeMessage string
	code         int
	rangeCode    int
}

// Date returns a validation rule that checks if a string value is in a format that can be parsed into a date.
// The format of the date should be specified as the layout parameter which accepts the same value as that for time.Parse.
// For example,
//    validation.Date(time.ANSIC)
//    validation.Date("02 Jan 06 15:04 MST")
//    validation.Date("2006-01-02")
//
// By calling Min() and/or Max(), you can let the Date rule to check if a parsed date value is within
// the specified date range.
//
// An empty value is considered valid. Use the Required rule to make sure a value is not empty.
func Date(layout string) *DateRule {
	return &DateRule{
		layout:    layout,
		code:      1102,
		rangeCode: 1103,
	}
}

// Min sets the minimum date range. A zero value means skipping the minimum range validation.
func (r *DateRule) Min(min time.Time) *DateRule {
	r.min = min
	return r
}

// Max sets the maximum date range. A zero value means skipping the maximum range validation.
func (r *DateRule) Max(max time.Time) *DateRule {
	r.max = max

	return r
}

// Validate checks if the given value is a valid date.
func (r *DateRule) Validate(value interface{}) (code int, args []interface{}) {
	value, isNil := Indirect(value)
	if isNil || IsEmpty(value) {
		return
	}
	str, code := EnsureString(value)
	if code != 0 {
		return
	}
	date, errt := time.Parse(r.layout, str)
	if errt != nil {
		code = r.code
		return
	}
	if !r.min.IsZero() && r.min.After(date) || !r.max.IsZero() && date.After(r.max) {
		code = r.rangeCode
		return
	}
	return
}
