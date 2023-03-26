package models

import (
	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"
)

type User struct {
	gorm.Model
	Name     string `json:"name"`
	Username string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     uint   `json:"role"`
	Profile  Profile
}

type TokenRequest struct {
	Email    string `json:"email"`
	Password string `json:"password"`
}

//hashpassword
func (user *User) Hashpassword(password string) error {
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)

	if err != nil {
		return err
	}

	user.Password = string(bytes)

	return nil
}

//checkpassword

func (user *User) CheckPassword(password string) error {
	err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password))

	if err != nil {
		return err
	}

	return nil
}
