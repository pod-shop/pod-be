package user

import (
	"errors"
	"pod-be/cmd/db"
	"pod-be/pkg/modules/user/dto"

	"gorm.io/gorm"
)

func GetUserByID(id string) (*dto.UserDTO, error) {
	var user db.User // Adjust according to your actual import path and structure

	result := db.DB.First(&user, "id = ?", id)
	if result.Error != nil {
		if errors.Is(result.Error, gorm.ErrRecordNotFound) {
			return nil, errors.New("user not found")
		}
		return nil, result.Error
	}

	return &dto.UserDTO{
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone, // Keep the nullability
	}, nil
}

func CreateUser(name, email *string, phone string) (*dto.UserDTO, error) {
	user := db.User{
		Name:  name,
		Email: email,
		Phone: phone, // Keep the nullability
	}

	if err := db.DB.Create(&user).Error; err != nil {
		return nil, err
	}

	return &dto.UserDTO{
		Name:  user.Name,
		Email: user.Email,
		Phone: user.Phone, // Keep the nullability
	}, nil
}
