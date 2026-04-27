package main

import (
	"fmt"

	gnet "github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
)

type NetworkSnapshot struct {
}

func monitorNetwork() {
	processes, err := process.Processes()
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, p := range processes {
		conns, err := gnet.ConnectionsPid("tcp", p.Pid)
		fmt.Println(conns)
		if err != nil {
			continue
		}

		establishedConns := 0
		for _, conn := range conns {
			if conn.Status == "ESTABLISHED" {
				establishedConns++
			}
		}

		if establishedConns > 0 {
			name, _ := p.Name()
			fmt.Printf("pid=%d name=%s established_tcp=%d\n", p.Pid, name, establishedConns)
		}
	}
}
