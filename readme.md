# Go Monitor

A simple, lightweight monitoring agent written in Go.

> **Work In Progress**  
> This project is actively being developed. Expect bugs, rapid changes, and evolving features.  
> Use it for internal testing and experimentation until a stable release is tagged.

## Features

- System hostname, OS, and platform info
- CPU info (model, usage %, core count)
- Memory usage (total, used, percent)
- Disk usage per mountpoint
- Total running processes
- Cross-platform support (Linux, macOS, Windows)

## Setup

1. **Clone the repo**:

```bash
   git clone https://github.com/your-org/go-monitor
   cd go-monitor
````

2. **Install dependencies**:

   ```bash
   go mod tidy
   ```

3. **Build and run the monitor**:

    ```bash
   go run ./src/main/main.go
   ```


## Project Structure

```
go-monitor/
├── data/               # JSON output written here
├── src/
│   ├── main.go         # Entry point
│   ├── utils.go        # Collection & file writing logic
│   └── models.go       # Structs for JSON structure
├── go.mod
└── README.md
```


## Output Example

```json
{
  "hostname": "my-server",
  "os": "linux",
  "platform": "ubuntu",
  "cpu": {
    "model_name": "Intel(R) Xeon(R)",
    "vendor_id": "GenuineIntel",
    "cores": 8,
    "mhz": 2400,
    "usage_percent": 12.3
  },
  "memory": {
    "total_bytes": 16777216000,
    "used_bytes": 8452346880,
    "used_percent": 50.35
  },
  "disks": [
    {
      "mountpoint": "/",
      "total_bytes": 256000000000,
      "used_bytes": 128000000000,
      "used_percent": 50.0
    }
  ],
  "process_count": 187,
  "timestamp": "2025-06-01T12:00:00Z"
}
```

## TODO

* [ ] Add network interface stats
* [ ] Optional JSON compression
* [ ] Send data to remote API
* [ ] Expose as REST or WebSocket server
* [ ] Add system uptime and user session count

## License

MIT – use freely, modify openly.

Made with ❤️ by maybeizen
