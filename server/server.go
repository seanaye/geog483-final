package main

import (
	"context"
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/seanaye/geog483-final/server/graph"
	"github.com/seanaye/geog483-final/server/graph/generated"
	"github.com/seanaye/geog483-final/server/pkg/directive"
	"github.com/seanaye/geog483-final/server/pkg/redis"
	"github.com/seanaye/geog483-final/server/pkg/middleware"

	"github.com/go-chi/chi"
)

const defaultPort = "8080"
const defaultRedis = "127.0.0.1:6379"

var ctx = context.Background()

func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}

	redis_addr := os.Getenv("REDIS_HOST")
	if redis_addr == "" {
		redis_addr = defaultRedis
	}

	service := &redis.RedisService{Host: redis_addr}
	service.Clear()

	router := chi.NewRouter()
	router.Use(middleware.Auth(*service))

	conf := generated.Config{Resolvers: &graph.Resolver{Session: service, User: service}}
	conf.Directives.Auth = directive.Auth

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(conf))

	router.Handle("/", playground.Handler("GraphQL playground", "/query"))
	router.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	err := http.ListenAndServe(":8080", router)
	if err != nil {
		panic(err)
	}
}
