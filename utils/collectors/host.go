package collectors

import (
	"fmt"

	"github.com/shirou/gopsutil/v3/host"
)

type HostCollector struct{}

func NewHostCollector() *HostCollector {
	return &HostCollector{}
}

func (h *HostCollector) Collect() (*host.InfoStat, error) {
	hostInfo, err := host.Info()
	if err != nil {
		return nil, fmt.Errorf("failed to get host info: %w", err)
	}

	return hostInfo, nil
} 