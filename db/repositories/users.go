package repositories

import (
	"github.com/fusion44/ll-backend/graph/model"
	"github.com/go-pg/pg/v9"
)

// UsersRepository contains all functions regarding users
type UsersRepository struct {
	DB *pg.DB
}

// GetUserByID returns the user for the given ID
func (r *UsersRepository) GetUserByID(id string) (*model.User, error) {
	var user model.User
	err := r.DB.Model(&user).Where("id = ?", id).First()

	if err != nil {
		return nil, err
	}

	return &user, err
}

// GetUsers returns all users in the database
func (r *UsersRepository) GetUsers() ([]*model.User, error) {
	var users []*model.User
	err := r.DB.Model(&users).Order("id").Select()

	if err != nil {
		return nil, err
	}

	return users, nil
}
