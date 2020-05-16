package validator

import "fmt"

// MinLength checks a minimum string length
func (v *Validator) MinLength(field, value string, minLength int) bool {
	if _, ok := v.Errors[field]; ok {
		return false
	}

	if len(value) < minLength {
		v.Errors[field] = fmt.Sprintf("%s must be at least %d characters long", field, minLength)
		return false
	}

	return true
}
