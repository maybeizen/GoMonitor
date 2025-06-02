package collectors

import (
	"fmt"
	"monitor/models"

	"github.com/shirou/gopsutil/v3/net"
)

type NetworkCollector struct{}

func NewNetworkCollector() *NetworkCollector {
	return &NetworkCollector{}
}

func (n *NetworkCollector) Collect() ([]models.NetworkInfo, error) {
	interfaces, err := net.Interfaces()
	if err != nil {
		return nil, fmt.Errorf("failed to get network interfaces: %w", err)
	}

	counters, err := net.IOCounters(true) // per interface
	if err != nil {
		return nil, fmt.Errorf("failed to get network IO counters: %w", err)
	}

	counterMap := make(map[string]net.IOCountersStat)
	for _, c := range counters {
		counterMap[c.Name] = c
	}

	var networks []models.NetworkInfo
	for _, iface := range interfaces {
		if isLoopback(iface.Flags) {
			continue
		}

		counter, exists := counterMap[iface.Name]
		if !exists {
			continue
		}

		networks = append(networks, models.NetworkInfo{
			Name:        iface.Name,
			BytesSent:   counter.BytesSent,
			BytesRecv:   counter.BytesRecv,
			PacketsSent: counter.PacketsSent,
			PacketsRecv: counter.PacketsRecv,
			Errin:       counter.Errin,
			Errout:      counter.Errout,
			Dropin:      counter.Dropin,
			Dropout:     counter.Dropout,
		})
	}

	return networks, nil
}

func isLoopback(flags []string) bool {
	for _, flag := range flags {
		if flag == "loopback" {
			return true
		}
	}
	return false
} 