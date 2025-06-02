package collectors

import (
	"fmt"
	"monitor/models"

	"github.com/shirou/gopsutil/v3/mem"
)

type MemoryCollector struct{}

func NewMemoryCollector() *MemoryCollector {
	return &MemoryCollector{}
}

func (m *MemoryCollector) Collect() (models.MemoryInfo, error) {
	memStats, err := mem.VirtualMemory()
	if err != nil {
		return models.MemoryInfo{}, fmt.Errorf("failed to get memory info: %w", err)
	}

	return models.MemoryInfo{
		Total:       memStats.Total,
		Used:        memStats.Used,
		UsedPercent: memStats.UsedPercent,
		Available:   memStats.Available,
		Free:        memStats.Free,
	}, nil
} 