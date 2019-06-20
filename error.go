package validation

import (
	"encoding/json"
	"errors"
	"fmt"
	"sort"
)

type (
	ExternalError interface {
		error
		ExternalError() error
		GetCode() int
	}

	jsonError struct {
		ErrCode int    `json:"errCode"`
		Err     string `json:"error"`
	}
	externalError struct {
		code int
		error
	}

	// Errors represents the validation errors that are indexed by struct field names, map or slice keys.
	Errors map[string]ExternalError
)

func (e *externalError) ExternalError() error {
	return e.error
}

func (e *externalError) Error() string {
	if e.GetCode() == 0 {
		return fmt.Sprintf("%v", e.error)
	}
	return fmt.Sprintf("%v, ErrCode: %v", e.error, e.GetCode())
}

// NewExternalError wraps a given error into an InternalError.
func NewExternalError(err error, code int) ExternalError {
	return &externalError{error: err, code: code}
}

// InternalError returns the actual error that it wraps around.
func (e *externalError) GetCode() int {
	return e.code
}

// Error returns the error string of Errors.
func (es Errors) Error() string {
	if len(es) == 0 {
		return ""
	}

	keys := []string{}
	for key := range es {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	s := ""
	for i, key := range keys {
		if i > 0 {
			s += "; "
		}
		if errs, ok := es[key].(Errors); ok {
			s += fmt.Sprintf("%v: (%v)", key, errs.Error())
		} else {
			s += fmt.Sprintf("%v: %v", key, es[key].Error())
		}
	}
	return s + "."
}

func (es Errors) ExternalError() error {
	if len(es) == 0 {
		return nil
	}

	keys := []string{}
	for key := range es {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	s := ""
	for i, key := range keys {
		if i > 0 {
			s += "; "
		}
		if errs, ok := es[key].(Errors); ok {
			s += fmt.Sprintf("%v: (%v)", key, errs)
		} else {
			s += fmt.Sprintf("%v: %v", key, es[key].Error())
		}
	}
	return errors.New(s)
}

func (es Errors) GetCode() int {
	if len(es) == 0 {
		return 0
	}

	keys := []string{}
	for key := range es {
		keys = append(keys, key)
	}
	sort.Strings(keys)

	s := ""
	for i, key := range keys {
		if i > 0 {
			s += "; "
		}
		if errs, ok := es[key].(ExternalError); ok {
			s += fmt.Sprintf("%v: (%v)", key, errs.GetCode())
		} else {
			s += fmt.Sprintf("%v: %v", key, es[key].Error())
		}
	}
	return 0
}

// MarshalJSON converts the Errors into a valid JSON.
func (es Errors) MarshalJSON() ([]byte, error) {
	errs := map[string]interface{}{}
	for key, err := range es {

		if ms, ok := err.(json.Marshaler); ok {
			errs[key] = ms
		} else {
			errs[key] = jsonError{Err: err.ExternalError().Error(), ErrCode: err.GetCode()}
		}
	}
	return json.Marshal(errs)
}

// Filter removes all nils from Errors and returns back the updated Errors as an error.
// If the length of Errors becomes 0, it will return nil.
func (es Errors) Filter() error {
	for key, value := range es {
		if value == nil {
			delete(es, key)
		}
	}
	if len(es) == 0 {
		return nil
	}
	return es
}
