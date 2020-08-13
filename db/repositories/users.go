package repositories

import (
	"fmt"

	"github.com/fusion44/couchless-backend/graph/model"
	"github.com/go-pg/pg/v9"
)

// UsersRepository contains all functions regarding users
type UsersRepository struct {
	DB *pg.DB
}

// CreateUser inserts a user into the database
func (r *UsersRepository) CreateUser(tx *pg.Tx, user *model.User) (*model.User, error) {
	_, err := tx.Model(user).Returning("*").Insert()
	return user, err
}

// GetUserByField returns the user for the given field
func (r *UsersRepository) GetUserByField(field, value string) (*model.User, error) {
	var user model.User
	err := r.DB.Model(&user).Where(fmt.Sprintf("%v = ?", field), value).First()
	return &user, err
}

// GetUserByID returns the user for the given ID
func (r *UsersRepository) GetUserByID(id string) (*model.User, error) {
	return r.GetUserByField("id", id)
}

// GetUserByEmail returns the user for the given email
func (r *UsersRepository) GetUserByEmail(email string) (*model.User, error) {
	return r.GetUserByField("email", email)
}

// GetUserByUsername returns the user for the given email
func (r *UsersRepository) GetUserByUsername(username string) (*model.User, error) {
	return r.GetUserByField("username", username)
}

// GetUsers returns all users in the database
func (r *UsersRepository) GetUsers() ([]*model.User, error) {
	var users []*model.User
	err := r.DB.Model(&users).Order("id").Select()
	return users, err
}
