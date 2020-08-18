package bi

import (
	validation "github.com/cadyrov/govalidation"
	"github.com/cadyrov/govalidation/is"
	"strconv"
	"unicode/utf8"
)

var Inn12 = &inn12Rule{code: 2820}

type inn12Rule struct {
	code int
}

func (inn *inn12Rule) Validate(value interface{}) (code int, args []interface{}) {
	value, isNil := validation.Indirect(value)
	if isNil || validation.IsEmpty(value) {
		return
	}
	var s string
	switch value.(type) {
	case string:
		s = value.(string)
	case int:
		s = strconv.Itoa(value.(int))
	case int64:
		s = strconv.FormatInt(value.(int64), 10)
	default:
		code = 2821
		return
	}

	if code, args = is.Digit.Validate(s); code != 0 {
		return
	}

	if utf8.RuneCountInString(s) != 12 {
		code = 2822
		return
	}

	coefficients11 := []int64{7, 2, 4, 10, 3, 5, 9, 4, 6, 8}
	coefficients12 := []int64{3, 7, 2, 4, 10, 3, 5, 9, 4, 6, 8}
	if !checkInnDigits(s, coefficients11) || !checkInnDigits(s, coefficients12) {
		code = 2823
		return
	}
	return
}

func (r *inn12Rule) Error(message string) *inn12Rule {
	return &inn12Rule{
		code: 2802,
	}
}
