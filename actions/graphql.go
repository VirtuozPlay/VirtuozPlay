package actions

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gobuffalo/buffalo"
	"virtuozplay/graph"
	"virtuozplay/models"
	"virtuozplay/models/repository"
)

func init() {
	// Initialize the GraphQL server
	srv := handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{
		Performances: repository.NewPerformancesRepository(models.DB),
	}}))

	// Supported ways to submit queries
	srv.AddTransport(transport.GRAPHQL{})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})
	srv.AddTransport(transport.Websocket{})

	srv.SetQueryCache(lru.New(1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})

	GraphQLHandler = buffalo.WrapHandler(srv)
	GraphQLPlaygroundHandler = buffalo.WrapHandler(playground.Handler("GraphQL playground", "/graphql"))
}

// GraphQLHandler Manages the GraphQL endpoint of VirtuozPlay
var GraphQLHandler buffalo.Handler

// GraphQLPlaygroundHandler gives access to an interactive GraphQL playground in the browser
var GraphQLPlaygroundHandler buffalo.Handler
