package loader

import (
	"context"
	"net/http"
	"time"

	gcontext "github.com/fusion44/ll-backend/context"
	"github.com/fusion44/ll-backend/graph/model"
	"github.com/go-pg/pg/v9"
)

// UserLoaderMiddleware adds User-DataLoaders to the handler
func UserLoaderMiddleware(db *pg.DB, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		userloader := UserLoader{
			maxBatch: 100,
			wait:     1 * time.Millisecond,
			fetch: func(ids []string) ([]*model.User, []error) {
				var users []*model.User

				err := db.Model(&users).Where("id in (?)", pg.In(ids)).Select()

				if err != nil {
					return nil, []error{err}
				}

				// The users slice is not guaranteed to be in the right order
				// so we have to sort the according to the input ids
				u := make(map[string]*model.User, len(users))
				for _, user := range users {
					u[user.ID] = user
				}

				result := make([]*model.User, len(ids))

				for i, id := range ids {
					result[i] = u[id]
				}

				return result, []error{err}
			},
		}

		ctx := context.WithValue(r.Context(), gcontext.KeyUserloaderMiddleware, &userloader)

		next.ServeHTTP(w, r.WithContext(ctx))
	})
}

// GetUserLoader returns the UserLoader from the given context
func GetUserLoader(ctx context.Context) *UserLoader {
	return ctx.Value(gcontext.KeyUserloaderMiddleware).(*UserLoader)
}
