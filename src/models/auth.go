package models

import (
	"encoding/base64"
	"errors"
	"github.com/Arenelin/List-of-current-affairs/src/utils"
	"golang.org/x/crypto/argon2"
)

func (user *User) HashPassword() error {
	time := uint32(1)
	memory := uint32(64 * 1024)
	threads := uint8(4)
	keyLen := uint32(32)
	salt, err := utils.GenerateSalt(16)

	if err != nil {
		return err
	}

	hashed := argon2.IDKey([]byte(user.Password), salt, time, memory, threads, keyLen)

	user.Password = base64.StdEncoding.EncodeToString(salt) + "$" + base64.StdEncoding.EncodeToString(hashed)
	return nil
}

func GetUserByEmail(email string) (*User, error) {
	var user User

	if err := Database.Where("email = ?", email).First(&user).Error; err != nil {
		return nil, errors.New("user not found")
	}

	return &user, nil
}

func (user *User) Register() (*User, error) {

	if user, err := GetUserByEmail(user.Email); user != nil && err == nil {
		return nil, errors.New("user already exists")
	}

	if err := user.HashPassword(); err != nil {
		return nil, errors.New("failed to create user password")
	}

	if err := Database.Model(&user).Create(&user).Error; err != nil {
		return nil, errors.New("failed to create user")
	}

	return user, nil
}

func (user *User) Login() (*User, error) {

	existedUser, err := GetUserByEmail(user.Email)
	if err != nil {
		return nil, errors.New("email or password incorrect")
	}

	if isCorrectPassword := utils.ComparePassword(user.Password, existedUser.Password); !isCorrectPassword {
		return nil, errors.New("email or password incorrect")
	}

	return existedUser, nil
}
