package main

import (
	"log"
	"net/http"
	"os"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/seanaye/geog483-final/server/graph"
	"github.com/seanaye/geog483-final/server/graph/generated"
)

const defaultPort = "8080"
const default_redis = "127.0.0.1:6379"

var ctx = context.Background()



func main() {
	port := os.Getenv("PORT")
	if port == "" {
		port = defaultPort
	}
	redis_addr := os.Getenv("REDIS_HOST")

	if redis_addr == "" {
		redis_addr = default_redis
	}

	userService := &redis.UserService{Addr: redis_addr}

	srv := handler.NewDefaultServer(generated.NewExecutableSchema(generated.Config{Resolvers: &graph.Resolver{User: userService}}))

	http.Handle("/", playground.Handler("GraphQL playground", "/query"))
	http.Handle("/query", srv)

	log.Printf("connect to http://localhost:%s/ for GraphQL playground", port)
	log.Fatal(http.ListenAndServe(":"+port, nil))
}
