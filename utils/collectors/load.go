package collectors

import (
	"fmt"
	"monitor/models"
	"runtime"

	"github.com/shirou/gopsutil/v3/load"
)

type LoadCollector struct{}

func NewLoadCollector() *LoadCollector {
	return &LoadCollector{}
}

func (l *LoadCollector) Collect() (models.LoadInfo, error) {
	if runtime.GOOS == "windows" {
		return models.LoadInfo{
			Load1:  0,
			Load5:  0,
			Load15: 0,
		}, nil
	}

	loadStats, err := load.Avg()
	if err != nil {
		return models.LoadInfo{}, fmt.Errorf("failed to get load averages: %w", err)
	}

	return models.LoadInfo{
		Load1:  loadStats.Load1,
		Load5:  loadStats.Load5,
		Load15: loadStats.Load15,
	}, nil
} 