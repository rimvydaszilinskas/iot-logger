package models

type Device struct {
	VerboseName string `json:"verbose_name"`
}

func (d *Device) Validate() map[string]string {
	l := len(d.VerboseName)

	if l < 3 || l > 20 {
		return map[string]string{"verbose_name": "must be between 3 and 20 characters long"}
	}
	return nil
}

type DeviceEntry struct {
	LocalIP  string `json:"local_ip" binding:"required"`
	PublicIP string `json:"public_ip" binding:"required"`
}
