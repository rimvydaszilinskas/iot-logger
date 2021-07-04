package models

import (
	"crypto/rand"

	"golang.org/x/crypto/bcrypt"
	"gorm.io/gorm"

	"github.com/rimvydaszilinskas/announcer-backend/utils"
)

type User struct {
	gorm.Model
	Email    string `json:"email" gorm:"type:varchar(60);unique"`
	Password string `json:"-"`
	Token    string `json:"-"`
}

func (u *User) SetPassword(password string) (*string, error) {
	hashedBytes, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)

	if err != nil {
		return nil, err
	}

	hashedPassword := string(hashedBytes)
	return &hashedPassword, err
}

func (u *User) ComparePasswordHash(password string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(u.Password), []byte(password))
	return err != nil
}

func (u *User) GenerateToken() {
	b := make([]byte, 40)
	if _, err := rand.Read(b); err != nil {
		return
	}
	u.Token = utils.GenerateToken(40)
}
