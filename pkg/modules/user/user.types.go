package user

import "github.com/graphql-go/graphql"

// Define GraphQL Type for User
var UserType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "User",
		Fields: graphql.Fields{
			"id":    &graphql.Field{Type: graphql.String},
			"email": &graphql.Field{Type: graphql.String},
			"name":  &graphql.Field{Type: graphql.String},
			"phone": &graphql.Field{Type: graphql.String},
		},
	},
)
