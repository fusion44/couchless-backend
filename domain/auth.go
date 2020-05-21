package domain

import (
	"context"
	"errors"
	"log"

	"github.com/fusion44/ll-backend/graph/model"
	"github.com/fusion44/ll-backend/middleware"
	"github.com/fusion44/ll-backend/validator"

	gcontext "github.com/fusion44/ll-backend/context"
)

// Register registers a new user
func (d *Domain) Register(ctx context.Context, input model.RegisterInput) (*model.AuthResponse, error) {
	_, err := d.UsersRepo.GetUserByEmail(input.Email)
	if err == nil {
		return nil, errors.New("Email is already in use")
	}

	_, err = d.UsersRepo.GetUserByUsername(input.Username)
	if err == nil {
		return nil, errors.New("Username is already in use")
	}

	user := &model.User{
		Username: input.Username,
		Email:    input.Email,
	}

	errMsg := errors.New("Something went wrong")

	err = user.HashPassword(input.Password)
	if err != nil {
		log.Printf("Password hashing failed")
		return nil, errMsg
	}

	tx, err := d.UsersRepo.DB.Begin()
	if err != nil {
		log.Printf("Error creating transaction: %v", err)
		return nil, errMsg
	}

	defer tx.Rollback()

	if _, err := d.UsersRepo.CreateUser(tx, user); err != nil {
		log.Printf("Error creating user: %v", err)
		return nil, err // tx.Rollback is called
	}

	if err := tx.Commit(); err != nil {
		log.Printf("Error while commiting: %v", err)
		return nil, err // tx.Rollback is called
	}

	cfg, err := gcontext.GetConfigFromContext(ctx)
	if err != nil {
		log.Printf("No config in context: %v", err)
		return nil, err
	}

	token, err := user.GenToken(cfg)
	if err != nil {
		log.Printf("Error while generating token: %v", err)
		return nil, err
	}

	return &model.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil

	// Rollback is called, but on an empty transaction,
	// as it was committed above
}

// Login logs a user in
func (d *Domain) Login(ctx context.Context, input model.LoginInput) (*model.AuthResponse, error) {
	valid, _ := input.Validate()
	if !valid {
		return nil, ErrBadCredentials
	}

	var user *model.User
	var err error

	v := validator.New()
	if v.IsEmail("email", input.Username) {
		user, err = d.UsersRepo.GetUserByEmail(input.Username)
	} else if v.MinLength("usernameLength", input.Username, 2) {
		user, err = d.UsersRepo.GetUserByUsername(input.Username)
	} else {
		return nil, ErrBadCredentials
	}

	if err != nil {
		// Input valid, but no user found
		log.Printf("No user found with username: %s", input.Username)
		return nil, ErrBadCredentials
	}

	err = user.ComparePassword(input.Password)
	if err != nil {
		return nil, ErrBadCredentials
	}

	cfg, err := gcontext.GetConfigFromContext(ctx)
	if err != nil {
		log.Printf("No config in context: %v", err)
		return nil, err
	}

	token, err := user.GenToken(cfg)
	if err != nil {
		return nil, ErrBadCredentials
	}

	return &model.AuthResponse{
		AuthToken: token,
		User:      user,
	}, nil
}

// GetUserByID retrieves a user object for the given id
func (d *Domain) GetUserByID(ctx context.Context, id string) (*model.User, error) {
	u, err := middleware.GetCurrentUserFromContext(ctx)
	if err != nil {
		return nil, ErrUnauthenticated
	}

	if u.ID == id {
		return u, nil
	}

	a, err := d.UsersRepo.GetUserByID(id)

	if err != nil {
		return nil, err
	}

	return a, nil
}
