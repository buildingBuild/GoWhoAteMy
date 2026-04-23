package main

import (
	"fmt"
	"log"
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
		if delta >= 15.0 {
			fmt.Printf("SPIKE: pid=%d name=%s prev=%.2f current=%.2f delta=%.2f\n",
				cur.PID, cur.Name, prev.CPUPercent, cur.CPUPercent, delta)
		}
	}
}

func monitorCpu() {
	log.Println("Starting CPU monitor...")
	previousSnapshots := map[int32]CPUSnapshot{}

	for {

		currentSnapshots, err := monitorCPUSnapshot()
		if err != nil {
			fmt.Println("Error:", err)
			return
		}

		detectSpikes(currentSnapshots, previousSnapshots)

		limit := 10
		if len(currentSnapshots) < limit {
			limit = len(currentSnapshots)
		}

		for _, s := range currentSnapshots[:limit] {
			fmt.Println(s.PID, s.Name, s.CPUPercent)
		}

		previousSnapshots = make(map[int32]CPUSnapshot)
		for _, s := range currentSnapshots {
			previousSnapshots[s.PID] = s
		}

		time.Sleep(10 * time.Second)
	}
}
