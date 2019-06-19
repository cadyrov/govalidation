package bi

import (
	"errors"
	validation "github.com/cadyrov/govalidation"
	"github.com/cadyrov/govalidation/is"
	"strconv"
	"unicode/utf8"
)

var OPGNIp = &ogrnIpRule{message: validation.MsgByCode(2804), code: 2804}

type ogrnIpRule struct {
	message string
	code    int
}

func (inn *ogrnIpRule) Validate(value interface{}) validation.ExternalError {
	value, isNil := validation.Indirect(value)
	if isNil || validation.IsEmpty(value) {
		return nil
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
		return validation.NewExternalError(errors.New("can't parse value "), 2804)
	}

	if err := is.Digit.Validate(s); err != nil {
		return err
	}

	if utf8.RuneCountInString(s) != 15 {
		return validation.NewExternalError(errors.New("only 15 digits"), 2804)
	}

	s14 := s[:len(s)-1]
	s15 := s[14]
	i14, _ := strconv.ParseInt(string(s14), 10, 64)
	i15, _ := strconv.ParseInt(string(s15), 10, 64)
	os := i14 % 13
	sn := strconv.FormatInt(os, 10)
	snX, _ := strconv.ParseInt(string(sn[len(sn)-1]), 10, 12)
	if snX != i15 {
		return validation.NewExternalError(errors.New("invalid control data"), 2804)
	}
	return nil
}

func (r *ogrnIpRule) Error(message string) *ogrnIpRule {
	return &ogrnIpRule{
		message: message,
		code:    2804,
	}
}
