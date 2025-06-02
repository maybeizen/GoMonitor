package collectors

import (
	"fmt"
	"monitor/models"
	"time"
)

type SystemCollector struct {
	config *models.Config
	cpu    *CPUCollector
	memory *MemoryCollector
	disk   *DiskCollector
	host   *HostCollector
	load   *LoadCollector
	process *ProcessCollector
	network *NetworkCollector
}

func NewSystemCollector(config *models.Config) *SystemCollector {
	return &SystemCollector{
		config:  config,
		cpu:     NewCPUCollector(),
		memory:  NewMemoryCollector(),
		disk:    NewDiskCollector(),
		host:    NewHostCollector(),
		load:    NewLoadCollector(),
		process: NewProcessCollector(config.MaxProcessCount),
		network: NewNetworkCollector(),
	}
}

func (s *SystemCollector) Collect() (*models.SystemInfo, error) {
	now := time.Now()
	timestamp := now.Format(time.RFC3339)

	hostInfo, err := s.host.Collect()
	if err != nil {
		return nil, fmt.Errorf("failed to collect host info: %w", err)
	}

	cpuInfo, err := s.cpu.Collect()
	if err != nil {
		return nil, fmt.Errorf("failed to collect CPU info: %w", err)
	}

	memInfo, err := s.memory.Collect()
	if err != nil {
		return nil, fmt.Errorf("failed to collect memory info: %w", err)
	}

	diskInfo, err := s.disk.Collect()
	if err != nil {
		fmt.Printf("Warning: Failed to collect disk info: %v\n", err)
		diskInfo = []models.DiskInfo{}
	}

	loadInfo, err := s.load.Collect()
	if err != nil {
		fmt.Printf("Warning: Failed to collect load info: %v\n", err)
		loadInfo = models.LoadInfo{}
	}

	processCount := 0
	if s.config.IncludeProcesses {
		count, err := s.process.Collect()
		if err != nil {
			fmt.Printf("Warning: Failed to collect process info: %v\n", err)
		} else {
			processCount = count
		}
	}

	var networks []models.NetworkInfo
	if s.config.IncludeNetworks {
		networks, err = s.network.Collect()
		if err != nil {
			fmt.Printf("Warning: Failed to collect network info: %v\n", err)
			networks = []models.NetworkInfo{}
		}
	}

	return &models.SystemInfo{
		Hostname:     hostInfo.Hostname,
		OS:           hostInfo.OS,
		Platform:     hostInfo.Platform,
		Uptime:       hostInfo.Uptime,
		CPU:          cpuInfo,
		Memory:       memInfo,
		Disks:        diskInfo,
		Networks:     networks,
		Load:         loadInfo,
		ProcessCount: processCount,
		Timestamp:    timestamp,
		CollectedAt:  now,
	}, nil
} 