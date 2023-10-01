package errorhandling

import "errors"

// Predefined errors
var (
	ErrUserNotFound    = errors.New("user not found")
	ErrInvalidPassword = errors.New("invalid password")
	ErrInvalidEmail    = errors.New("invalid email")
)

// GraphQLError is our custom error type that will also serve as a GraphQL error.
type GraphQLError struct {
	Message string `json:"message"`
}

// NewGraphQLError creates a new GraphQLError instance.
func NewGraphQLError(message string) *GraphQLError {
	return &GraphQLError{Message: message}
}

// Error implements the error interface for GraphQLError.
func (e *GraphQLError) Error() string {
	return e.Message
}
