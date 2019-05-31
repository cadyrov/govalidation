package bi

import (
	"errors"
	validation "github.com/cadyrov/ozzo-validation"
	"github.com/cadyrov/ozzo-validation/is"
	"strconv"
	"unicode/utf8"
)

var Inn12 = &inn12Rule{message: " is not correct", code: 202}

type inn12Rule struct {
	message string
	code    int
}

func (inn *inn12Rule) Validate(value interface{}) error {
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

	if utf8.RuneCountInString(s) != 12 {
		return errors.New("only 12 digits")
	}

	coefficients11 := []int64{7, 2, 4, 10, 3, 5, 9, 4, 6, 8}
	coefficients12 := []int64{3, 7, 2, 4, 10, 3, 5, 9, 4, 6, 8}
	if !checkInnDigits(s, coefficients11) || !checkInnDigits(s, coefficients12) {
		return errors.New("control sum is invalid")
	}
	return nil
}

func (r *inn12Rule) Error(message string) *inn12Rule {
	return &inn12Rule{
		message: message,
		code:    202,
	}
}
