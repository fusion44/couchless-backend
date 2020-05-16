package domain

import (
	"errors"

	"github.com/fusion44/ll-backend/db/repositories"
	"github.com/fusion44/ll-backend/graph/model"
)

// App errors
var (
	ErrBadCredentials  = errors.New("Login credentials not valid")
	ErrUnauthenticated = errors.New("Unauthenticated")
	ErrUnauthorized    = errors.New("Unauthorized")
	ErrInternalServer  = errors.New("Internal server error")
	ErrInvalidInput    = errors.New("Input not valid")
)

// Domain contains all business logic
type Domain struct {
	UsersRepo    repositories.UsersRepository
	ActivityRepo repositories.ActivitiesRepository
}

// NewDomain creates a new Domain instance
func NewDomain(usersRepo repositories.UsersRepository, activityRepo repositories.ActivitiesRepository) *Domain {
	return &Domain{UsersRepo: usersRepo, ActivityRepo: activityRepo}
}

// Ownable makes an object ownable by an user
type Ownable interface {
	IsOwner(user *model.User) bool
}
