package bi

import (
	validation "github.com/cadyrov/govalidation"
	"github.com/cadyrov/govalidation/is"
	"strconv"
	"unicode/utf8"
)

var Snils = &snilsRule{code: 2808}

type snilsRule struct {
	code int
}

func (inn *snilsRule) Validate(value interface{}) (code int, args []interface{}) {
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
		code = 2880
		return
	}

	if code, args = is.Digit.Validate(s); code != 0 {
		return
	}

	if utf8.RuneCountInString(s) != 11 {
		code = 2880
		return
	}

	sumSnils := int64(0)
	for i := 1; i < 10; i++ {
		x, _ := strconv.ParseInt(string(s[9-i]), 10, 64)
		sumSnils += int64(i) * x
	}
	cntrl := snilsControl(sumSnils)
	must, _ := strconv.ParseInt(string(s[9:]), 10, 64)
	if must != cntrl {
		code = 2880
		return
	}

	return
}

func snilsControl(input int64) int64 {
	if input < 100 {
		return input
	} else if input == 100 {
		return 0
	} else {
		return snilsControl(input % 101)
	}

}
