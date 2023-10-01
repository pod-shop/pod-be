package dto

// Define the Data Transfer Object for a user
type UserDTO struct {
	ID    string `json:"id"`
	Name  string `json:"name"`
	Email string `json:"email"`
	Phone string `json:"phone"`
}
type UserInputDTO struct {
	ID string `json:"id"`
}
