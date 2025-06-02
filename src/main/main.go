package main

import (
	"fmt"
	"log"
	_ "monitor/src/models"
	"monitor/src/utils"
	"time"
)

func main() {
	log.Println("Starting system monitor...")
	for {
		info, err := utils.CollectSystemInfo()
		if err != nil {
			log.Println("Error collecting system info:", err)
			continue
		}

		err = utils.WriteJSONToFile(info)
		if err != nil {
			log.Println("Error writing to file:", err)
		} else {
			fmt.Println("Wrote system info to /logs/data.json")
		}

		time.Sleep(3 * time.Second)
	}
}
