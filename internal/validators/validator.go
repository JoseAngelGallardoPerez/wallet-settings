package validators

import (
	"reflect"

	ut "github.com/go-playground/universal-translator"
)

// Validator interface provide methods to validate structures
type Validator interface {
	// Validate a structure
	Struct(current interface{}) error
}

type ValidationError struct {
	Err error
	// Original structure contains slice of Setting structures.
	// We should store original indexes of slice items to build correct Error source in an HTTP response.
	// key - name of the new struct and its field. For example "AutoLogout.Status"
	// value - index of field in original data structure
	Indexes map[string]int
}

// implementation of github.com/go-playground/validator/v10.FieldError interface
type FieldError struct {
	tag         string
	ns          string
	field       string
	structfield string
	value       interface{}
	param       string
	kind        reflect.Kind
	typ         reflect.Type
}

func (s FieldError) Tag() string {
	return s.tag
}

func (s FieldError) ActualTag() string {
	return s.tag
}

func (s FieldError) Namespace() string {
	return s.ns
}

func (s FieldError) StructNamespace() string {
	return s.ns
}

func (s FieldError) Field() string {
	return s.field
}

func (s FieldError) StructField() string {
	return s.structfield
}

func (s FieldError) Value() interface{} {
	return s.value
}

func (s FieldError) Param() string {
	return s.param
}

func (s FieldError) Kind() reflect.Kind {
	return s.kind
}

func (s FieldError) Type() reflect.Type {
	return s.typ
}

func (s FieldError) Translate(ut ut.Translator) string {
	return ""
}
