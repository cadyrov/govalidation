package validation

import (
	"fmt"
	"mime/multipart"
	"reflect"
	"time"
)

type ThresholdRule struct {
	threshold interface{}
	operator  int
	message   string
}

const (
	greaterThan = iota
	greaterEqualThan
	lessThan
	lessEqualThan
)

// Min is a validation rule that checks if a value is greater or equal than the specified value.
// By calling Exclusive, the rule will check if the value is strictly greater than the specified value.
// Note that the value being checked and the threshold value must be of the same type.
// Only int, uint, float and time.Time types are supported.
// An empty value is considered valid. Please use the Required rule to make sure a value is not empty.
func Min(min interface{}) *ThresholdRule {
	return &ThresholdRule{
		threshold: min,
		operator:  greaterEqualThan,
		message:   fmt.Sprintf("must be no less than %v", min),
	}
}

// Max is a validation rule that checks if a value is less or equal than the specified value.
// By calling Exclusive, the rule will check if the value is strictly less than the specified value.
// Note that the value being checked and the threshold value must be of the same type.
// Only int, uint, float and time.Time types are supported.
// An empty value is considered valid. Please use the Required rule to make sure a value is not empty.
func Max(max interface{}) *ThresholdRule {
	return &ThresholdRule{
		threshold: max,
		operator:  lessEqualThan,
		message:   fmt.Sprintf("must be no greater than %v", max),
	}
}

// Exclusive sets the comparison to exclude the boundary value.
func (r *ThresholdRule) Exclusive() *ThresholdRule {
	if r.operator == greaterEqualThan {
		r.operator = greaterThan

		r.message = fmt.Sprintf("must be greater than %v", r.threshold)

		return r
	}

	if r.operator == lessEqualThan {
		r.operator = lessThan

		r.message = fmt.Sprintf("must be less than %v", r.threshold)
	}

	return r
}

// Validate checks if the given value is valid or not.
func (r *ThresholdRule) Validate(value interface{}) (code int, args []interface{}) {
	value, isNil := Indirect(value)
	if isNil || IsEmpty(value) {
		return
	}

	rv := reflect.ValueOf(r.threshold)
	switch rv.Kind() {
	case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		checkInt := true
		slicePtrHeaders, ok := value.([]*multipart.FileHeader)
		if ok {
			checkInt = false
			var sum int64
			for i := range slicePtrHeaders {
				sum += slicePtrHeaders[i].Size
			}
			if r.compareInt(rv.Int(), sum) {
				return
			}
		}
		sliceHeaders, ok := value.([]multipart.FileHeader)
		if ok {
			checkInt = false
			var sum int64
			for i := range sliceHeaders {
				sum += sliceHeaders[i].Size
			}
			if r.compareInt(rv.Int(), sum) {
				return
			}
		}
		ptrHeader, ok := value.(*multipart.FileHeader)
		if ok {
			checkInt = false
			if r.compareInt(rv.Int(), ptrHeader.Size) {
				return
			}
		}
		header, ok := value.(multipart.FileHeader)
		if ok {
			checkInt = false
			if r.compareInt(rv.Int(), header.Size) {
				return
			}
		}
		if checkInt == true {
			v, err := ToInt(value)
			if err != nil {
				code = 1000

				return
			}

			if r.compareInt(rv.Int(), v) {
				return
			}
		}

	case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
		v, err := ToUint(value)
		if err != nil {
			code = 1000
			return
		}
		if r.compareUint(rv.Uint(), v) {
			return
		}

	case reflect.Float32, reflect.Float64:
		v, err := ToFloat(value)
		if err != nil {
			code = 1000
			return
		}
		if r.compareFloat(rv.Float(), v) {
			return
		}

	case reflect.Struct:
		t, ok := r.threshold.(time.Time)
		if !ok {
			code = 1000
			return
		}
		v, ok := value.(time.Time)
		if !ok {
			code = 1000
			return
		}
		if v.IsZero() || r.compareTime(t, v) {
			return
		}

	default:
		code = 1000
		return

	}
	code = 1000
	return
}

func (r *ThresholdRule) compareInt(threshold, value int64) bool {
	switch r.operator {
	case greaterThan:
		return value > threshold
	case greaterEqualThan:
		return value >= threshold
	case lessThan:
		return value < threshold
	default:
		return value <= threshold
	}
}

func (r *ThresholdRule) compareUint(threshold, value uint64) bool {
	switch r.operator {
	case greaterThan:
		return value > threshold
	case greaterEqualThan:
		return value >= threshold
	case lessThan:
		return value < threshold
	default:
		return value <= threshold
	}
}

func (r *ThresholdRule) compareFloat(threshold, value float64) bool {
	switch r.operator {
	case greaterThan:
		return value > threshold
	case greaterEqualThan:
		return value >= threshold
	case lessThan:
		return value < threshold
	default:
		return value <= threshold
	}
}

func (r *ThresholdRule) compareTime(threshold, value time.Time) bool {
	switch r.operator {
	case greaterThan:
		return value.After(threshold)
	case greaterEqualThan:
		return value.After(threshold) || value.Equal(threshold)
	case lessThan:
		return value.Before(threshold)
	default:
		return value.Before(threshold) || value.Equal(threshold)
	}
}
