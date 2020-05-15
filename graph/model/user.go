package model

import (
	"time"

	"github.com/dgrijalva/jwt-go"

	"golang.org/x/crypto/bcrypt"

	gcontext "github.com/fusion44/ll-backend/context"
)

// User represents all user data
type User struct {
	ID        string     `json:"id"`
	Username  string     `json:"username"`
	Email     string     `json:"email"`
	Password  string     `json:"-"`
	IPAddress string     `json:"ipAddress"`
	CreatedAt time.Time  `json:"createdAt"`
	UpdatedAt time.Time  `json:"updatedAt"`
	DeletedAt *time.Time `json:"-" pg:",soft_delete"`
}

// HashPassword one-way hashes the password of the user
func (u *User) HashPassword(password string) error {
	bytePW := []byte(password)
	passwordHash, err := bcrypt.GenerateFromPassword(bytePW, bcrypt.DefaultCost)
	if err != nil {
		return err
	}

	u.Password = string(passwordHash)

	return nil
}

// ComparePassword checks the given password for validity
func (u *User) ComparePassword(pw string) error {
	bytePW := []byte(pw)
	byteHashedPW := []byte(u.Password)
	return bcrypt.CompareHashAndPassword(byteHashedPW, bytePW)
}

// GenToken generates an auth token for the user
func (u *User) GenToken(cfg *gcontext.Config) (*AuthToken, error) {
	expiresAt := time.Now().Add(time.Hour * 24 * 7)
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.StandardClaims{
		ExpiresAt: expiresAt.Unix(),
		Id:        u.ID,
		IssuedAt:  time.Now().Unix(),
		Issuer:    cfg.AppName,
	})

	accessToken, err := token.SignedString([]byte(cfg.JWTSecret))
	if err != nil {
		return nil, err
	}

	return &AuthToken{
		AccessToken: accessToken,
		ExpiredAt:   expiresAt,
	}, nil
}
