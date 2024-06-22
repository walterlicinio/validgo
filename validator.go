package validgo

import (
	"reflect"
	"strings"
)

type Validator struct {
	data   any
	errors []error
}

func New(data any) *Validator {
	return &Validator{data: data}
}

func (v *Validator) Field(name string) *FieldValidator {
	value := v.getField(name)
	return &FieldValidator{validator: v, fieldName: name, fieldValue: value}
}

func (v *Validator) HasErrors() bool {
	return len(v.errors) > 0
}

func (v *Validator) Errors() []error {
	return v.errors
}

func (v *Validator) getField(name string) any {
	current := reflect.ValueOf(v.data)
	for _, part := range strings.Split(name, ".") {
		if current.Kind() == reflect.Ptr {
			current = current.Elem()
		}
		if current.Kind() != reflect.Struct {
			return nil
		}
		current = current.FieldByName(part)
		if !current.IsValid() {
			return nil
		}
	}
	return current.Interface()
}

func (v Validator) addError(err error) {
	v.errors = append(v.errors, err)
}
