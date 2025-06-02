# GoMonitor

A lightweight, cross-platform system monitoring tool that collects system information and can output to files or API endpoints.

![GitHub release (latest by date)](https://img.shields.io/github/v/release/maybeizen/gomonitor)
![License](https://img.shields.io/github/license/maybeizen/gomonitor)

## Features

- **Cross-Platform**: Works on Windows, Linux, and macOS
- **Detailed System Information**:
  - CPU usage and details
  - Memory usage
  - Disk usage for all mounted volumes
  - Network interface statistics
  - System load (Unix/Linux)
  - Process count
- **Multiple Output Options**:
  - Save to JSON files
  - Send to REST API endpoints
  - Optional compression

## Download

Download the latest release for your platform from the [Releases page](https://github.com/maybeizen/gomonitor/releases).

Available platforms:

- Windows (x64 and x86)
- Linux (x64 and ARM64)
- macOS (Intel and Apple Silicon)

## Quick Start

1. Download the appropriate zip file for your platform
2. Extract the zip file
3. Run the monitor executable:

```bash
# Windows
gomonitor-1.0.0-windows-amd64.exe

# Linux/macOS
chmod +x gomonitor-1.0.0-linux-amd64  # Make executable first
./gomonitor-1.0.0-linux-amd64
```

On first run, a default `config.json` file will be created.

## Configuration

The monitor uses a `config.json` file in the same directory as the executable. Edit this file to:

- Change the monitoring frequency
- Configure where to save the data
- Set up API endpoints
- Enable/disable certain collectors

Example config for saving to a file and sending to an API:

```json
{
  "monitor_interval": 10,
  "outputs": [
    {
      "type": "file",
      "file_path": "data/system-stats.json"
    },
    {
      "type": "api",
      "api_url": "https://your-monitoring-api.com/metrics",
      "api_method": "POST",
      "api_key": "your-api-key"
    }
  ],
  "include_networks": true
}
```

## Windows Users

For complete disk information, you may need to run the monitor as Administrator.

## For Developers

If you want to build from source or contribute to the project, see [DEVELOPING.md](DEVELOPING.md).

## License

[MIT](license)
