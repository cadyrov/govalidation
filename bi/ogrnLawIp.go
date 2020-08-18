package bi

import (
	validation "github.com/cadyrov/govalidation"
)

var ORGNLawIp = &ogrnLawIpRule{code: 2860}

type ogrnLawIpRule struct {
	message string
	code    int
}

func (oilp *ogrnLawIpRule) Validate(value interface{}) (code int, args []interface{}) {
	value, isNil := validation.Indirect(value)
	if isNil || validation.IsEmpty(value) {
		return
	}
	errLaw, _ := OGRNLaw.Validate(value)
	errIp, _ := OPGNIp.Validate(value)

	if errLaw != 0 && errIp != 0 {
		code = oilp.code
		return
	}
	return
}

func (r *ogrnLawIpRule) Error(message string) *ogrnLawIpRule {
	return &ogrnLawIpRule{
		message: message,
		code:    2806,
	}
}
