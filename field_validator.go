package validgo

import (
	"fmt"
	"reflect"
	"regexp"
	"strings"
)

type FieldValidator struct {
	validator  *Validator
	fieldName  string
	fieldValue any
}

// IsString checks if the field is a string
func (f *FieldValidator) IsString() *FieldValidator {
	if _, ok := f.fieldValue.(string); !ok {
		f.validator.addError(fmt.Errorf("%s must be a string", f.fieldName))
	}
	return f
}

// MinLength checks if the field is at least the specified length
func (f *FieldValidator) MinLength(length int) *FieldValidator {
	if str, ok := f.fieldValue.(string); ok {
		if len(str) < length {
			f.validator.addError(fmt.Errorf("%s must be at least %d characters long", f.fieldName, length))
		}
	}
	return f
}

// MaxLength checks if the field is at most the specified length
func (f *FieldValidator) MaxLength(length int) *FieldValidator {
	if str, ok := f.fieldValue.(string); ok {
		if len(str) > length {
			f.validator.addError(fmt.Errorf("%s must be at most %d characters long", f.fieldName, length))
		}
	}
	return f
}

// IsEmail checks if the field is a valid email address
func (f *FieldValidator) IsEmail() *FieldValidator {
	if str, ok := f.fieldValue.(string); ok {
		emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
		if !emailRegex.MatchString(str) {
			f.validator.addError(fmt.Errorf("%s must be a valid email address", f.fieldName))
		}
	}
	return f
}

// Custom allows for custom validation logic
// The function should return true if the validation is successful
func (f *FieldValidator) Custom(fn func(any) bool) *FieldValidator {
	if !fn(f.fieldValue) {
		f.validator.addError(fmt.Errorf("%s custom validation is invalid", f.fieldName))
	}
	return f
}

// Transform allows for custom transformation logic
// The function should return the transformed value
func (f *FieldValidator) Transform(fn func(any) any) *FieldValidator {
	f.fieldValue = fn(f.fieldValue)
	f.updateField()
	return f
}

// updateField updates the field in the data structure
func (f *FieldValidator) updateField() {
	current := reflect.ValueOf(f.validator.data)
	parts := strings.Split(f.fieldName, ".")
	for i, part := range parts {
		if current.Kind() == reflect.Ptr {
			current = current.Elem()
		}
		if i == len(parts)-1 {
			current.FieldByName(part).Set(reflect.ValueOf(f.fieldValue))
		} else {
			current = current.FieldByName(part)
		}
	}
}
