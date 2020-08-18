package validation

import (
	"reflect"
)

func MultipleOf(threshold interface{}) *multipleOfRule {
	return &multipleOfRule{
		threshold,
		1106,
	}
}

type multipleOfRule struct {
	threshold interface{}
	code      int
}

func (r *multipleOfRule) Validate(value interface{}) (code int, args []interface{}) {
	rv := reflect.ValueOf(r.threshold)
	args = append(args, rv.Kind())
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		v, err := ToInt(value)
		if err != nil {
			code = r.code
			return
		}
		if v%rv.Int() == 0 {
			code = r.code
			return
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		v, err := ToUint(value)
		if err != nil {
			code = r.code
			return
		}

		if v%rv.Uint() == 0 {
			code = r.code
			return
		}
	default:
		code = r.code
		return
	}

	return
}
