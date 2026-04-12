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
	var mode int
	err := getBasicDeviceInfo(&deviceInfo)
	if err != nil {
		fmt.Fprintln(os.Stderr, "Could not get device info: ", err)
		os.Exit(1)
	}

	fmt.Println("\nWelcome to GoWhoAteMyCPU!")
	fmt.Println("")
	fmt.Println("More info on features:")
	fmt.Println("https://github.com/buildingBuild/GoWhoAteMyCPU")
	fmt.Println("")
	displayDeviceInfo(deviceInfo)

	fmt.Println(Green + "\nOPTIONS: " + Reset)
	fmt.Println("1. CPU Monitoring Mode:")
	fmt.Println("2. Memory Monitoring Mode:")
	fmt.Println("3. Network Monitoring Mode:")
	fmt.Println("5. Hybrid Monitoring Mode:")
	fmt.Print("Select a Mode: ")
	fmt.Scan(&mode)
}
