package collectors

import (
	"fmt"
	"monitor/models"
	"runtime"
	"sort"
	"strings"

	"github.com/shirou/gopsutil/v3/disk"
)

type DiskCollector struct{}

func NewDiskCollector() *DiskCollector {
	return &DiskCollector{}
}

func (d *DiskCollector) Collect() ([]models.DiskInfo, error) {
	var partitions []disk.PartitionStat
	var err error

	if runtime.GOOS == "windows" {
		partitions, err = disk.Partitions(false)
	} else {
		partitions, err = disk.Partitions(true)
	}
	
	if err != nil {
		return nil, fmt.Errorf("failed to get disk partitions: %w", err)
	}

	if len(partitions) == 0 {
		return nil, fmt.Errorf("no disk partitions found")
	}

	var diskData []models.DiskInfo
	for _, p := range partitions {
		if !isValidPartition(p) {
			continue
		}

		usage, err := disk.Usage(p.Mountpoint)
		if err != nil {
			fmt.Printf("Warning: Could not access disk %s: %v\n", p.Mountpoint, err)
			continue
		}

		if usage.Total == 0 {
			continue
		}

		diskData = append(diskData, models.DiskInfo{
			Mountpoint:  usage.Path,
			Total:       usage.Total,
			Used:        usage.Used,
			UsedPercent: usage.UsedPercent,
			FileSystem:  p.Fstype,
		})
	}

	if len(diskData) == 0 {
		fmt.Println("No disk information could be collected, returning empty disk list")
		return []models.DiskInfo{}, nil
	}

	sort.Slice(diskData, func(i, j int) bool {
		return diskData[i].Mountpoint < diskData[j].Mountpoint
	})

	return diskData, nil
}

func isValidPartition(p disk.PartitionStat) bool {
	if runtime.GOOS == "windows" {
		if strings.HasPrefix(p.Mountpoint, "\\\\") {
			return false
		}
		if p.Mountpoint == "" || len(p.Mountpoint) < 2 {
			return false
		}
		if len(p.Mountpoint) >= 2 && p.Mountpoint[1] != ':' {
			return false
		}
		
		isRemovable := false
		for _, opt := range p.Opts {
			if strings.Contains(strings.ToLower(opt), "removable") {
				isRemovable = true
				break
			}
		}
		return !isRemovable
	}

	if p.Mountpoint == "" {
		return false
	}

	virtualFS := []string{"devfs", "tmpfs", "devtmpfs", "proc", "sysfs", "cgroup"}
	for _, fs := range virtualFS {
		if p.Fstype == fs {
			return false
		}
	}
	
	return true
} 