package bi

import (
	"errors"
	validation "github.com/cadyrov/ozzo-validation"
	"github.com/cadyrov/ozzo-validation/is"
	"strconv"
	"unicode/utf8"
)

var Inn10 = &inn10Rule{message: " is not correct", code: 201}

type inn10Rule struct {
	message string
	code    int
}

func (inn *inn10Rule) Validate(value interface{}) error {
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

	if utf8.RuneCountInString(s) != 10 {
		return errors.New("only 10 digits")
	}

	coefficients := []int64{2, 4, 10, 3, 5, 9, 4, 6, 8}

	if !checkInnDigits(s, coefficients) {
		return errors.New("control sum is invalid")
	}
	return nil
}

func (r *inn10Rule) Error(message string) *inn10Rule {
	return &inn10Rule{
		message: message,
		code:    201,
	}
}
