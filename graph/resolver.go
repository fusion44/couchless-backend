//go:generate go run github.com/99designs/gqlgen --verbose

package graph

import "github.com/fusion44/ll-backend/db/repositories"

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	UsersRepo    repositories.UsersRepository
	ActivityRepo repositories.ActivitiesRepository
}
