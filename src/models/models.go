package models

type CPUInfo struct {
	ModelName string  `json:"model_name"`
	VendorID  string  `json:"vendor_id"`
	Cores     int32   `json:"cores"`
	Mhz       float64 `json:"mhz"`
	Usage     float64 `json:"usage_percent"`
}

type MemoryInfo struct {
	Total       uint64  `json:"total_bytes"`
	Used        uint64  `json:"used_bytes"`
	UsedPercent float64 `json:"used_percent"`
}

type DiskInfo struct {
	Mountpoint  string  `json:"mountpoint"`
	Total       uint64  `json:"total_bytes"`
	Used        uint64  `json:"used_bytes"`
	UsedPercent float64 `json:"used_percent"`
}

type SystemInfo struct {
	Hostname     string     `json:"hostname"`
	OS           string     `json:"os"`
	Platform     string     `json:"platform"`
	CPU          CPUInfo    `json:"cpu"`
	Memory       MemoryInfo `json:"memory"`
	Disks        []DiskInfo `json:"disks"`
	ProcessCount int        `json:"process_count"`
	Timestamp    string     `json:"timestamp"`
}
