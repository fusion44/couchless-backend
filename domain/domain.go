package domain

import (
	"errors"

	"github.com/fusion44/ll-backend/db/repositories"
	"github.com/fusion44/ll-backend/graph/model"
)

// App errors
var (
	ErrBadCredentials           = errors.New("Login credentials not valid")
	ErrUnauthenticated          = errors.New("Unauthenticated")
	ErrUnauthorized             = errors.New("Unauthorized")
	ErrInternalServer           = errors.New("Internal server error")
	ErrInvalidInput             = errors.New("Input not valid")
	ErrUnableToProcess          = errors.New("Unable to process FIT file")
	ErrDuplicateActivityForFile = errors.New("Duplicate activity for FIT file")
)

// Domain contains all business logic
type Domain struct {
	UsersRepo      repositories.UsersRepository
	ActivityRepo   repositories.ActivitiesRepository
	FileRepository repositories.FileDescRepository
	StatsRepo      repositories.StatsRepository
}

// NewDomain creates a new Domain instance
func NewDomain(
	usersRepo repositories.UsersRepository,
	activityRepo repositories.ActivitiesRepository,
	fileRepo repositories.FileDescRepository,
	statsRepo repositories.StatsRepository) *Domain {
	return &Domain{
		UsersRepo:      usersRepo,
		ActivityRepo:   activityRepo,
		FileRepository: fileRepo,
		StatsRepo:      statsRepo}
}

// Ownable makes an object ownable by an user
type Ownable interface {
	IsOwner(user *model.User) bool
}
