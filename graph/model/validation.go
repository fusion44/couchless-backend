package model

import "github.com/fusion44/couchless-backend/validator"

const (
	minUsernameLength = 2
	minPasswordLength = 8
)

// Validate all register input fields
func (r RegisterInput) Validate() (bool, map[string]string) {
	v := validator.New()

	v.Require("email", r.Email)
	v.IsEmail("email", r.Email)

	v.Require("password", r.Password)
	v.MinLength("password", r.Password, minPasswordLength)

	v.Require("confirmPassword", r.ConfirmPassword)
	v.MinLength("confirmPassword", r.ConfirmPassword, minPasswordLength)

	v.Equals("password", r.Password, "confirmPassword", r.ConfirmPassword)

	v.Require("username", r.Username)
	v.MinLength("username", r.Username, minUsernameLength)

	return v.IsValid(), v.Errors
}

// Validate all login input fields
func (l LoginInput) Validate() (bool, map[string]string) {
	v := validator.New()

	v.Require("username", l.Username)
	v.MinLength("username", l.Username, minUsernameLength)

	v.Require("password", l.Password)
	v.MinLength("password", l.Password, minPasswordLength)

	return v.IsValid(), v.Errors
}
