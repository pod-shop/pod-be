package user

// Define the Data Transfer Object for a user
type UserDTO struct {
	ID    string `json:"id"`
	Email string `json:"email"`
}
