package models

type NetworkInterface struct {
	Name         string `json:"name"`
	HardwareAddr string `json:"hardware_address"`
	LocalIp      string `json:"local_ip"`
	NetworkIp    string `json:"network_ip"`
	PublicIp     string `json:"public_ip"`
}
