package models

import (
	"github.com/shirou/gopsutil/cpu"
	"github.com/shirou/gopsutil/host"
)

type SystemDetails struct {
	OS           *host.InfoStat `json:"os"`
	CPU          []cpu.InfoStat `json:"cpu"`
	TotalUsage   float64        `json:"total_usage"`
	PerCoreUsage []float64      `json:"per_core_usage"`
}