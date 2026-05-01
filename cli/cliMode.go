package cli

import (
	"fmt"

	"go-who-ate-my-cpu/monitoring"
)

type CommandLineOptions struct {
	CPU             bool `help:"Run CPU monitoring mode."`
	Memory          bool `help:"Run memory monitoring mode."`
	Network         bool `help:"Run network monitoring mode."`
	Computer        bool `help:"Run full computer monitoring mode with notifications."`
	IntervalSeconds int  `name:"interval" default:"10" help:"Seconds between computer checks."`
}

func RunSelectedMode(options CommandLineOptions) bool {
	switch {
	case options.CPU:
		fmt.Println("CPU Monitoring Mode: ")
		monitoring.MonitorCPU()
	case options.Memory:
		fmt.Println("Memory Monitoring Mode: ")
		monitoring.MonitorMemory()
	case options.Network:
		fmt.Println("Network Monitoring Mode: ")
		monitoring.MonitorNetwork()
	case options.Computer:
		fmt.Println("Computer Monitoring Mode: ")
		monitoring.MonitorComputer(options.IntervalSeconds)
	default:
		return false
	}

	return true
}
