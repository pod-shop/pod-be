package user

import (
	"pod-be/pkg/errorhandling"
	"pod-be/pkg/modules/user/dto"

	"github.com/graphql-go/graphql"
)

var QueryType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "Query",
		Fields: graphql.Fields{
			"user": &graphql.Field{
				Type: UserType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String), // Ensure that id is provided
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					input := dto.UserInputDTO{
						ID: p.Args["id"].(string),
					}

					user, err := GetUserByID(input.ID)
					if err != nil {
						return nil, errorhandling.NewGraphQLError(errorhandling.ErrUserNotFound.Error())
					}

					return dto.UserDTO{
						ID:    user.ID,
						Name:  user.Name,
						Email: user.Email,
						Phone: user.Phone,
					}, nil
				},
			},
		},
	},
)
var MutationType = graphql.NewObject(graphql.ObjectConfig{
	Name: "Mutation",
	Fields: graphql.Fields{
		"createUser": &graphql.Field{
			Type:        UserType, // This should be your UserType defined in the dto package
			Description: "Create a new user",
			Args: graphql.FieldConfigArgument{
				"name":  &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"email": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
				"phone": &graphql.ArgumentConfig{Type: graphql.String},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				name, _ := params.Args["name"].(string)
				email, _ := params.Args["email"].(string)
				phone, _ := params.Args["phone"].(string)

				user, err := CreateUser(name, email, phone)
				if err != nil {
					return nil, err
				}
				return user, nil
			},
		},
	},
})
var Schema, _ = graphql.NewSchema(
	graphql.SchemaConfig{
		Query:    QueryType,
		Mutation: MutationType,
	},
)
