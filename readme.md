# Go Monitor

A cross-platform system monitoring tool written in Go that collects system information and can output to files or API endpoints.

## Features

- Collects detailed system information:

  - CPU usage and details
  - Memory usage
  - Disk usage for all mounted volumes
  - System load (Unix/Linux)
  - Process count
  - Host information

- Multiple output options:

  - JSON file output
  - API endpoint (POST/PUT/PATCH)
  - Optional compression

- Cross-platform support:
  - Windows
  - Linux
  - macOS

## Installation

### Prerequisites

- Go 1.18 or newer

### Building from source

1. Clone the repository
2. Run `make deps` to install dependencies
3. Run `make build` to build the application

The binary will be created in the `bin` directory.

## Usage

Run the monitor with the default configuration:

```bash
# Linux/macOS
./bin/monitor

# Windows
bin\monitor.exe
```

### Configuration

Configuration is stored in `config.json`. The default configuration will be created if the file doesn't exist.

#### Configuration Options

```json
{
  "monitor_interval": 3,
  "outputs": [
    {
      "type": "file",
      "file_path": "data/data.json"
    },
    {
      "type": "api",
      "api_url": "http://example.com/api/metrics",
      "api_method": "POST",
      "api_key": "your-api-key"
    }
  ],
  "log_level": "info",
  "include_networks": true,
  "include_processes": true,
  "max_process_count": 1000,
  "enable_compression": false
}
```

- `monitor_interval`: How often to collect system information, in seconds
- `outputs`: Array of output configurations
  - File output:
    - `type`: "file"
    - `file_path`: Path to write the JSON data
  - API output:
    - `type`: "api"
    - `api_url`: URL of the API endpoint
    - `api_method`: HTTP method (POST, PUT, PATCH)
    - `api_key`: Optional API key for authentication
- `log_level`: Logging level (info, debug, warning, error)
- `include_networks`: Whether to collect network interface statistics
- `include_processes`: Whether to collect process count
- `max_process_count`: Maximum number of processes to count
- `enable_compression`: Whether to enable compression for file output (appends .gz to filename)

## Development

### Project Structure

```
monitor/
├── models/          # Data models
├── utils/           # Utility functions
│   ├── collectors/  # System information collectors
│   └── outputs/     # Output handlers
├── config.json      # Configuration file
├── go.mod           # Go module definition
├── main.go          # Main application entry point
└── Makefile         # Build automation
```

### Makefile Commands

- `make build`: Build the application
- `make clean`: Remove built binaries
- `make run`: Run the application
- `make deps`: Install dependencies
- `make test`: Run tests
- `make release`: Build optimized release binary
- `make config-api`: Create a sample config file with API endpoint

## License

MIT
