package collectors

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/process"
)

type ProcessCollector struct {
	MaxCount int
}

func NewProcessCollector(maxCount int) *ProcessCollector {
	return &ProcessCollector{
		MaxCount: maxCount,
	}
}

func (p *ProcessCollector) Collect() (int, error) {
	processes, err := process.Processes()
	if err != nil {
		return 0, fmt.Errorf("failed to get process list: %w", err)
	}

	count := len(processes)
	if count > p.MaxCount {
		count = p.MaxCount
	}

	return count, nil
} 