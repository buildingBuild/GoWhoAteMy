package main

import "fmt"

func runInteractiveMenu() {
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
			monitorMemory()
		case 3:
			fmt.Println("Network Monitoring Mode: ")
			monitorNetwork()
		case 5:
			fmt.Println("Hybrid Monitoring Mode: ")
			monitorHybrid(10)
		default:
			fmt.Println("ERROR ERROR ERROR")
		}
	}
}
