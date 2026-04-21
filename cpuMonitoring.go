package main

import (
	"log"
	"os/exec"
	"time"
)

type CPUSnapshot struct {
	PID  int32
	Name string
	CPU  float64
	Time time.Time
} // will capture every second

func monitorCpu() {
	log.Println("Starting CPU monitor...")
	for {
		out, err := exec.Command("top", "-l1", "-s0").Output()

		if err != nil {
			log.Println("Error:", err)
			time.Sleep(1 * time.Second)
			continue
		}
		log.Printf("CPU Stats:\n%s", string(out))
		time.Sleep(1 * time.Second)
	}

}
