package bi

import (
	validation "github.com/cadyrov/govalidation"
	"github.com/cadyrov/govalidation/is"
	"strconv"
	"unicode/utf8"
)

var Inn10 = &inn10Rule{code: 2810}

type inn10Rule struct {
	code int
}

func (inn *inn10Rule) Validate(value interface{}) (code int, args []interface{}) {
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
		code = inn.code
		return
	}

	if code, args = is.Digit.Validate(s); code != 0 {
		return
	}

	if utf8.RuneCountInString(s) != 10 {
		code = 2811
		return
	}

	coefficients := []int64{2, 4, 10, 3, 5, 9, 4, 6, 8}

	if !checkInnDigits(s, coefficients) {
		code = 2812
		return
	}
	return
}
