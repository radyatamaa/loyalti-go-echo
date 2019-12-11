package graphQL

import (
	"github.com/graphql-go/graphql"
)

// Root holds a pointer to a graphql object
type Root struct {
	Query *graphql.Object
}

type Page struct{
	page int `json:"null"`
}

// NewRoot returns base query type. This is where we add all the base queries
func NewRoot() *Root {
	// Create a resolver holding our database. Resolver can be found in resolvers.go

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
						Resolve: SongResolver,
					},
					"merchant" : &graphql.Field{
						Type:graphql.NewList(merchantType),
						Args: graphql.FieldConfigArgument{
							"page": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"size": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"sort": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"email" : &graphql.ArgumentConfig{
								Type: graphql.String,
							},
							"id" : &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
						},
						Resolve: MerchantResolver,
					},
					"category" : &graphql.Field{
						Type:graphql.NewList(categoryType),
						Args: graphql.FieldConfigArgument{
							"page": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"size": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"sort": &graphql.ArgumentConfig{
								Type:graphql.Int,
							},
						},
						Resolve: MerchantCategoryResolver,
					},
					"card" : &graphql.Field{
						Type:graphql.NewList(cardType),
						Args: graphql.FieldConfigArgument{
							"page": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"size": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"sort": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
						},
						Resolve: MerchantCardResolver,
					},
					"socialmedia" : &graphql.Field{
						Type:graphql.NewList(socialmediaType),
						Args:graphql.FieldConfigArgument{
							"page": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"size": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
							"sort": &graphql.ArgumentConfig{
								Type: graphql.Int,
							},
						},
						Resolve: SocialMediaResolver,
					},
					"program" : &graphql.Field{
						Type:graphql.NewList(programType),
						Args:graphql.FieldConfigArgument{
							"page": &graphql.ArgumentConfig{
								Type:graphql.Int,
							},
							"size": &graphql.ArgumentConfig{
								Type:graphql.Int,
							},
							"sort": &graphql.ArgumentConfig{
								Type:graphql.Int,
							},
							"category": &graphql.ArgumentConfig{
								Type:graphql.Int,
							},
							"id": &graphql.ArgumentConfig{
								Type:graphql.Int,
							},
						},
						Resolve: ProgramResolver,
					},
					"special" : &graphql.Field{
						Type:graphql.NewList(specialprogramType),
						Args:graphql.FieldConfigArgument{
							"page": &graphql.ArgumentConfig{
								Type:graphql.Int,
							},
							"size": &graphql.ArgumentConfig{
								Type:graphql.Int,
							},
							"sort": &graphql.ArgumentConfig{
								Type:graphql.Int,
							},
							"category": &graphql.ArgumentConfig{
								Type:graphql.Int,
							},
							"id": &graphql.ArgumentConfig{
								Type:graphql.Int,
							},
						},
						Resolve: SpecialProgramResolver,
					},
					"outlet" : &graphql.Field{
						Type:graphql.NewList(outletType),
						Args:graphql.FieldConfigArgument{
							"page": &graphql.ArgumentConfig{
								Type:graphql.Int,
							},
							"size": &graphql.ArgumentConfig{
								Type:graphql.Int,
							},
							"id": &graphql.ArgumentConfig{
								Type:graphql.Int,
							},
						},
						Resolve: OutletResolver,
					},
					"employee" : &graphql.Field{
						Type:graphql.NewList(employeeType),
						Args:graphql.FieldConfigArgument{
							"page": &graphql.ArgumentConfig{
								Type:graphql.Int,
							},
							"size": &graphql.ArgumentConfig{
								Type:graphql.Int,
							},
							"sort": &graphql.ArgumentConfig{
								Type:graphql.Int,
							},
						},
						Resolve: EmployeeResolver,
					},
				},
			},
		),
	}
	return &root
}
