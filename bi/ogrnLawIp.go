package bi

import (
	"errors"
	validation "github.com/cadyrov/govalidation"
)

var ORGNLawIp = &ogrnLawIpRule{message: validation.MsgByCode(2806), code: 2806}

type ogrnLawIpRule struct {
	message string
	code    int
}

func (inn *ogrnLawIpRule) Validate(value interface{}) validation.ExternalError {
	value, isNil := validation.Indirect(value)
	if isNil || validation.IsEmpty(value) {
		return nil
	}
	errLaw := OGRNLaw.Validate(value)
	errIp := OPGNIp.Validate(value)

	if errLaw != nil && errIp != nil {
		return validation.NewExternalError(errors.New(" 13 or 15 digits and control value "), 2806)
	}
	return nil
}

func (r *ogrnLawIpRule) Error(message string) *ogrnLawIpRule {
	return &ogrnLawIpRule{
		message: message,
		code:    2806,
	}
}
