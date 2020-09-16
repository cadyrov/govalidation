package bi

import (
	"strconv"
	"unicode/utf8"

	validation "github.com/cadyrov/govalidation"
	"github.com/cadyrov/govalidation/is"
)

var OPGNIp = &ogrnIpRule{code: 2850}

type ogrnIpRule struct {
	code int
}

func (oip *ogrnIpRule) Validate(value interface{}) (code int, args []interface{}) {
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
		code = oip.code
		return
	}

	if code, args = is.Digit.Validate(s); code != 0 {
		return
	}

	if utf8.RuneCountInString(s) != 15 {
		code = 2851
		return
	}

	s14 := s[:len(s)-1]
	s15 := s[14]
	i14, _ := strconv.ParseInt(string(s14), 10, 64)
	i15, _ := strconv.ParseInt(string(s15), 10, 64)
	os := i14 % 13
	sn := strconv.FormatInt(os, 10)
	snX, _ := strconv.ParseInt(string(sn[len(sn)-1]), 10, 12)
	if snX != i15 {
		code = 2552
		return
	}
	return
}

func (r *ogrnIpRule) Error(message string) *ogrnIpRule {
	return &ogrnIpRule{
		code: 2850,
	}
}
