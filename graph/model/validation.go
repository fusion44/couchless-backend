package model

import "github.com/fusion44/ll-backend/validator"

// Validate all register input fields
func (r RegisterInput) Validate() (bool, map[string]string) {
	v := validator.New()

	v.Require("email", r.Email)
	v.IsEmail("email", r.Email)

	v.Require("password", r.Password)
	v.MinLength("password", r.Password, 6)

	v.Require("confirmPassword", r.ConfirmPassword)
	v.MinLength("confirmPassword", r.ConfirmPassword, 6)

	v.Equals("password", r.Password, "confirmPassword", r.ConfirmPassword)

	v.Require("username", r.Username)
	v.MinLength("username", r.Username, 2)

	return v.IsValid(), v.Errors
}

// Validate all login input fields
func (l LoginInput) Validate() (bool, map[string]string) {
	v := validator.New()

	v.Require("password", l.Password)

	return v.IsValid(), v.Errors
}
