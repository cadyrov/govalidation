package validation

import (
	"reflect"
	"strings"

	"github.com/cadyrov/goerr"
	"github.com/cadyrov/govalidation/verror"
)

var (
	// ErrStructPointer is the error that a struct being validated is not specified as a pointer.
	ErrStructPointer = verror.NewGoErr(1001)
)

type (
	// ErrFieldPointer is the error that a field is not specified as a pointer.
	ErrFieldPointer int

	// ErrFieldNotFound is the error that a field cannot be found in the struct.
	ErrFieldNotFound int

	// FieldRules represents a rule set associated with a struct field.
	FieldRules struct {
		fieldPtr interface{}
		rules    []Rule
	}
)

// Error returns the error string of ErrFieldPointer.

func (e ErrFieldPointer) GetCode() int {
	return 1002
}

func (e ErrFieldNotFound) GetCode() int {
	return 1003
}

func ValidateStruct(structPtr interface{}, fields ...*FieldRules) goerr.IError {
	value := reflect.ValueOf(structPtr)
	if value.Kind() != reflect.Ptr || !value.IsNil() && value.Elem().Kind() != reflect.Struct {
		return ErrStructPointer
	}
	if value.IsNil() {
		// treat a nil struct pointer as valid
		return nil
	}
	value = value.Elem()

	errs := verror.NewErrStack("validation_error")

	for _, fr := range fields {
		fv := reflect.ValueOf(fr.fieldPtr)
		if fv.Kind() != reflect.Ptr {
			return verror.NewGoErr(1002)
		}
		ft := findStructField(value, fv)
		if ft == nil {
			return verror.NewGoErr(1003)
		}
		if err := Validate(fv.Elem().Interface(), fr.rules...); err != nil {
			if ft.Anonymous {
				// merge errors from anonymous struct field
				if es, ok := err.(verror.ErrStack); ok {
					details := es.GetDetails()
					for i := range details {
						errs.PushDetail(details[i])
					}
					continue
				}
			}
			err.SetID(getErrorFieldName(ft))
			errs.PushDetail(err)
		}
	}

	if len(errs.GetDetails()) > 0 {
		return errs
	}
	return nil
}

// Field specifies a struct field and the corresponding validation rules.
// The struct field must be specified as a pointer to it.
func Field(fieldPtr interface{}, rules ...Rule) *FieldRules {
	return &FieldRules{
		fieldPtr: fieldPtr,
		rules:    rules,
	}
}

// findStructField looks for a field in the given struct.
// The field being looked for should be a pointer to the actual struct field.
// If found, the field info will be returned. Otherwise, nil will be returned.
func findStructField(structValue reflect.Value, fieldValue reflect.Value) *reflect.StructField {
	ptr := fieldValue.Pointer()
	for i := structValue.NumField() - 1; i >= 0; i-- {
		sf := structValue.Type().Field(i)
		if ptr == structValue.Field(i).UnsafeAddr() {
			// do additional type comparison because it's possible that the address of
			// an embedded struct is the same as the first field of the embedded struct
			if sf.Type == fieldValue.Elem().Type() {
				return &sf
			}
		}
		if sf.Anonymous {
			// delve into anonymous struct to look for the field
			fi := structValue.Field(i)
			if sf.Type.Kind() == reflect.Ptr {
				fi = fi.Elem()
			}
			if fi.Kind() == reflect.Struct {
				if f := findStructField(fi, fieldValue); f != nil {
					return f
				}
			}
		}
	}
	return nil
}

// getErrorFieldName returns the name that should be used to represent the validation error of a struct field.
func getErrorFieldName(f *reflect.StructField) string {
	if tag := f.Tag.Get(ErrorTag); tag != "" {
		if cps := strings.SplitN(tag, ",", 2); cps[0] != "" {
			return cps[0]
		}
	}
	return f.Name
}
