package bi

import (
	"errors"
	validation "github.com/cadyrov/govalidation"
	"github.com/cadyrov/govalidation/is"
	"strconv"
)

var OkatoOkpo = &okatoOkpoRule{message: " is not correct", code: 208}

type okatoOkpoRule struct {
	message string
	code    int
}

func (inn *okatoOkpoRule) Validate(value interface{}) error {
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
		return errors.New("invalid control value ")
	}
	return nil
}

func (r *okatoOkpoRule) Error(message string) *snilsRule {
	return &snilsRule{
		message: message,
		code:    207,
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
