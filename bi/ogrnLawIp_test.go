package bi

import (
	validation "github.com/cadyrov/govalidation"
	"testing"
)

func TestOgrnLawIp(t *testing.T) {
	slInt := make([]interface{}, 0)
	slInt = append(slInt, "csacsa", "12", "dcdcd", "12", "csa", 44545455, "1177746126040")

	for i, val := range slInt {
		vl := val
		err := validation.Validate(vl, ORGNLawIp)
		if i == (len(slInt) - 1) {
			if err != nil {
				t.Fatal(err.Error())
			}
		} else {
			if err == nil {
				t.Fatal("mast have error", val)
			}
		}
	}

}
