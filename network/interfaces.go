package network

import (
	"fmt"
	"net"
	"os/exec"
	"strings"

	"github.com/rimvydaszilinskas/announcer-backend/models"
)

func getActiveNetworkInterfaceNameAndIp() (*string, *string, *string, error) {
	cmd := exec.Command("ip", "route", "get", "1.1.1.1")
	output, err := cmd.Output()

	if err != nil {
		return nil, nil, nil, err
	}

	elements := strings.Split(string(output), " ")

	localIp, deviceName, subnet := elements[2], elements[4], elements[6]

	return &localIp, &deviceName, &subnet, nil
}

func GetActiveNetworkInterface() (*models.NetworkInterface, error) {
	var iFace net.Interface
	localIp, deviceName, subnet, err := getActiveNetworkInterfaceNameAndIp()

	if err != nil {
		return nil, fmt.Errorf("error retrieving active network interface details - %w", err)
	}

	ifaces, err := net.Interfaces()

	if err != nil {
		return nil, fmt.Errorf("error retrieving network interfaces - %w", err)
	}

	for _, iface := range ifaces {
		if iface.Name == *deviceName {
			iFace = iface
		}
	}

	publicIp, err := getPublicIp()

	if err != nil {
		return nil, fmt.Errorf("error retrieving public IP - %w", err)
	}

	return &models.NetworkInterface{
		Name:         *deviceName,
		HardwareAddr: iFace.HardwareAddr.String(),
		LocalIp:      *localIp,
		NetworkIp:    *subnet,
		PublicIp:     publicIp,
	}, nil
}
