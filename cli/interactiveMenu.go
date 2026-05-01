package cli

import (
	"fmt"

	"go-who-ate-my-cpu/display"
	"go-who-ate-my-cpu/monitoring"
)

func RunInteractiveMenu(deviceInfo display.DeviceInfo) {
	for {
		display.DisplayPrompt(deviceInfo)

		var mode int

		_, err := fmt.Scan(&mode)
		if err != nil {
			fmt.Println("Please enter a number")
			continue
		}

		switch mode {
		case 1:
			fmt.Println("CPU Monitoring Mode: ")
			monitoring.MonitorCPU()
		case 2:
			fmt.Println("Memory Monitoring Mode: ")
			monitoring.MonitorMemory()
		case 3:
			fmt.Println("Network Monitoring Mode: ")
			monitoring.MonitorNetwork()
		case 5:
			fmt.Println("Hybrid Monitoring Mode: ")
			monitoring.MonitorComputer(10)
		default:
			fmt.Println("ERROR ERROR ERROR")
		}
	}
}
