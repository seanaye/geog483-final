package main

import (
	"context"
	"net/http"
	"os"
	"time"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/gorilla/websocket"
	"github.com/seanaye/geog483-final/server/graph"
	"github.com/seanaye/geog483-final/server/graph/generated"
	"github.com/seanaye/geog483-final/server/pkg/directive"
	"github.com/seanaye/geog483-final/server/pkg/middleware"
	"github.com/seanaye/geog483-final/server/pkg/redis"

	"github.com/go-chi/chi"
	"github.com/go-chi/cors"
)

const defaultRedis = "redis:6379"

var ctx = context.Background()

func main() {

	redis_addr := os.Getenv("REDIS_HOST")
	if redis_addr == "" {
		redis_addr = defaultRedis
	}

	// create database service
	service := &redis.RedisService{Host: redis_addr}
	service.Clear()

	// create router and add middleware
	router := chi.NewRouter()

	router.Use(cors.Handler(cors.Options{
		AllowedOrigins: []string{"https://*", "http://*"},
		AllowedMethods: []string{"GET", "POST", "PUT", "DELETE", "OPTIONS"},
		AllowedHeaders: []string{"Accept", "Authorization", "Content-Type", "X-CSRF-Token"},
	}))
	router.Use(middleware.Auth(*service))
	//////

	conf := generated.Config{Resolvers: &graph.Resolver{Session: service, User: service, Message: service}}
	conf.Directives.Auth = directive.Auth

	srv := handler.New(generated.NewExecutableSchema(conf))

	// Configure transport settings
	srv.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
		InitFunc: middleware.WSInit(*service),
	})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.MultipartForm{})

	srv.SetQueryCache(lru.New(1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	srv.Use(extension.Introspection{})
	//////////

	// serve gqlgen app via chi router
	router.Handle("/query", srv)

	err := http.ListenAndServe("0.0.0.0:8080", router)
	if err != nil {
		panic(err)
	}
}
