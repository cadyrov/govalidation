package validation

import (
	"database/sql"
	"testing"

	"reflect"

	"github.com/stretchr/testify/assert"
)

func validateMe(s string) bool {
	return s == "me"
}

func TestNewStringRule(t *testing.T) {
	v := NewStringRule(validateMe, "abc", 3035)
	assert.NotNil(t, v.validate)
	assert.Equal(t, "abc", v.message)
}

func TestStringValidator_Error(t *testing.T) {
	v := NewStringRule(validateMe, "abc", 3035)
	assert.Equal(t, "abc", v.message)
	v2 := v.Error("correct", 3030)
	assert.Equal(t, "correct", v2.message)
	assert.Equal(t, "abc", v.message)
}

func TestStringValidator_Validate(t *testing.T) {
	v := NewStringRule(validateMe, "wrong", 3035)

	value := "me"

	err := v.Validate(value)
	assert.Nil(t, err)

	err = v.Validate(&value)
	assert.Nil(t, err)

	value = ""

	err = v.Validate(value)
	assert.Nil(t, err)

	err = v.Validate(&value)
	assert.Nil(t, err)

	nullValue := sql.NullString{String: "me", Valid: true}
	err = v.Validate(nullValue)
	assert.Nil(t, err)

	nullValue = sql.NullString{String: "", Valid: true}
	err = v.Validate(nullValue)
	assert.Nil(t, err)

	var s *string
	err = v.Validate(s)
	assert.Nil(t, err)

	err = v.Validate("not me")
	if assert.NotNil(t, err) {
		assert.Equal(t, "wrong, ErrCode: 3035", err.Error())
	}

	err = v.Validate(100)
	if assert.NotNil(t, err) {
		assert.NotEqual(t, "wrong, ErrCode: 3035", err.Error())
	}

	v2 := v.Error("Wrong!", 3035)
	err = v2.Validate("not me")
	if assert.NotNil(t, err) {
		assert.Equal(t, "Wrong!, ErrCode: 3035", err.Error())
	}
}

func TestGetErrorFieldName(t *testing.T) {
	type A struct {
		T0 string
		T1 string `json:"t1"`
		T2 string `json:"t2,omitempty"`
		T3 string `json:",omitempty"`
		T4 string `json:"t4,x1,omitempty"`
	}
	tests := []struct {
		tag   string
		field string
		name  string
	}{
		{"t1", "T0", "T0"},
		{"t2", "T1", "t1"},
		{"t3", "T2", "t2"},
		{"t4", "T3", "T3"},
		{"t5", "T4", "t4"},
	}
	a := reflect.TypeOf(A{})
	for _, test := range tests {
		field, _ := a.FieldByName(test.field)
		assert.Equal(t, test.name, getErrorFieldName(&field), test.tag)
	}
}
