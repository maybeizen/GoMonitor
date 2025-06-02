package utils

import (
	"encoding/json"
	"monitor/src/models"
	"os"
	"path/filepath"
	"time"

	"github.com/shirou/gopsutil/v3/cpu"
	"github.com/shirou/gopsutil/v3/disk"
	"github.com/shirou/gopsutil/v3/host"
	"github.com/shirou/gopsutil/v3/mem"
	"github.com/shirou/gopsutil/v3/process"
)

func CollectSystemInfo() (*models.SystemInfo, error) {
	hostInfo, _ := host.Info()
	cpuStats, _ := cpu.Info()
	cpuPercent, _ := cpu.Percent(0, false)
	memStats, _ := mem.VirtualMemory()
	procs, _ := process.Processes()
	timestamp := time.Now().Format(time.RFC3339)
	
	cpuData := models.CPUInfo{}
	if len(cpuStats) > 0 {
		cpuData = models.CPUInfo{
			ModelName: cpuStats[0].ModelName,
			VendorID:  cpuStats[0].VendorID,
			Cores:     cpuStats[0].Cores,
			Mhz:       cpuStats[0].Mhz,
			Usage:     cpuPercent[0],
		}
	}

	memData := models.MemoryInfo{
		Total:       memStats.Total,
		Used:        memStats.Used,
		UsedPercent: memStats.UsedPercent,
	}

	partitions, _ := disk.Partitions(false)
	var diskData []models.DiskInfo
	for _, p := range partitions {
		usage, err := disk.Usage(p.Mountpoint)
		if err == nil {
			diskData = append(diskData, models.DiskInfo{
				Mountpoint:  usage.Path,
				Total:       usage.Total,
				Used:        usage.Used,
				UsedPercent: usage.UsedPercent,
			})
		}
	}

	return &models.SystemInfo{
		Hostname:     hostInfo.Hostname,
		OS:           hostInfo.OS,
		Platform:     hostInfo.Platform,
		CPU:          cpuData,
		Memory:       memData,
		Disks:        diskData,
		ProcessCount: len(procs),
		Timestamp:    timestamp,
	}, nil
}

func WriteJSONToFile(info *models.SystemInfo) error {
	data, err := json.MarshalIndent(info, "", "  ")
	if err != nil {
		return err
	}

	cwd, err := os.Getwd()
	if err != nil {
		return err
	}

	outputDir := filepath.Join(cwd, "data")
	err = os.MkdirAll(outputDir, 0755)
	if err != nil {
		return err
	}

	filePath := filepath.Join(outputDir, "data.json")
	return os.WriteFile(filePath, data, 0644)
}
