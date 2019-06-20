// Copyright 2016 Qiang Xue. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

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
		if errs, ok := es[key].(ExternalError); ok {
			s += fmt.Sprintf("%v: %v, ErrCode: %v", key, errs.Error(), errs.GetCode())
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
		if errs, ok := es[key].(ExternalError); ok {
			s += fmt.Sprintf("%v: %v, ErrCode: %v", key, errs.Error(), errs.GetCode())
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
			errs[key] = err.Error()
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
