package actions

import (
	"net/http"
	"virtuozplay/graph"
	"virtuozplay/models/repository"

	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gobuffalo/buffalo"
	"github.com/gorilla/websocket"
)

func lazyInit(c buffalo.Context) {
	performances := c.Value("performances").(*repository.Performances)
	songs := c.Value("songs").(*repository.Songs)
	users := c.Value("users").(*repository.Users)

	// Initialize the GraphQL server
	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		Performances: performances,
		Songs:        songs,
		Users:        users,
	}}))

	// Supported ways to submit queries
	srv.AddTransport(transport.GRAPHQL{})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})
	srv.AddTransport(transport.Websocket{Upgrader: websocket.Upgrader{CheckOrigin: func(r *http.Request) bool { return true }}})

	srv.SetQueryCache(lru.New(1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	graphQLHandler = buffalo.WrapHandler(srv)
	graphQLPlaygroundHandler = buffalo.WrapHandler(playground.Handler("GraphQL playground", "/graphql"))
}

// GraphQLHandler Manages the GraphQL endpoint of VirtuozPlay
func GraphQLHandler(c buffalo.Context) error {
	if graphQLHandler == nil {
		lazyInit(c)
	}
	return graphQLHandler(c)
}

// GraphQLPlaygroundHandler gives access to an interactive GraphQL playground in the browser
func GraphQLPlaygroundHandler(c buffalo.Context) error {
	if graphQLPlaygroundHandler == nil {
		lazyInit(c)
	}
	return graphQLPlaygroundHandler(c)
}

var graphQLHandler buffalo.Handler
var graphQLPlaygroundHandler buffalo.Handler
