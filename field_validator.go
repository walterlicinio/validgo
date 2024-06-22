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

func (f *FieldValidator) IsString() *FieldValidator {
	if _, ok := f.fieldValue.(string); !ok {
		f.validator.addError(fmt.Errorf("%s must be a string", f.fieldName))
	}
	return f
}

func (f *FieldValidator) MinLength(length int) *FieldValidator {
	if str, ok := f.fieldValue.(string); ok {
		if len(str) < length {
			f.validator.addError(fmt.Errorf("%s must be at least %d characters long", f.fieldName, length))
		}
	}
	return f
}

func (f *FieldValidator) MaxLength(length int) *FieldValidator {
	if str, ok := f.fieldValue.(string); ok {
		if len(str) > length {
			f.validator.addError(fmt.Errorf("%s must be at most %d characters long", f.fieldName, length))
		}
	}
	return f
}

func (f *FieldValidator) IsInt() *FieldValidator {
	if _, ok := f.fieldValue.(int); !ok {
		f.validator.addError(fmt.Errorf("%s must be an integer", f.fieldName))
	}
	return f
}

func (f *FieldValidator) Min(min int) *FieldValidator {
	if val, ok := f.fieldValue.(int); ok {
		if val < min {
			f.validator.addError(fmt.Errorf("%s must be at least %d", f.fieldName, min))
		}
	}
	return f
}

func (f *FieldValidator) Max(max int) *FieldValidator {
	if val, ok := f.fieldValue.(int); ok {
		if val > max {
			f.validator.addError(fmt.Errorf("%s must be at most %d", f.fieldName, max))
		}
	}
	return f
}

func (f *FieldValidator) IsEmail() *FieldValidator {
	if str, ok := f.fieldValue.(string); ok {
		emailRegex := regexp.MustCompile(`^[a-z0-9._%+\-]+@[a-z0-9.\-]+\.[a-z]{2,4}$`)
		if !emailRegex.MatchString(str) {
			f.validator.addError(fmt.Errorf("%s must be a valid email address", f.fieldName))
		}
	}
	return f
}

func (f *FieldValidator) Custom(fn func(any) bool) *FieldValidator {
	if !fn(f.fieldValue) {
		f.validator.addError(fmt.Errorf("%s custom validation is invalid", f.fieldName))
	}
	return f
}

func (f *FieldValidator) Transform(fn func(any) any) *FieldValidator {
	f.fieldValue = fn(f.fieldValue)
	f.updateField()
	return f
}

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
