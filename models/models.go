package models

import "time"

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
	Available   uint64  `json:"available_bytes"`
	Free        uint64  `json:"free_bytes"`
}

type DiskInfo struct {
	Mountpoint  string  `json:"mountpoint"`
	Total       uint64  `json:"total_bytes"`
	Used        uint64  `json:"used_bytes"`
	UsedPercent float64 `json:"used_percent"`
	FileSystem  string  `json:"filesystem"`
}

type NetworkInfo struct {
	Name        string `json:"name"`
	BytesSent   uint64 `json:"bytes_sent"`
	BytesRecv   uint64 `json:"bytes_recv"`
	PacketsSent uint64 `json:"packets_sent"`
	PacketsRecv uint64 `json:"packets_recv"`
	Errin       uint64 `json:"errin"`
	Errout      uint64 `json:"errout"`
	Dropin      uint64 `json:"dropin"`
	Dropout     uint64 `json:"dropout"`
}

type LoadInfo struct {
	Load1  float64 `json:"load1"`
	Load5  float64 `json:"load5"`
	Load15 float64 `json:"load15"`
}

type OutputType string

const (
	FileOutput OutputType = "file"
	APIOutput OutputType = "api"
)

type OutputConfig struct {
	Type      OutputType `json:"type"`
	FilePath  string     `json:"file_path,omitempty"`
	APIURL    string     `json:"api_url,omitempty"`
	APIKey    string     `json:"api_key,omitempty"`
	APIMethod string     `json:"api_method,omitempty"`
}

type Config struct {
	MonitorInterval   int           `json:"monitor_interval"`
	Outputs           []OutputConfig `json:"outputs"`
	LogLevel          string        `json:"log_level"`
	IncludeNetworks   bool          `json:"include_networks"`
	IncludeProcesses  bool          `json:"include_processes"`
	MaxProcessCount   int           `json:"max_process_count"`
	EnableCompression bool          `json:"enable_compression"`
}

type SystemInfo struct {
	Hostname     string       `json:"hostname"`
	OS           string       `json:"os"`
	Platform     string       `json:"platform"`
	Uptime       uint64       `json:"uptime"`
	CPU          CPUInfo      `json:"cpu"`
	Memory       MemoryInfo   `json:"memory"`
	Disks        []DiskInfo   `json:"disks"`
	Networks     []NetworkInfo `json:"networks,omitempty"`
	Load         LoadInfo     `json:"load"`
	ProcessCount int          `json:"process_count"`
	Timestamp    string       `json:"timestamp"`
	CollectedAt  time.Time    `json:"collected_at"`
}
