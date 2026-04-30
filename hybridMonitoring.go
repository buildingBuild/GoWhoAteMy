package main

import (
	"fmt"
	"time"

	"github.com/shirou/gopsutil/mem"
)

func monitorHybrid() {
	fmt.Println("Hybrid Monitoring Mode running. Press Ctrl+C to stop.")

	previousCPUSnapshots := map[int32]CPUSnapshot{}
	previousMemorySnapshots := map[int32]MemorySnapshot{}
	previousNetworkPorts := map[string]bool{}
	previousNetworkConnections := map[int32]int{}

	for {
		fmt.Println("\n--- Hybrid Snapshot ---")
		monitorCpuHybrid(previousCPUSnapshots)
		monitorMemoryHybrid(previousMemorySnapshots)
		fmt.Println("\nNetwork activity:")
		monitorNetworkHybrid(previousNetworkPorts, previousNetworkConnections)

		time.Sleep(100 * time.Second)
	}
}

func monitorNetworkHybrid(previousPorts map[string]bool, previousConnections map[int32]int) {
	currentSnapshots, err := monitorNetworkSnapshot()
	if err != nil {
		fmt.Println("Network error:", err)
		return
	}

	detectNetworkChanges(currentSnapshots, previousPorts, previousConnections)
	printNetworkReport(currentSnapshots)
}

func monitorCpuHybrid(previousSnapshots map[int32]CPUSnapshot) {
	currentSnapshots, err := monitorCPUSnapshot()
	if err != nil {
		fmt.Println("CPU error:", err)
		return
	}

	detectSpikes(currentSnapshots, previousSnapshots)

	limit := 5
	if len(currentSnapshots) < limit {
		limit = len(currentSnapshots)
	}

	fmt.Println("\nTop CPU processes:")
	for _, s := range currentSnapshots[:limit] {
		fmt.Printf("pid=%d name=%s cpu=%.2f%%\n", s.PID, s.Name, s.CPUPercent)
	}

	for _, s := range currentSnapshots {
		previousSnapshots[s.PID] = s
	}
}

func monitorMemoryHybrid(previousSnapshots map[int32]MemorySnapshot) {
	virtualMemory, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("Memory error:", err)
		return
	}

	if virtualMemory.UsedPercent >= 85 {
		message := fmt.Sprintf("Memory usage is %.2f%%. Available memory is %.1f GB.",
			virtualMemory.UsedPercent,
			bytesToGB(virtualMemory.Available),
		)
		go sendNotification("GoWhoAteMyCPU Memory Pressure", message)
	}

	swapMemory, err := mem.SwapMemory()
	if err != nil {
		fmt.Println("Swap error:", err)
		return
	}

	if swapMemory.UsedPercent >= 50 {
		message := fmt.Sprintf("Swap is %.2f%% used. Your Mac may feel slower.", swapMemory.UsedPercent)
		go sendNotification("GoWhoAteMyCPU Swap Alert", message)
	}

	currentSnapshots, err := monitorMemorySnapshot()
	if err != nil {
		fmt.Println("Memory process error:", err)
		return
	}

	detectMemoryGrowth(currentSnapshots, previousSnapshots)
	printMemoryReport(currentSnapshots, virtualMemory.UsedPercent, virtualMemory.Available, swapMemory.UsedPercent, swapMemory.Used)

	for _, s := range currentSnapshots {
		previousSnapshots[s.PID] = s
	}
}
