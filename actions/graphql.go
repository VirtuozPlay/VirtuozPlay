package actions

import (
	"github.com/99designs/gqlgen/graphql/handler"
	"github.com/99designs/gqlgen/graphql/handler/extension"
	"github.com/99designs/gqlgen/graphql/handler/lru"
	"github.com/99designs/gqlgen/graphql/handler/transport"
	"github.com/99designs/gqlgen/graphql/playground"
	"github.com/gobuffalo/buffalo"
	"virtuozplay/graph"
)

var srv *handler.Server

func init() {
	// Initialize the GraphQL server
	srv = handler.New(graph.NewExecutableSchema(graph.Config{Resolvers: &graph.Resolver{}}))

	// Supported ways to submit queries
	srv.AddTransport(transport.GRAPHQL{})
	srv.AddTransport(transport.Options{})
	srv.AddTransport(transport.GET{})
	srv.AddTransport(transport.POST{})
	srv.AddTransport(transport.MultipartForm{})

	srv.SetQueryCache(lru.New(1000))

	srv.Use(extension.Introspection{})
	srv.Use(extension.AutomaticPersistedQuery{
		Cache: lru.New(100),
	})
}

// GraphQLHandler Manages the GraphQL endpoint of VirtuozPlay
func GraphQLHandler() buffalo.Handler {
	return buffalo.WrapHandler(srv)
}

// GraphQLPlaygroundHandler gives access to an interactive GraphQL playground in the browser
func GraphQLPlaygroundHandler() buffalo.Handler {
	return buffalo.WrapHandler(playground.Handler("GraphQL playground", "/graphql"))
}
