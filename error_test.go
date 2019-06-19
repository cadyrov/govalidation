// Copyright 2016 Qiang Xue. All rights reserved.
// Use of this source code is governed by a MIT-style
// license that can be found in the LICENSE file.

package validation

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewInternalError(t *testing.T) {
	err := NewInternalError(errors.New("abc"))
	if assert.NotNil(t, err.InternalError()) {
		assert.Equal(t, "abc", err.InternalError().Error())
	}
}

func TestErrors_Error(t *testing.T) {
	errs := Errors{
		"B": NewExternalError(errors.New("B1"), 1),
		"C": NewExternalError(errors.New("C1"), 2),
		"A": NewExternalError(errors.New("A1"), 3),
	}
	assert.Equal(t, "A: A1; B: B1; C: C1.", errs.Error())

	errs = Errors{
		"B": NewExternalError(errors.New("B1"), 1),
	}
	assert.Equal(t, "B: B1.", errs.Error())

	errs = Errors{}
	assert.Equal(t, "", errs.Error())
}

func TestErrors_MarshalMessage(t *testing.T) {
	errs := Errors{
		"A": NewExternalError(errors.New("A1"), 1),
		"B": Errors{
			"2": NewExternalError(errors.New("B1"), 2),
		},
	}
	errsJSON, err := errs.MarshalJSON()
	assert.Nil(t, err)
	assert.Equal(t, "{\"A\":\"A1\",\"B\":{\"2\":\"B1\"}}", string(errsJSON))
}

func TestErrors_Filter(t *testing.T) {
	errs := Errors{
		"B": NewExternalError(errors.New("B1"), 2),
		"C": nil,
		"A": NewExternalError(errors.New("A1"), 1),
	}
	err := errs.Filter()
	assert.Equal(t, 2, len(errs))
	if assert.NotNil(t, err) {
		assert.Equal(t, "A: A1; B: B1.", err.Error())
	}

	errs = Errors{}
	assert.Nil(t, errs.Filter())

	errs = Errors{
		"B": nil,
		"C": nil,
	}
	assert.Nil(t, errs.Filter())
}
