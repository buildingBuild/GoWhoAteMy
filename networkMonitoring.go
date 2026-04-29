package main

import (
	"fmt"

	gnet "github.com/shirou/gopsutil/net"
	"github.com/shirou/gopsutil/process"
)

type UDPListeningPorts struct {
	Number int
	Text   string
}

type TCPListeningPorts struct {
	Number int
	Text   string
}

func monitorNetwork() {

	processes, err := process.Processes()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	for _, p := range processes {
		tcpConns, err := gnet.ConnectionsPid("tcp", p.Pid)
		if err != nil {
			continue
		}

		udpConns, err := gnet.ConnectionsPid("udp", p.Pid)
		if err != nil {
			continue
		}

		establishedTcpConns := 0
		establishedUdpConns := 0
		tcpListeningPorts := []TCPListeningPorts{}
		udpListeningPorts := []UDPListeningPorts{}

		for _, conn := range tcpConns {
			if conn.Status == "ESTABLISHED" {
				establishedTcpConns++
			}

			if conn.Status == "LISTEN" {
				tcpListeningPorts = append(tcpListeningPorts, TCPListeningPorts{
					Number: int(conn.Laddr.Port),
					Text:   conn.Laddr.IP,
				})
			}
		}

		for _, conn := range udpConns {
			if conn.Status == "ESTABLISHED" {
				establishedUdpConns++
			}

			if conn.Laddr.Port > 0 {
				udpListeningPorts = append(udpListeningPorts, UDPListeningPorts{
					Number: int(conn.Laddr.Port),
					Text:   conn.Laddr.IP,
				})
			}
		}

		if establishedTcpConns > 0 {
			name, _ := p.Name()
			fmt.Printf("pid=%d name=%s established_tcp=%d\n", p.Pid, name, establishedTcpConns)
		}

		if establishedUdpConns > 0 {
			name, _ := p.Name()
			fmt.Printf("pid=%d name=%s established_udp=%d\n", p.Pid, name, establishedUdpConns)
		}

		if len(tcpListeningPorts) > 0 {
			name, _ := p.Name()
			for _, port := range tcpListeningPorts {
				fmt.Printf("pid=%d name=%s listening for TCP traffic on port %d from %s\n", p.Pid, name, port.Number, port.Text)
			}
		}

		if len(udpListeningPorts) > 0 {
			name, _ := p.Name()
			for _, port := range udpListeningPorts {
				fmt.Printf("pid=%d name=%s listening for UDP traffic on port %d from %s\n", p.Pid, name, port.Number, port.Text)
			}
		}
	}
}
