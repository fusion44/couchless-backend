package validator

// Validation interface for all validators
type Validation interface {
	Validate() (bool, map[string]string)
}

// Validator implements various data validators
type Validator struct {
	Errors map[string]string
}

// New makes a new Validator instance
func New() *Validator {
	return &Validator{Errors: make(map[string]string)}
}

// IsValid returns true if there are no errors found
func (v *Validator) IsValid() bool {
	return len(v.Errors) == 0
}
