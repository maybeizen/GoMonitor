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

const configPath = "config.json" // dont touch this

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile) // because debugging without file names is pain

	config, err := utils.LoadConfig(configPath)
	if err != nil {
		log.Fatalf("Failed to load config: %v", err) // this app dies like my dreams of Payge actually liking me back
	}

	ctx, cancel := context.WithCancel(context.Background()) // context is life. cancel it like my weekend plans.
	defer cancel()

	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM) // ah yes, UNIX signals. Nature's ctrl+c

	go func() {
		sig := <-sigChan
		log.Printf("Received signal %v, shutting down...", sig) // goodbye cruel world
		cancel()
	}()

	collector := collectors.NewSystemCollector(config) // it collects. nice

	var outputHandlers []outputs.OutputHandler
	for _, outputConfig := range config.Outputs {
		handler, err := outputs.NewOutputHandler(outputConfig)
		if err != nil {
			log.Printf("Warning: Failed to create output handler: %v", err) // passive-aggressive note
			continue
		}
		outputHandlers = append(outputHandlers, handler)
	}

	if len(outputHandlers) == 0 {
		log.Fatal("No valid output handlers configured") // config.json had one FUCKIGN JOB TF ETJJSOJN
	}

	log.Printf("Starting monitoring with interval: %ds", config.MonitorInterval)
	for _, handler := range outputHandlers {
		log.Printf("Output handler: %s", handler.Name()) // flexing all the handlers like I flex knowledge of Go to Payge (she doesnt care)
	}

	monitorLoop(ctx, collector, outputHandlers, config.MonitorInterval)

	log.Println("Shutdown complete") // fades out with 'Leave Out All The Rest' playing
}

func monitorLoop(ctx context.Context, collector *collectors.SystemCollector,
	handlers []outputs.OutputHandler, interval int) {

	ticker := time.NewTicker(time.Duration(interval) * time.Second) // tick tock, its bug o clock
	defer ticker.Stop()

	if err := collectAndWrite(collector, handlers); err != nil {
		log.Printf("Initial collection error: %v", err) // "its like im paranoid lookin over my logs"
	}

	for {
		select {
		case <-ctx.Done():
			return // peace out
		case <-ticker.C:
			if err := collectAndWrite(collector, handlers); err != nil {
				log.Printf("Error in monitoring cycle: %v", err) // everythign was perfectly fine until it wasnt
			}
		}
	}
}

func collectAndWrite(collector *collectors.SystemCollector, handlers []outputs.OutputHandler) error {
	log.Println("Collecting system information...") // just making sure the computer isnt dead yet

	info, err := collector.Collect()
	if err != nil {
		return fmt.Errorf("failed to collect system info: %w", err)
	}

	if len(info.Disks) == 0 {
		log.Println("Warning: No disk information collected. This may be normal if using restricted permissions.") // or if your disks just gave up like me in gym class
	} else {
		log.Printf("Collected information for %d disk(s)", len(info.Disks)) // more disk than i have friends
	}

	var wg sync.WaitGroup
	errChan := make(chan error, len(handlers))

	for _, handler := range handlers {
		wg.Add(1)
		go func(h outputs.OutputHandler) {
			defer wg.Done()

			start := time.Now()
			if err := h.Write(info); err != nil {
				errChan <- fmt.Errorf("%s: %v", h.Name(), err) // "i tried so hard and got so far, but in the end, the handler failed"
				return
			}

			elapsed := time.Since(start)
			log.Printf("Wrote system info to %s (took %v)", h.Name(), elapsed) // wow it actually worked?? wild
		}(handler)
	}

	wg.Wait()
	close(errChan)

	var errs []error
	for err := range errChan {
		errs = append(errs, err) // im gonna start harvesting my fuckin desk
	}

	if len(errs) > 0 {
		if len(errs) == 1 {
			return errs[0] // one error = bad but  many errors = Go moment
		}
		return fmt.Errorf("%v (and %d more errors)", errs[0], len(errs)-1) // and then everything just WENT TO SHIT
	}

	return nil // no errs? hell mustve froze over
}
