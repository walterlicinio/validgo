package validgo

import (
	"reflect"
	"strings"
)

type Validator struct {
	data   any
	errors []error
}

// New creates a new Validator instance
func New(data any) *Validator {
	return &Validator{data: data}
}

// Field returns a FieldValidator for a specific field in the data structure
func (v *Validator) Field(name string) *FieldValidator {
	value := v.getField(name)
	return &FieldValidator{validator: v, fieldName: name, fieldValue: value}
}

// HasErrors returns true if any errors occurred during validation
func (v *Validator) HasErrors() bool {
	return len(v.errors) > 0
}

// Errors returns a list of errors that occurred during validation
func (v *Validator) Errors() []error {
	return v.errors
}

// getField retrieves a field from the data structure
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

func (v *Validator) addError(err error) {
	v.errors = append(v.errors, err)
}
