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

func monitorCpu() {

	log.Println("Starting CPU monitor...")
	allProcesses, err := process.Processes()

	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, p := range allProcesses {
		p.CPUPercent()

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

	limit := 10
	if len(snapshots) < limit {
		limit = len(snapshots)
	}

	for _, s := range snapshots[:limit] {
		fmt.Println(s.PID, s.Name, s.CPUPercent)
	}
}
