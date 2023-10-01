package user

import "github.com/graphql-go/graphql"

// Define GraphQL Type for User
var UserTypeInput = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id":    &graphql.Field{Type: graphql.String},
			"email": &graphql.Field{Type: graphql.String},
		},
	},
)
