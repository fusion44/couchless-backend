package validator

import "fmt"

// Equals checks for a valid email on a string
func (v *Validator) Equals(fieldA string, valueA interface{}, fieldB string, valueB interface{}) bool {
	if _, ok := v.Errors[fieldA]; ok {
		return false
	}

	if valueA != valueB {
		v.Errors[fieldA] = fmt.Sprintf("%s must equal %s", fieldA, fieldB)
		return false
	}

	return true
}
