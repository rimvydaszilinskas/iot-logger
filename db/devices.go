package db

import (
	"errors"

	"github.com/rimvydaszilinskas/announcer-backend/models"
	"gorm.io/gorm"
)

func (db *Connection) GetUserDevices(user *models.User) ([]models.Device, bool, error) {
	var devices []models.Device

	err := db.DB.Find(&devices, "owner_id = ?", user.ID).Error

	if err != nil {
		return nil, false, err
	}

	return devices, true, nil
}

func (db *Connection) GetDeviceByToken(token string) (*models.Device, bool, error) {
	var device models.Device

	if err := db.DB.Preload("Entries").First(&device, "token = ?", token).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, false, nil
		}
		return nil, false, err
	}

	return &device, true, nil
}

func (db *Connection) GetUserDevice(user *models.User, id string) (*models.Device, bool, error) {
	var device models.Device

	if err := db.DB.Preload("Entries").First(&device, "id = ? AND owner_id = ?", id, user.ID).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, false, nil
		}
		return nil, false, err
	}
	return &device, true, nil
}
