package user

import (
	"pod-be/pkg/errorhandling"

	"github.com/graphql-go/graphql"
)

var QueryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: UserTypeInput,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{Type: graphql.String},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					id, ok := p.Args["id"].(string)
					if ok {
						user, err := GetUserByID(id) // Assuming a GetUserByID function is defined
						if err != nil {
							// Return GraphQL error if user is not found
							return nil, errorhandling.NewGraphQLError(errorhandling.ErrUserNotFound.Error())
						}
						return user, nil
					}
					// Return error if ID is not provided or invalid
					return nil, errorhandling.NewGraphQLError("Invalid ID")
				},
			},
		},
	},
)
var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query: QueryType,
	},
)
