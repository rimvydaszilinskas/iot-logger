package db

import (
	"errors"

	"github.com/rimvydaszilinskas/announcer-backend/models"
	"gorm.io/gorm"
)

func (db *Connection) GetUserByEmail(email string) (*models.User, bool, error) {
	var user models.User

	if err := db.DB.First(&user, "email = ?", email).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, false, nil
		}
		return nil, false, err
	}

	return &user, true, nil
}

func (db *Connection) GetUserByToken(token string) (*models.User, bool, error) {
	var user models.User

	if err := db.DB.First(&user, "token = ?", token).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, false, nil
		}
		return nil, false, err
	}

	return &user, true, nil
}
