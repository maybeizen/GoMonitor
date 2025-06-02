package collectors

import (
	"fmt"
	"monitor/models"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
)

type CPUCollector struct{}

func NewCPUCollector() *CPUCollector {
	return &CPUCollector{}
}

func (c *CPUCollector) Collect() (models.CPUInfo, error) {
	cpuStats, err := cpu.Info()
	if err != nil {
		return models.CPUInfo{}, fmt.Errorf("failed to get CPU info: %w", err)
	}

	cpuPercent, err := cpu.Percent(time.Second, false)
	if err != nil {
		return models.CPUInfo{}, fmt.Errorf("failed to get CPU usage: %w", err)
	}

	var usage float64
	if len(cpuPercent) > 0 {
		usage = cpuPercent[0]
	}

	if len(cpuStats) == 0 {
		return models.CPUInfo{Usage: usage}, nil
	}

	return models.CPUInfo{
		ModelName: cpuStats[0].ModelName,
		VendorID:  cpuStats[0].VendorID,
		Cores:     cpuStats[0].Cores,
		Mhz:       cpuStats[0].Mhz,
		Usage:     usage,
	}, nil
} 