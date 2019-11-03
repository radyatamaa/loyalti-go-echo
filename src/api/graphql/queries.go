package graphQL

import (
	"github.com/graphql-go/graphql"
)

// Root holds a pointer to a graphql object
type Root struct {
	Query *graphql.Object
}

// NewRoot returns base query type. This is where we add all the base queries
func NewRoot() *Root {
	// Create a resolver holding our database. Resolver can be found in resolvers.go
	resolver := SongResolver

	// Create a new Root that describes our base query set up. In this
	// example we have a user query that takes one argument called name
	root := Root{
		Query: graphql.NewObject(
			graphql.ObjectConfig{
				Name: "Query",
				Fields: graphql.Fields{
					"songs": &graphql.Field{
						// Slice of User type which can be found in types.go
						Type: graphql.NewList(songType),
						Resolve: resolver,
					},
				},
			},
		),
	}
	return &root
}
