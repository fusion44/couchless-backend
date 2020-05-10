package db

import (
	"context"
	"fmt"

	"github.com/go-pg/pg/v9"

	gcontext "github.com/fusion44/ll-backend/context"
)

// Logger ...
type Logger struct{}

// BeforeQuery is called before a query is executed
func (d Logger) BeforeQuery(ctx context.Context, q *pg.QueryEvent) (context.Context, error) {
	return ctx, nil
}

// AfterQuery logs the formatted query after a query was executed
func (d Logger) AfterQuery(ctx context.Context, q *pg.QueryEvent) error {
	fmt.Println(q.FormattedQuery())
	return nil
}

// New creates a new database connection
func New(cfg *gcontext.Config) *pg.DB {
	return pg.Connect(&pg.Options{
		User:     cfg.DBUser,
		Password: cfg.DBPassword,
		Database: cfg.DBName,
		Addr:     fmt.Sprintf("%s:%s", cfg.DBHost, cfg.DBPort),
	})
}
