package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/go-chi/chi/middleware"

	loader "github.com/fusion44/ll-backend/db/loaders"
	"github.com/fusion44/ll-backend/db/repositories"
	"github.com/go-chi/chi"
	"github.com/rs/cors"

	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/99designs/gqlgen/handler"
	"github.com/fusion44/ll-backend/db"
	"github.com/fusion44/ll-backend/graph/generated"
	"github.com/fusion44/ll-backend/graph/resolver"
	projMiddleware "github.com/fusion44/ll-backend/middleware"

	gcontext "github.com/fusion44/ll-backend/context"
)

// AppConfig holds the global configuration
var AppConfig *gcontext.Config

func main() {
	AppConfig := gcontext.LoadConfig(".")
	DB := db.New(AppConfig)
	DB.AddQueryHook(db.Logger{})
	defer DB.Close()

	userRepo := repositories.UsersRepository{DB: DB}
	activityRepo := repositories.ActivitiesRepository{DB: DB}

	router := chi.NewRouter()
	router.Use(cors.New(cors.Options{
		AllowedOrigins:   []string{fmt.Sprintf("http://localhost:%s", AppConfig.ServerPort)},
		AllowCredentials: true,
		Debug:            true,
	}).Handler)
	router.Use(middleware.RequestID)
	router.Use(middleware.Logger)
	router.Use(projMiddleware.AuthMiddleware(AppConfig, &userRepo))
	router.Use(gcontext.ConfigMiddleware(AppConfig))

	c := generated.Config{Resolvers: &resolver.Resolver{
		ActivityRepo: activityRepo,
		UsersRepo:    userRepo,
	}}

	queryHander := handler.GraphQL(generated.NewExecutableSchema(c))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", loader.UserLoaderMiddleware(DB, queryHander))

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", AppConfig.ServerPort)
	log.Fatal(http.ListenAndServe(":"+AppConfig.ServerPort, router))
}
