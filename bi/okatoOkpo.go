package bi

import (
	"strconv"

	validation "github.com/cadyrov/govalidation"
	"github.com/cadyrov/govalidation/is"
)

var OkatoOkpo = &okatoOkpoRule{code: 2870}

type okatoOkpoRule struct {
	code int
}

func (inn *okatoOkpoRule) Validate(value interface{}) (code int, args []interface{}) {
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
		code = 2870
		return
	}

	if code, args = is.Digit.Validate(s); code != 0 {
		return
	}

	digits := make([]int64, 0)
	controlDigit := int64(0)
	for _, val := range s {
		x, _ := strconv.ParseInt(string(val), 10, 64)
		controlDigit = x
		digits = append(digits, x)
	}

	cn := controlStat(digits, 1, 10)
	if cn == 10 {
		cn = controlStat(digits, 3, 10)
	}
	if cn == 10 {
		cn = 0
	}

	if cn != controlDigit {
		code = 2870
		return
	}
	return
}

func (r *okatoOkpoRule) Error(message string) *snilsRule {
	return &snilsRule{
		code: 2807,
	}
}

func controlStat(digits []int64, a int64, z int64) int64 {
	current := a
	sum := int64(0)
	for i := 0; i < len(digits); i++ {
		if current == 11 {
			current = a
		}
		sum += digits[i] * current
		current++
	}
	return sum % 11
}
