package bi

import (
	"errors"
	validation "github.com/cadyrov/govalidation"
	"github.com/cadyrov/govalidation/is"
	"strconv"
	"unicode/utf8"
)

var OGRNLaw = &ogrnLawRule{message: validation.MsgByCode(2805), code: 2805}

type ogrnLawRule struct {
	message string
	code    int
}

func (inn *ogrnLawRule) Validate(value interface{}) validation.ExternalError {
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
		return validation.NewExternalError(errors.New("can't parse value "), 2805)
	}

	if err := is.Digit.Validate(s); err != nil {
		return err
	}

	if utf8.RuneCountInString(s) != 13 {
		return validation.NewExternalError(errors.New("only 13 digits"), 2805)
	}

	s12 := s[:len(s)-1]
	s13 := s[12]
	i12, _ := strconv.ParseInt(string(s12), 10, 64)
	i13, _ := strconv.ParseInt(string(s13), 10, 64)
	os := i12 % 11
	sn := strconv.FormatInt(os, 10)
	snX, _ := strconv.ParseInt(string(sn[len(sn)-1]), 10, 12)
	if snX != i13 {
		return validation.NewExternalError(errors.New("invalid control data"), 2805)
	}
	return nil
}

func (r *ogrnLawRule) Error(message string) *ogrnLawRule {
	return &ogrnLawRule{
		message: message,
		code:    2805,
	}
}
