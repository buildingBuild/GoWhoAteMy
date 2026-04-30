package main

import "fmt"

func runSelectedMode(cli CommandLineOptions) bool {
	switch {
	case cli.CPU:
		fmt.Println("CPU Monitoring Mode: ")
		monitorCpu()
	case cli.Memory:
		fmt.Println("Memory Monitoring Mode: ")
		monitorMemory()
	case cli.Network:
		fmt.Println("Network Monitoring Mode: ")
		monitorNetwork()
	case cli.Computer:
		fmt.Println("Computer Monitoring Mode: ")
		monitorHybrid(cli.IntervalSeconds)
	default:
		return false
	}

	return true
}
