package models

import (
	"fmt"

	"github.com/rimvydaszilinskas/iot-logger/utils"
	"gorm.io/gorm"
)

type Device struct {
	gorm.Model
	VerboseName string

	OwnerID int
	Owner   User `gorm:"constraint:OnDelete:CASCADE"`

	Token string
}

func (d *Device) GenerateToken() {
	d.Token = utils.GenerateToken(40)
}

func (d *Device) GetRedisKey() string {
	return fmt.Sprintf("device-entry-%d-%d", d.OwnerID, d.ID)
}
