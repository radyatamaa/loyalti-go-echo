package graphQL

import (
	"github.com/graphql-go/graphql"
)

var songType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Song",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.String,
		},
		"album": &graphql.Field{
			Type: graphql.String,
		},
		"title": &graphql.Field{
			Type: graphql.String,
		},
		"duration": &graphql.Field{
			Type: graphql.String,
		},
	},
})

var merchantType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Merchant",
	Fields: graphql.Fields{
		"id": &graphql.Field{
			Type: graphql.Int,
		},
		"created_at": &graphql.Field{
			Type: graphql.String,
		},
		"update_at": &graphql.Field{
			Type: graphql.String,
		},
		"delete_at": &graphql.Field{
			Type: graphql.DateTime,
		},
		"merchant_name": &graphql.Field{
			Type: graphql.String,
		},
		"phone_number": &graphql.Field{
			Type: graphql.String,
		},
		"email": &graphql.Field{
			Type: graphql.String,
		},
	},
})