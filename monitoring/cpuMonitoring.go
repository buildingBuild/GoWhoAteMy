package monitoring

import (
	"fmt"
	"sort"
	"time"

	"github.com/shirou/gopsutil/process"
)

type CPUSnapshot struct {
	PID        int32
	Name       string
	CPUPercent float64
	Time       time.Time
}

func monitorCPUSnapshot() ([]CPUSnapshot, error) {
	allProcesses, err := process.Processes()
	if err != nil {
		return nil, err
	}

	for _, p := range allProcesses {
		_, _ = p.CPUPercent()
	}

	time.Sleep(1 * time.Second)

	snapshots := []CPUSnapshot{}

	for _, p := range allProcesses {
		cpu, err := p.CPUPercent()
		if err != nil {
			continue
		}

		name, _ := p.Name()

		snapshots = append(snapshots, CPUSnapshot{
			PID:        p.Pid,
			Name:       name,
			CPUPercent: cpu,
			Time:       time.Now(),
		})
	}

	sort.Slice(snapshots, func(i, j int) bool {
		return snapshots[i].CPUPercent > snapshots[j].CPUPercent
	})

	return snapshots, nil
}

func detectSpikes(current []CPUSnapshot, previous map[int32]CPUSnapshot) {
	for _, cur := range current {
		prev, ok := previous[cur.PID]
		if !ok {
			continue
		}

		delta := cur.CPUPercent - prev.CPUPercent
		if delta >= 8.0 {
			message := fmt.Sprintf("%s CPU jumped from %.2f%% to %.2f%% (+%.2f%%)",
				cur.Name, prev.CPUPercent, cur.CPUPercent, delta)

			fmt.Printf("SPIKE: pid=%d name=%s prev=%.2f current=%.2f delta=%.2f\n",
				cur.PID, cur.Name, prev.CPUPercent, cur.CPUPercent, delta)
			go sendNotification("GoWhoAteMyCPU CPU Alert", message)
		}
	}
}

func MonitorCPU() {
	currentSnapshots, err := monitorCPUSnapshot()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	limit := 10
	if len(currentSnapshots) < limit {
		limit = len(currentSnapshots)
	}

	fmt.Println("Top CPU processes:")
	for _, s := range currentSnapshots[:limit] {
		fmt.Printf("pid=%d name=%s cpu=%.2f%%\n", s.PID, s.Name, s.CPUPercent)
	}
}
