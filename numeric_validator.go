package validgo

import "fmt"

// IsInt checks if the field is an integer
func (f *FieldValidator) IsInt() *FieldValidator {
	if _, ok := f.fieldValue.(int); !ok {
		f.validator.addError(fmt.Errorf("%s must be an integer", f.fieldName))
	}
	return f
}

// Min checks if the field is at least the specified value
func (f *FieldValidator) Min(min int) *FieldValidator {
	if val, ok := f.fieldValue.(int); ok {
		if val < min {
			f.validator.addError(fmt.Errorf("%s must be at least %d", f.fieldName, min))
		}
	}
	return f
}

// Max checks if the field is at most the specified value
func (f *FieldValidator) Max(max int) *FieldValidator {
	if val, ok := f.fieldValue.(int); ok {
		if val > max {
			f.validator.addError(fmt.Errorf("%s must be at most %d", f.fieldName, max))
		}
	}
	return f
}

// IsFloat checks if the field is a float
func (f *FieldValidator) IsFloat() *FieldValidator {
	if _, ok := f.fieldValue.(float64); !ok {
		f.validator.addError(fmt.Errorf("%s must be a float", f.fieldName))
	}
	return f
}

// MinFloat checks if the field is at least the specified value
func (f *FieldValidator) MinFloat(min float64) *FieldValidator {
	if val, ok := f.fieldValue.(float64); ok {
		if val < min {
			f.validator.addError(fmt.Errorf("%s must be at least %f", f.fieldName, min))
		}
	}
	return f
}

// MaxFloat checks if the field is at most the specified value
func (f *FieldValidator) MaxFloat(max float64) *FieldValidator {
	if val, ok := f.fieldValue.(float64); ok {
		if val > max {
			f.validator.addError(fmt.Errorf("%s must be at most %f", f.fieldName, max))
		}
	}
	return f
}

// IsPositive checks if the field is a positive number
// Also checks if the field is a positive integer or float
func (f *FieldValidator) IsPositive() *FieldValidator {
	switch v := f.fieldValue.(type) {
	case int:
		if v <= 0 {
			f.validator.addError(fmt.Errorf("%s must be a positive integer", f.fieldName))
		}
	case float64:
		if v <= 0 {
			f.validator.addError(fmt.Errorf("%s must be a positive float", f.fieldName))
		}
	default:
		f.validator.addError(fmt.Errorf("%s must be a positive number", f.fieldName))
	}
	return f
}
