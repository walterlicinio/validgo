package validgo

import (
	"fmt"
	"regexp"
)

func (f *FieldValidator) Matches(pattern string) *FieldValidator {
	if str, ok := f.fieldValue.(string); ok {
		matched, _ := regexp.MatchString(pattern, str)
		if !matched {
			f.validator.addError(fmt.Errorf("%s must match pattern %s", f.fieldName, pattern))
		}
	}
	return f
}

func (f *FieldValidator) IsAlpha() *FieldValidator {
	return f.Matches("^[a-zA-Z]+$")
}

func (f *FieldValidator) IsAlphanumeric() *FieldValidator {
	return f.Matches("^[a-zA-Z0-9]+$")
}

func (f *FieldValidator) IsNumeric() *FieldValidator {
	return f.Matches("^[0-9]+$")
}
