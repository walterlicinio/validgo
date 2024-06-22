package validgo

import (
	"fmt"
	"regexp"
)

// Matches checks if the field matches the specified regex pattern
func (f *FieldValidator) Matches(pattern string) *FieldValidator {
	if str, ok := f.fieldValue.(string); ok {
		matched, _ := regexp.MatchString(pattern, str)
		if !matched {
			f.validator.addError(fmt.Errorf("%s must match pattern %s", f.fieldName, pattern))
		}
	}
	return f
}

// IsAlpha checks if the field contains only alphabetic characters
func (f *FieldValidator) IsAlpha() *FieldValidator {
	return f.Matches("^[a-zA-Z]+$")
}

// IsAlphanumeric checks if the field contains only alphanumeric characters
func (f *FieldValidator) IsAlphanumeric() *FieldValidator {
	return f.Matches("^[a-zA-Z0-9]+$")
}

// IsNumeric checks if the field contains only numeric characters
func (f *FieldValidator) IsNumeric() *FieldValidator {
	return f.Matches("^[0-9]+$")
}
