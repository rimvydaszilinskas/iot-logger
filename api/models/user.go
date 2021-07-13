package models

import (
	"errors"
	"fmt"
	"net/mail"

	dbModels "github.com/rimvydaszilinskas/iot-logger/models"
	"gorm.io/gorm"
)

type AuthenticationUser struct {
	Email    string `json:"email" form:"email"`
	Password string `json:"password" form:"password"`
}

func (u *AuthenticationUser) ToUserModel() *dbModels.User {
	user := &dbModels.User{
		Email: u.Email,
	}

	user.SetPassword(u.Password)

	return user
}

// Validate checks user validity for creating a new user
func (u *AuthenticationUser) Validate(db *gorm.DB) map[string]string {
	errorsMap := map[string]string{}

	if len(u.Email) == 0 {
		errorsMap["email"] = "email cannot be empty"
	}
	if len(u.Password) < 6 || len(u.Password) > 50 {
		errorsMap["password"] = "password has to be between 6 and 50 characters long"
	}
	if _, err := mail.ParseAddress(u.Email); err != nil {
		errorsMap["email"] = fmt.Sprintf("%s is not a valid email", u.Email)
	}

	var existingUser dbModels.User

	if err := db.First(&existingUser, "email = ?", u.Email).Error; err != nil {

		if !errors.Is(err, gorm.ErrRecordNotFound) {
			return map[string]string{
				"error": "internal server error",
			}
		}

	}

	if existingUser.ID != 0 {
		errorsMap["email"] = "email already exists"
	}

	return errorsMap
}
