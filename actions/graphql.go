package actions

import (
	"github.com/gobuffalo/buffalo"
	"github.com/graphql-go/graphql"
	"github.com/graphql-go/handler"
	"log"
)

var schema = makeSchema()

// GraphQLHandler Manages the GraphQL endpoint of VirtuozPlay
func GraphQLHandler() buffalo.Handler {
	devMode := ENV != "production"

	h := handler.New(&handler.Config{
		Schema:     &schema,
		Pretty:     devMode,
		GraphiQL:   devMode,
		Playground: devMode,
	})

	return buffalo.WrapHandler(h)
}

// makeSchema Creates the GraphQL schema for VirtuozPlay
// TODO: replace by gqlgen, and actual data
func makeSchema() graphql.Schema {
	virtuozPlayObject := graphql.NewObject(graphql.ObjectConfig{
		Name: "VirtuozPlay",
		Fields: graphql.Fields{
			"version": &graphql.Field{
				Type: graphql.String,
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					return "0.1.0", nil
				},
			},
		},
	})
	rootQuery := graphql.ObjectConfig{Name: "RootQuery", Fields: graphql.Fields{
		"virtuozPlay": &graphql.Field{
			Type: virtuozPlayObject,
			Resolve: func(p graphql.ResolveParams) (interface{}, error) {
				return p.Source, nil
			},
		},
	}}

	schema, err := graphql.NewSchema(graphql.SchemaConfig{Query: graphql.NewObject(rootQuery)})
	if err != nil {
		log.Fatalf("failed to create new schema, error: %v", err)
	}
	return schema
}
