package main

import (
	"fmt"
	"os"
)

type DeviceInfo struct {
	OS                string
	Platform          string
	PlatformVersion   string
	CPU               string
	CPUUsage          float64
	CPUCores          int32
	DiskUsagePercent  float64
	MemoryUsedPercent float64
}

var deviceInfo DeviceInfo

func main() {
	err := getBasicDeviceInfo(&deviceInfo)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not get device info: ", err)
		os.Exit(1)
	}

	for {
		displayPrompt()

		var mode int

		_, err := fmt.Scan(&mode)
		if err != nil {
			fmt.Println("Please enter a number")
			continue
		}

		switch mode {
		case 1:
			fmt.Println("CPU Monitoring Mode: ")
			monitorCpu()
		case 2:
			fmt.Println("Memory Monitoring Mode: ")
			monitorMemorySnapshot()
		case 3:
			fmt.Println("Network Monitoring Mode: ")
			monitorNetwork()
		case 5:
			fmt.Println("Hybrid Monitoring Mode: ")
		default:
			fmt.Println("ERROR ERROR ERROR")
		}
	}
}
