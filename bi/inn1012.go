package bi

import (
	"errors"
	validation "github.com/cadyrov/ozzo-validation"

	"strconv"
)

func checkInnDigits(inn string, coefficients []int64) bool {

	n := int64(0)
	innSl := make([]int64, 12)

	for i, val := range inn {
		innSl[i], _ = strconv.ParseInt(string(val), 10, 64)
	}
	for i, k := range coefficients {
		n += k * innSl[i]
	}

	n = n % 11 % 10

	if n == innSl[len(coefficients)] {
		return true
	}
	return false
}

var Inn1012 = &inn1012Rule{message: " is not correct", code: 203}

type inn1012Rule struct {
	message string
	code    int
}

func (inn *inn1012Rule) Validate(value interface{}) error {
	value, isNil := validation.Indirect(value)
	if isNil || validation.IsEmpty(value) {
		return nil
	}
	err10 := Validate(value)
	err12 := Validate(value)

	if err10 != nil && err12 != nil {
		return errors.New(" 10 or 12 digits and control value ")
	}
	return nil
}

func (r *inn1012Rule) Error(message string) *inn1012Rule {
	return &inn1012Rule{
		message: message,
		code:    203,
	}
}
