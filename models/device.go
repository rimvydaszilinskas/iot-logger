package models

import (
	"github.com/rimvydaszilinskas/announcer-backend/utils"
	"gorm.io/gorm"
)

type Device struct {
	gorm.Model
	VerboseName string

	OwnerID int
	Owner   User `gorm:"constraint:OnDelete:CASCADE"`

	Token string

	Entries []DeviceEntry `gorm:"many2many:entries;"`
}

type DeviceEntry struct {
	gorm.Model

	PublicIP string `json:"public_ip"`
	LocalIP  string `json:"local_ip"`
}

func (d *Device) GenerateToken() {
	d.Token = utils.GenerateToken(40)
}
