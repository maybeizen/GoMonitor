package main

import (
	"context"
	"fmt"
	"log"
	"monitor/utils"
	"monitor/utils/collectors"
	"monitor/utils/outputs"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

const configPath = "config.json"

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	config, err := utils.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err)
	}

	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		sig := <-sigChan
		log.Printf("Received signal %v, shutting down...", sig)
		cancel()
	}()

	collector := collectors.NewSystemCollector(config)

	var outputHandlers []outputs.OutputHandler
	for _, outputConfig := range config.Outputs {
		handler, err := outputs.NewOutputHandler(outputConfig)
		if err != nil {
			log.Printf("Warning: Failed to create output handler: %v", err)
			continue
		}
		outputHandlers = append(outputHandlers, handler)
	}

	if len(outputHandlers) == 0 {
		log.Fatal("No valid output handlers configured")
	}

	log.Printf("Starting monitoring with interval: %ds", config.MonitorInterval)
	for _, handler := range outputHandlers {
		log.Printf("Output handler: %s", handler.Name())
	}

	monitorLoop(ctx, collector, outputHandlers, config.MonitorInterval)

	log.Println("Shutdown complete")
}

func monitorLoop(ctx context.Context, collector *collectors.SystemCollector, 
	handlers []outputs.OutputHandler, interval int) {
	
	ticker := time.NewTicker(time.Duration(interval) * time.Second)
	defer ticker.Stop()

	if err := collectAndWrite(collector, handlers); err != nil {
		log.Printf("Initial collection error: %v", err)
	}

	for {
		select {
		case <-ctx.Done():
			return
		case <-ticker.C:
			if err := collectAndWrite(collector, handlers); err != nil {
				log.Printf("Error in monitoring cycle: %v", err)
			}
		}
	}
}

func collectAndWrite(collector *collectors.SystemCollector, handlers []outputs.OutputHandler) error {
	log.Println("Collecting system information...")

	info, err := collector.Collect()
	if err != nil {
		return fmt.Errorf("failed to collect system info: %w", err)
	}

	if len(info.Disks) == 0 {
		log.Println("Warning: No disk information collected. This may be normal if using restricted permissions.")
	} else {
		log.Printf("Collected information for %d disk(s)", len(info.Disks))
	}

	var wg sync.WaitGroup
	errChan := make(chan error, len(handlers))

	for _, handler := range handlers {
		wg.Add(1)
		go func(h outputs.OutputHandler) {
			defer wg.Done()
			
			start := time.Now()
			if err := h.Write(info); err != nil {
				errChan <- fmt.Errorf("%s: %v", h.Name(), err)
				return
			}
			
			elapsed := time.Since(start)
			log.Printf("Wrote system info to %s (took %v)", h.Name(), elapsed)
		}(handler)
	}

	wg.Wait()
	close(errChan)

	var errs []error
	for err := range errChan {
		errs = append(errs, err)
	}

	if len(errs) > 0 {
		if len(errs) == 1 {
			return errs[0]
		}
		return fmt.Errorf("%v (and %d more errors)", errs[0], len(errs)-1)
	}

	return nil
}
