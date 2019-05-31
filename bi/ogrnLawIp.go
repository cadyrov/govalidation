package bi

import (
	"errors"
	validation "github.com/cadyrov/govalidation"
)

var ORGNLawIp = &ogrnLawIpRule{message: " is not correct", code: 206}

type ogrnLawIpRule struct {
	message string
	code    int
}

func (inn *ogrnLawIpRule) Validate(value interface{}) error {
	value, isNil := validation.Indirect(value)
	if isNil || validation.IsEmpty(value) {
		return nil
	}
	errLaw := Validate(value)
	errIp := Validate(value)

	if errLaw != nil && errIp != nil {
		return errors.New(" 13 or 15 digits and control value ")
	}
	return nil
}

func (r *ogrnLawIpRule) Error(message string) *ogrnLawIpRule {
	return &ogrnLawIpRule{
		message: message,
		code:    206,
	}
}
