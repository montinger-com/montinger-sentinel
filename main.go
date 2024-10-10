package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/montinger-com/montinger-sentinel/config"
	"github.com/montinger-com/montinger-sentinel/models"
	"github.com/rashintha/logger"
	"github.com/shirou/gopsutil/v4/cpu"
	"github.com/shirou/gopsutil/v4/host"
	"github.com/shirou/gopsutil/v4/mem"
)

func main() {
	cpuChannel := make(chan float64)
	memChannel := make(chan *mem.VirtualMemoryStat)
	osChannel := make(chan *host.InfoStat)

	for {
		go getCPU(cpuChannel)
		go getMemory(memChannel)
		go getOS(osChannel)

		cpuUsage := <-cpuChannel
		mem := <-memChannel
		os := <-osChannel

		monitor := models.Monitor{
			LastData: models.LastData{
				CPU: &models.CPU{
					UsedPercent: cpuUsage,
				},
				Memory: &models.Memory{
					Total:       mem.Total,
					Available:   mem.Available,
					Used:        mem.Used,
					UsedPercent: mem.UsedPercent,
				},
				OS: &models.OS{
					Type:            os.OS,
					Platform:        os.Platform,
					PlatformFamily:  os.PlatformFamily,
					PlatformVersion: os.PlatformVersion,
					KernelVersion:   os.KernelVersion,
					KernelArch:      os.KernelArch,
				},
				Uptime: os.Uptime,
			},
		}

		url := fmt.Sprintf("%v/monitors/%v/push", config.API_URL, config.UID)
		requestBody, err := json.Marshal(monitor)
		if err != nil {
			logger.Errorf("Error in converting to json: %v", err.Error())
			continue
		}

		req, err := http.NewRequest("POST", url, bytes.NewBuffer(requestBody))
		if err != nil {
			logger.Errorf("Error in creating request: %v", err.Error())
			continue
		}

		// Set headers
		req.Header.Set("Content-Type", "application/json")
		req.Header.Set("X-API-KEY", config.SECRET)

		// Send the request
		client := &http.Client{}
		resp, err := client.Do(req)
		if err != nil {
			logger.Errorf("Error in sending request: %v", err.Error())
			continue
		}
		defer resp.Body.Close()

		if resp.StatusCode != 200 {
			logger.Errorf("response Status: %v", resp.Status)
			logger.Infof("response Headers: %v", resp.Header)
		}

		time.Sleep(15 * time.Second)
	}
}

func getCPU(cpuChannel chan float64) {
	percentages, err := cpu.Percent(0, false)
	if err != nil {
		logger.Errorf("Error in getting CPU percentage: %v", err.Error())
	}
	cpuChannel <- percentages[0]
}

func getMemory(memChannel chan *mem.VirtualMemoryStat) {
	v, err := mem.VirtualMemory()
	if err != nil {
		logger.Errorf("Error in getting memory details: %v", err.Error())
	}
	memChannel <- v
}

func getOS(osChannel chan *host.InfoStat) {
	os, err := host.Info()
	if err != nil {
		logger.Errorf("Error in getting os details: %v", err.Error())
	}
	osChannel <- os
}
