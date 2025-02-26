package models

import (
	"errors"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Email    string `json:"email"`
	Password string `json:"password"`
	Name     string `json:"name"`
}

func GetUser(id uint) (*User, error) {
	var user User

	if err := Database.Where("id = ?", id).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (user *User) UpdateUser(id uint) (*User, error) {
	if user.Password != "" {
		if err := user.HashPassword(); err != nil {
			return nil, errors.New("failed to update user password")
		}
	}

	if err := Database.Model(&User{}).Where("id = ?", id).Updates(user).Error; err != nil {
		return nil, errors.New("failed to update user")
	}

	return user, nil
}
