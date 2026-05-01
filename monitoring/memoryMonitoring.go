package monitoring

import (
	"fmt"
	"sort"
	"time"

	"github.com/shirou/gopsutil/mem"
	"github.com/shirou/gopsutil/process"
)

type MemorySnapshot struct {
	PID           int32
	Name          string
	RSSBytes      uint64
	VMSBytes      uint64
	SwapBytes     uint64
	MemoryPercent float32
	Time          time.Time
}

func monitorMemorySnapshot() ([]MemorySnapshot, error) {
	allProcesses, err := process.Processes()
	if err != nil {
		return nil, err
	}

	snapshots := []MemorySnapshot{}

	for _, p := range allProcesses {
		memoryInfo, err := p.MemoryInfo()
		if err != nil {
			continue
		}

		memoryPercent, err := p.MemoryPercent()
		if err != nil {
			continue
		}

		name, _ := p.Name()

		snapshots = append(snapshots, MemorySnapshot{
			PID:           p.Pid,
			Name:          name,
			RSSBytes:      memoryInfo.RSS,
			VMSBytes:      memoryInfo.VMS,
			SwapBytes:     memoryInfo.Swap,
			MemoryPercent: memoryPercent,
			Time:          time.Now(),
		})
	}

	sort.Slice(snapshots, func(i, j int) bool {
		return snapshots[i].RSSBytes > snapshots[j].RSSBytes
	})

	return snapshots, nil
}

func detectMemoryGrowth(current []MemorySnapshot, previous map[int32]MemorySnapshot) {
	for _, cur := range current {
		prev, ok := previous[cur.PID]
		if !ok {
			continue
		}

		if cur.RSSBytes <= prev.RSSBytes {
			continue
		}

		growth := cur.RSSBytes - prev.RSSBytes
		if growth >= 100*1024*1024 {
			message := fmt.Sprintf("%s grew by %.1f MB", cur.Name, bytesToMB(growth))

			fmt.Printf("MEMORY GROWTH: pid=%d name=%s grew by %.1f MB\n",
				cur.PID, cur.Name, bytesToMB(growth))
			go sendNotification("GoWhoAteMyCPU Memory Alert", message)
		}
	}
}

func MonitorMemory() {
	virtualMemory, err := mem.VirtualMemory()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	swapMemory, err := mem.SwapMemory()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	currentSnapshots, err := monitorMemorySnapshot()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	printMemoryReport(currentSnapshots, virtualMemory.UsedPercent, virtualMemory.Available, swapMemory.UsedPercent, swapMemory.Used)
}

func bytesToMB(bytes uint64) float64 {
	return float64(bytes) / 1024 / 1024
}

func bytesToGB(bytes uint64) float64 {
	return float64(bytes) / 1024 / 1024 / 1024
}

func printMemoryReport(currentSnapshots []MemorySnapshot, memoryUsedPercent float64, memoryAvailable uint64, swapUsedPercent float64, swapUsed uint64) {
	fmt.Printf("\nSystem memory: %.2f%% used, %.1f GB available\n",
		memoryUsedPercent,
		bytesToGB(memoryAvailable),
	)
	fmt.Printf("Swap: %.2f%% used, %.1f GB used\n",
		swapUsedPercent,
		bytesToGB(swapUsed),
	)
	fmt.Println("Context: Swap means disk space used as backup memory when RAM is tight. High swap can make the computer feel slow.")

	limit := 10
	if len(currentSnapshots) < limit {
		limit = len(currentSnapshots)
	}

	fmt.Println("\nTop memory processes:")
	for _, s := range currentSnapshots[:limit] {
		fmt.Printf("pid=%d name=%s memory=%.1f MB ram=%.2f%% swap=%.1f MB\n",
			s.PID,
			s.Name,
			bytesToMB(s.RSSBytes),
			s.MemoryPercent,
			bytesToMB(s.SwapBytes),
		)
	}
}
