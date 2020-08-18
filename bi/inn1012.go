package bi

import (
	validation "github.com/cadyrov/govalidation"

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

var Inn1012 = &inn1012Rule{code: 2803}

type inn1012Rule struct {
	code int
}

func (inn *inn1012Rule) Validate(value interface{}) (code int, args []interface{}) {
	value, isNil := validation.Indirect(value)
	if isNil || validation.IsEmpty(value) {
		return
	}
	err10, _ := Inn10.Validate(value)
	err12, _ := Inn12.Validate(value)

	if err10 != 0 && err12 != 0 {
		code = inn.code
		return
	}
	return
}

func (r *inn1012Rule) Error(message string) *inn1012Rule {
	return &inn1012Rule{
		code: 2830,
	}
}
