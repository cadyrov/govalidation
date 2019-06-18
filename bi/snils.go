package bi

import (
	"errors"
	validation "github.com/cadyrov/govalidation"
	"github.com/cadyrov/govalidation/is"
	"strconv"
	"unicode/utf8"
)

var Snils = &snilsRule{message: validation.MsgByCode(1900), code: 1900}

type snilsRule struct {
	message string
	code    int
}

func (inn *snilsRule) Validate(value interface{}) error {
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
		return errors.New("can't parse value ")
	}

	if err := is.Digit.Validate(s); err != nil {
		return err
	}

	if utf8.RuneCountInString(s) != 11 {
		return errors.New("only 11 digits")
	}

	sumSnils := int64(0)
	for i := 1; i < 10; i++ {
		x, _ := strconv.ParseInt(string(s[9-i]), 10, 64)
		sumSnils += int64(i) * x
	}
	cntrl := snilsControl(sumSnils)
	must, _ := strconv.ParseInt(string(s[9:]), 10, 64)
	if must != cntrl {
		return errors.New("invalid control value ")
	}

	return nil
}

func (r *snilsRule) Error(message string) *snilsRule {
	return &snilsRule{
		message: message,
		code:    207,
	}
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
