// Package validation provides configurable and extensible rules for validating data of various types.
package validation

import (
	"fmt"
	"net/http"
	"reflect"
	"strconv"

	"github.com/cadyrov/goerr"
	"github.com/cadyrov/govalidation/verror"
)

type (
	// Validatable is the interface indicating the type implementing it supports data validation.
	Validatable interface {
		// Validate validates the data and returns an error if validation fails.
		Validate() (code int, args []interface{})
	}

	// Rule represents a validation rule.
	Rule interface {
		// Validate validates a value and returns a value if validation fails.
		Validate(value interface{}) (code int, args []interface{})
	}

	// RuleFunc represents a validator function.
	// You may wrap it as a Rule by calling By().
	RuleFunc func(value interface{}) (code int, args []interface{})
)

var (
	// ErrorTag is the struct tag name used to customize the error field name for a struct field.
	ErrorTag = "json"

	// Skip is a special validation rule that indicates all rules following it should be skipped.
	Skip = &skipRule{}

	validatableType = reflect.TypeOf((*Validatable)(nil)).Elem()
)

// Validate validates the given value and returns the validation error, if any.
//
// Validate performs validation using the following steps:
// - validate the value against the rules passed in as parameters
// - if the value is a map and the map values implement `Validatable`, call `Validate` of every map value
// - if the value is a slice or array whose values implement `Validatable`, call `Validate` of every element
func Validate(value interface{}, rules ...Rule) goerr.IError {
	for _, rule := range rules {
		if _, ok := rule.(*skipRule); ok {
			return nil
		}
		if code, args := rule.Validate(value); code != 0 {
			return verror.NewGoErr(code, args...)
		}
	}

	rv := reflect.ValueOf(value)
	if (rv.Kind() == reflect.Ptr || rv.Kind() == reflect.Interface) && rv.IsNil() {
		return nil
	}
	if v, ok := value.(Validatable); ok {
		if code, args := v.Validate(); code != 0 {
			return verror.NewGoErr(code, args...)
		}
		return nil
	}

	switch rv.Kind() {
	case reflect.Map:
		if rv.Type().Elem().Implements(validatableType) {
			return validateMap(rv)
		}
	case reflect.Slice, reflect.Array:
		if rv.Type().Elem().Implements(validatableType) {
			return validateSlice(rv)
		}
	case reflect.Ptr, reflect.Interface:
		return Validate(rv.Elem().Interface())
	}
	return nil
}

// validateMap validates a map of validatable elements
func validateMap(rv reflect.Value) goerr.IError {
	errs := verror.NewErrStack("validationError")
	for _, key := range rv.MapKeys() {
		if mv := rv.MapIndex(key).Interface(); mv != nil {
			if code, args := mv.(Validatable).Validate(); code != 0 {
				e := verror.NewGoErr(code, args)
				e.SetID(fmt.Sprintf("%v", key.Interface()))
				errs.PushDetail(e)
			}
		}
	}
	if len(errs.GetDetails()) > 0 {
		return errs
	}
	return nil
}

// validateMap validates a slice/array of validatable elements
func validateSlice(rv reflect.Value) goerr.IError {
	errs := verror.NewErrStack("validationError")
	l := rv.Len()
	for i := 0; i < l; i++ {
		if ev := rv.Index(i).Interface(); ev != nil {
			if code, args := ev.(Validatable).Validate(); code != 0 {
				e := verror.NewGoErr(code, args)
				e.SetID(strconv.Itoa(i))
				errs.PushDetail(e)
			}
		}
	}
	if len(errs.GetDetails()) > 0 {
		return errs.HTTP(http.StatusBadRequest)
	}
	return nil
}

type skipRule struct{}

func (r *skipRule) Validate(interface{}) (code int, args []interface{}) {
	return 0, nil
}

type inlineRule struct {
	f RuleFunc
}

func (r *inlineRule) Validate(value interface{}) (code int, args []interface{}) {
	return r.f(value)
}

// By wraps a RuleFunc into a Rule.
func By(f RuleFunc) Rule {
	return &inlineRule{f}
}
