package bi

import (
	validation "github.com/cadyrov/govalidation"
	"testing"
)

func TestInn1012(t *testing.T) {
	slInt := make([]interface{}, 0)
	slInt = append(slInt, "csacsa", "12", "dcdcd", "12", "csa", 44545455, "051204727800")

	for i, val := range slInt {
		vl := val
		err := validation.Validate(vl, Inn1012)
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
