//go:generate go run github.com/99designs/gqlgen --verbose

package resolver

import (
	"github.com/fusion44/couchless-backend/domain"
)

// This file will not be regenerated automatically.
//
// It serves as dependency injection for your app, add any dependencies you require here.

type Resolver struct {
	Domain *domain.Domain
}
