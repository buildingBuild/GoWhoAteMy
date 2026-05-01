package monitoring

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

type NetworkSnapshot struct {
	PID                 int32
	Name                string
	EstablishedTCPConns int
	EstablishedUDPConns int
	TCPListeningPorts   []TCPListeningPorts
	UDPListeningPorts   []UDPListeningPorts
}

type NetworkPortEvent struct {
	PID      int32
	Name     string
	Protocol string
	Port     int
	Address  string
}

func monitorNetworkSnapshot() ([]NetworkSnapshot, error) {
	processes, err := process.Processes()
	if err != nil {
		return nil, err
	}

	snapshots := []NetworkSnapshot{}

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

		if establishedTcpConns == 0 &&
			establishedUdpConns == 0 &&
			len(tcpListeningPorts) == 0 &&
			len(udpListeningPorts) == 0 {
			continue
		}

		name, _ := p.Name()
		snapshots = append(snapshots, NetworkSnapshot{
			PID:                 p.Pid,
			Name:                name,
			EstablishedTCPConns: establishedTcpConns,
			EstablishedUDPConns: establishedUdpConns,
			TCPListeningPorts:   tcpListeningPorts,
			UDPListeningPorts:   udpListeningPorts,
		})
	}

	return snapshots, nil
}

func MonitorNetwork() {
	snapshots, err := monitorNetworkSnapshot()
	if err != nil {
		fmt.Println("error:", err)
		return
	}

	printNetworkReport(snapshots)
}

func printNetworkReport(snapshots []NetworkSnapshot) {
	for _, snapshot := range snapshots {
		if snapshot.EstablishedTCPConns > 0 {
			fmt.Printf("pid=%d name=%s established_tcp=%d\n", snapshot.PID, snapshot.Name, snapshot.EstablishedTCPConns)
		}

		if snapshot.EstablishedUDPConns > 0 {
			fmt.Printf("pid=%d name=%s established_udp=%d\n", snapshot.PID, snapshot.Name, snapshot.EstablishedUDPConns)
		}

		if len(snapshot.TCPListeningPorts) > 0 {
			for _, port := range snapshot.TCPListeningPorts {
				fmt.Printf("pid=%d name=%s listening for TCP traffic on port %d from %s\n", snapshot.PID, snapshot.Name, port.Number, port.Text)
			}
		}

		if len(snapshot.UDPListeningPorts) > 0 {
			for _, port := range snapshot.UDPListeningPorts {
				fmt.Printf("pid=%d name=%s listening for UDP traffic on port %d from %s\n", snapshot.PID, snapshot.Name, port.Number, port.Text)
			}
		}
	}
}

func detectNetworkChanges(current []NetworkSnapshot, previousPorts map[string]bool, previousConnections map[int32]int) {
	currentPorts := map[string]NetworkPortEvent{}
	currentConnections := map[int32]int{}
	processNames := map[int32]string{}
	hasPreviousPorts := len(previousPorts) > 0
	hasPreviousConnections := len(previousConnections) > 0

	for _, snapshot := range current {
		currentConnections[snapshot.PID] = snapshot.EstablishedTCPConns + snapshot.EstablishedUDPConns
		processNames[snapshot.PID] = snapshot.Name

		for _, port := range snapshot.TCPListeningPorts {
			event := NetworkPortEvent{
				PID:      snapshot.PID,
				Name:     snapshot.Name,
				Protocol: "TCP",
				Port:     port.Number,
				Address:  port.Text,
			}
			currentPorts[networkPortKey(event)] = event
		}

		for _, port := range snapshot.UDPListeningPorts {
			event := NetworkPortEvent{
				PID:      snapshot.PID,
				Name:     snapshot.Name,
				Protocol: "UDP",
				Port:     port.Number,
				Address:  port.Text,
			}
			currentPorts[networkPortKey(event)] = event
		}
	}

	for key, event := range currentPorts {
		if hasPreviousPorts && !previousPorts[key] {
			message := fmt.Sprintf("%s started listening for %s traffic on port %d from %s",
				event.Name, event.Protocol, event.Port, event.Address)
			fmt.Println("NETWORK CHANGE:", message)
			go sendNotification("GoWhoAteMyCPU Network Alert", message)
		}
	}

	for pid, count := range currentConnections {
		previousCount := previousConnections[pid]
		if hasPreviousConnections && count > previousCount {
			name := processNames[pid]
			if name == "" {
				name = fmt.Sprintf("pid=%d", pid)
			}

			message := fmt.Sprintf("%s opened %d new established network connection(s)", name, count-previousCount)
			fmt.Println("NETWORK CHANGE:", message)
			go sendNotification("GoWhoAteMyCPU Connection Alert", message)
		}
	}

	replaceStringSet(previousPorts, currentPorts)
	replaceIntMap(previousConnections, currentConnections)
}

func networkPortKey(event NetworkPortEvent) string {
	return fmt.Sprintf("%d|%s|%d|%s", event.PID, event.Protocol, event.Port, event.Address)
}

func replaceStringSet(target map[string]bool, source map[string]NetworkPortEvent) {
	for key := range target {
		delete(target, key)
	}

	for key := range source {
		target[key] = true
	}
}

func replaceIntMap(target map[int32]int, source map[int32]int) {
	for key := range target {
		delete(target, key)
	}

	for key, value := range source {
		target[key] = value
	}
}
