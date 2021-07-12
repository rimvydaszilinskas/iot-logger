package system

import (
	"fmt"
	"time"

	"github.com/rimvydaszilinskas/announcer-backend/models"
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
)

func CollectSystem() (*models.SystemDetails, error) {
	host, err := host.Info()

	if err != nil {
		return nil, fmt.Errorf("error retrieving info about host - %w", err)
	}

	cpuInfo, err := cpu.Info()

	if err != nil {
		return nil, fmt.Errorf("error retrieving info about CPU - %w", err)
	}

	perCore, err := cpu.Percent(time.Second, true)

	if err != nil {
		return nil, fmt.Errorf("error retrieving CPU per core usage - %w", err)
	}

	totalUsage, err := cpu.Percent(time.Second, false)

	if err != nil {
		return nil, fmt.Errorf("error retrieving total CPU usage - %w", err)
	}

	return &models.SystemDetails{
		OS:           host,
		CPU:          cpuInfo,
		TotalUsage:   totalUsage[0],
		PerCoreUsage: perCore,
	}, nil
}
