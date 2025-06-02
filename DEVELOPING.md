# Developing GoMonitor

This document contains information for developers who want to build, modify, or contribute to GoMonitor.

## Development Prerequisites

- Go 1.18 or newer
- Make (optional, for using the Makefile)

## Building from Source

1. Clone the repository
2. Install dependencies:
   ```
   go mod tidy
   ```
3. Build the application:
   ```
   go build -o bin/monitor.exe .
   ```

## Project Structure

```
monitor/
├── models/          # Data models
├── utils/           # Utility functions
│   ├── collectors/  # System information collectors
│   └── outputs/     # Output handlers
├── config.json      # Configuration file
├── go.mod           # Go module definition
├── main.go          # Main application entry point
└── build-release.bat # Release build script
```

## Configuration Options

The application uses a config.json file with the following options:

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

## Building

Simply run the following:

### Windows:

```bash
go build -o dist/gomonitor.exe
```

### Unix (Linux, macOS, etc):

```bash
go build -o dist/gomonitor
```

This will compile GoMonitor into a binary executable in the `/dist` directory.
