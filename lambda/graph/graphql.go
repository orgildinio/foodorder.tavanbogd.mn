package graph

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/go-redis/redis"
	"github.com/gofiber/fiber/v2"
	"github.com/gorilla/websocket"
	"github.com/lambda-platform/lambda/config"
	lambdaPlayground "github.com/lambda-platform/lambda/graphql/playground"
	"github.com/valyala/fasthttp/fasthttpadaptor"
	"lambda/lambda/graph/generated"
	"net/http"
	"sync"
	"time"
)

type Cache struct {
	client redis.UniversalClient
	ttl    time.Duration
}

func Set(e *fiber.App) {

	graphqlConfig := generated.Config{Resolvers: &Resolver{

		mutex: sync.Mutex{},
	}}

	graphqlHandler := handler.New(generated.NewExecutableSchema(graphqlConfig))
	playgroundHandler := playground.Handler("GraphQL playground", "/query")
	lambdaPlaygroundHandler := lambdaPlayground.Handler("GraphQL playground", "/query")

	graphqlHandler.AddTransport(transport.Options{})
	graphqlHandler.AddTransport(transport.GET{})
	graphqlHandler.AddTransport(transport.POST{})
	graphqlHandler.AddTransport(transport.MultipartForm{})
	graphqlHandler.SetQueryCache(lru.New(1000))

	if config.Config.Graphql.Debug == "true" {
		graphqlHandler.Use(extension.Introspection{})
	}

	graphqlHandler.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})
	graphqlHandler.AddTransport(transport.Websocket{
		KeepAlivePingInterval: 10 * time.Second,
		Upgrader: websocket.Upgrader{
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	})

	e.Post("/query", func(c *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			graphqlHandler.ServeHTTP(writer, request)
		})(c.Context())
		return nil
	})

	e.Get("/play", func(c *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			playgroundHandler.ServeHTTP(writer, request)
		})(c.Context())
		return nil
	})

	e.Get("/play-full", func(c *fiber.Ctx) error {
		fasthttpadaptor.NewFastHTTPHandlerFunc(func(writer http.ResponseWriter, request *http.Request) {
			lambdaPlaygroundHandler.ServeHTTP(writer, request)
		})(c.Context())
		return nil
	})
}
