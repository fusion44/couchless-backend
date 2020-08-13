package middleware

import (
	"context"
	"net/http"
	"strings"

	"github.com/fusion44/couchless-backend/db/repositories"
	"github.com/fusion44/couchless-backend/graph/model"

	"github.com/dgrijalva/jwt-go/request"
	"github.com/pkg/errors"

	"github.com/dgrijalva/jwt-go"
	gcontext "github.com/fusion44/couchless-backend/context"
)

// AuthMiddleware is a basic middleware to check authentication
func AuthMiddleware(cfg *gcontext.Config, userRepo *repositories.UsersRepository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(
			func(w http.ResponseWriter, r *http.Request) {
				token, err := parseToken(cfg, r)
				if err != nil {
					next.ServeHTTP(w, r)
					return
				}

				claims, ok := token.Claims.(jwt.MapClaims)

				if !ok || !token.Valid {
					next.ServeHTTP(w, r)
					return
				}

				user, err := userRepo.GetUserByID(claims["jti"].(string))

				if err != nil {
					next.ServeHTTP(w, r)
					return
				}

				ctx := context.WithValue(r.Context(), gcontext.KeyCurrentUser, user)

				next.ServeHTTP(w, r.WithContext(ctx))
			})
	}
}

var authHeaderExtractor = &request.PostExtractionFilter{
	Extractor: request.HeaderExtractor{"Authorization"},
	Filter:    stripBearerPrefixFromToken,
}

func stripBearerPrefixFromToken(token string) (string, error) {
	bearer := "BEARER"
	if len(token) > len(bearer) && strings.ToUpper(token[0:len(bearer)]) == bearer {
		return token[len(bearer)+1:], nil
	}

	return token, nil
}

var tokenExtractor = &request.MultiExtractor{
	authHeaderExtractor,
	request.ArgumentExtractor{"access_token"},
}

func parseToken(cfg *gcontext.Config, r *http.Request) (*jwt.Token, error) {
	jwtToken, err := request.ParseFromRequest(r, tokenExtractor, func(token *jwt.Token) (interface{}, error) {
		t := []byte(cfg.JWTSecret)
		return t, nil
	})
	return jwtToken, errors.Wrap(err, "parseToken error: ")
}

// GetCurrentUserFromContext retrieves the current user from the context.
// Returns nil if there is none.
func GetCurrentUserFromContext(ctx context.Context) (*model.User, error) {
	errMsg := errors.New("No user in context")
	if ctx.Value(gcontext.KeyCurrentUser) == nil {
		return nil, errMsg
	}

	user, ok := ctx.Value(gcontext.KeyCurrentUser).(*model.User)

	if !ok || user.ID == "" {
		return nil, errMsg
	}

	return user, nil
}
