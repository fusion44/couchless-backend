package main

import (
	"log"
	"net/http"

	loader "github.com/fusion44/ll-backend/db/loaders"
	"github.com/fusion44/ll-backend/db/repositories"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/fusion44/ll-backend/db"
	"github.com/fusion44/ll-backend/graph/generated"
	"github.com/fusion44/ll-backend/graph/resolver"

	gcontext "github.com/fusion44/ll-backend/context"
)

func main() {
	cfg := gcontext.LoadConfig(".")
	DB := db.New(cfg)
	DB.AddQueryHook(db.Logger{})
	defer DB.Close()

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &resolver.Resolver{
		UsersRepo:    repositories.UsersRepository{DB: DB},
		ActivityRepo: repositories.ActivitiesRepository{DB: DB},
	}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", loader.UserLoaderMiddleware(DB, srv))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", cfg.ServerPort)
	log.Fatal(http.ListenAndServe(":"+cfg.ServerPort, nil))
}
