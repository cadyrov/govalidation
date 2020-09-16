package bi

import (
	"strconv"
	"unicode/utf8"

	validation "github.com/cadyrov/govalidation"
	"github.com/cadyrov/govalidation/is"
)

var OGRNLaw = &ogrnLawRule{code: 2804}

type ogrnLawRule struct {
	code int
}

func (ogl *ogrnLawRule) Validate(value interface{}) (code int, args []interface{}) {
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
		code = ogl.code
		return
	}

	if code, args = is.Digit.Validate(s); code != 0 {
		return
	}

	if utf8.RuneCountInString(s) != 13 {
		code = 2841
		return
	}

	s12 := s[:len(s)-1]
	s13 := s[12]
	i12, _ := strconv.ParseInt(string(s12), 10, 64)
	i13, _ := strconv.ParseInt(string(s13), 10, 64)
	os := i12 % 11
	sn := strconv.FormatInt(os, 10)
	snX, _ := strconv.ParseInt(string(sn[len(sn)-1]), 10, 12)
	if snX != i13 {
		code = 2842
		return
	}
	return
}

func (r *ogrnLawRule) Error(message string) *ogrnLawRule {
	return &ogrnLawRule{
		code: 2805,
	}
}
