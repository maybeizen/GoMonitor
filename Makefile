.PHONY: build clean run deps test release

# Detect OS
ifeq ($(OS),Windows_NT)
	BINARY_NAME = monitor.exe
	BINARY_PATH = bin\$(BINARY_NAME)
	MKDIR = if not exist bin mkdir bin
	RM = if exist $(BINARY_PATH) del /Q $(BINARY_PATH)
	GO = go
	RUN_PREFIX = 
else
	BINARY_NAME = monitor
	BINARY_PATH = bin/$(BINARY_NAME)
	MKDIR = mkdir -p bin
	RM = rm -f $(BINARY_PATH)
	GO = go
	RUN_PREFIX = ./
endif

# Default target
all: clean build

# Build the application
build:
	@echo "Building..."
	@$(MKDIR)
	$(GO) build -o $(BINARY_PATH) .

# Run the application
run:
	@echo "Running..."
	$(RUN_PREFIX)$(BINARY_PATH)

# Clean the binary
clean:
	@echo "Cleaning..."
	@$(RM)

# Install dependencies
deps:
	@echo "Installing dependencies..."
	$(GO) mod tidy

# Build for release
release: clean
	@echo "Building release version..."
	@$(MKDIR)
	$(GO) build -ldflags="-s -w" -o $(BINARY_PATH) .

# Run tests
test:
	@echo "Running tests..."
	$(GO) test ./...

# Sample config with API endpoint
config-api:
	@echo "Creating sample config with API endpoint..."
	@echo '{"monitor_interval":3,"outputs":[{"type":"file","file_path":"data/data.json"},{"type":"api","api_url":"http://example.com/api/metrics","api_method":"POST","api_key":"your-api-key"}],"log_level":"info","include_networks":true,"include_processes":true,"max_process_count":1000,"enable_compression":false}' > config.json

# Print variables (for debugging)
debug:
	@echo OS: $(OS)
	@echo BINARY_NAME: $(BINARY_NAME)
	@echo BINARY_PATH: $(BINARY_PATH)
	@echo GO: $(GO)
