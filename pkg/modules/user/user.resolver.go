package user

import (
	"fmt"
	"pod-be/pkg/errorhandling"
	"pod-be/pkg/modules/elasticsearch"
	"pod-be/pkg/modules/user/dto"

	"github.com/graphql-go/graphql"
)

var SchoolType = graphql.NewObject(
	graphql.ObjectConfig{
		Name: "School",
		Fields: graphql.Fields{
			"name": &graphql.Field{
				Type: graphql.String,
			},
			"description": &graphql.Field{
				Type: graphql.String,
			},
			"street": &graphql.Field{
				Type: graphql.String,
			},
			"city": &graphql.Field{
				Type: graphql.String,
			},
			"state": &graphql.Field{
				Type: graphql.String,
			},
			"zip": &graphql.Field{
				Type: graphql.String,
			},
			"location": &graphql.Field{
				Type: graphql.NewList(graphql.Float),
			},
			"fees": &graphql.Field{
				Type: graphql.Float,
			},
			"tags": &graphql.Field{
				Type: graphql.NewList(graphql.String),
			},
			"rating": &graphql.Field{
				Type: graphql.Float,
			},
		},
	},
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
			"school": &graphql.Field{
				Type: SchoolType,
				Args: graphql.FieldConfigArgument{
					"id": &graphql.ArgumentConfig{
						Type: graphql.NewNonNull(graphql.String),
					},
				},
				Resolve: func(p graphql.ResolveParams) (interface{}, error) {
					// Create a new Elasticsearch client
					es, err := elasticsearch.NewElasticsearchClient()
					if err != nil {
						return nil, fmt.Errorf("error creating the Elasticsearch client: %s", err)
					}

					// Create a new Elasticsearch service
					esService := NewElasticsearchService(es)

					// Call the GetSchoolByID method
					schoolDoc, err := esService.GetDocumentByID("school", p.Args["id"].(string))
					if err != nil {
						return nil, fmt.Errorf("error retrieving school: %s", err)
					}

					// Map Elasticsearch document fields to GraphQL school fields
					return schoolDoc["_source"], nil
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
				"name":  &graphql.ArgumentConfig{Type: graphql.String},
				"email": &graphql.ArgumentConfig{Type: graphql.String},
				"phone": &graphql.ArgumentConfig{Type: graphql.NewNonNull(graphql.String)},
			},
			Resolve: func(params graphql.ResolveParams) (interface{}, error) {
				nameArg, nameProvided := params.Args["name"].(string)
				emailArg, emailProvided := params.Args["email"].(string)
				phone, _ := params.Args["phone"].(string)

				var name *string
				var email *string
				if nameProvided {
					name = &nameArg
				}
				if emailProvided {
					email = &emailArg
				}
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
