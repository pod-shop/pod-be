package dto

// Define the Data Transfer Object for a user
type UserDTO struct {
	ID    string  `json:"id"`
	Name  *string `json:"name,omitempty"`
	Email *string `json:"email,omitempty"`
	Phone string  `json:"phone"`
}
type UserInputDTO struct {
	ID string `json:"id"`
}
