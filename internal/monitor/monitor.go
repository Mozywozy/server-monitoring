package monitor

import (
	"fmt"
	"net"
	"time"
	"server-monitor/internal/alert"
	"server-monitor/internal/storage"
)

type Monitor struct {
	Servers  []string
	Status   map[string]string
	Interval time.Duration
}

func NewMonitor() *Monitor {
	return &Monitor{
		Servers:  []string{"192.168.1.1", "google.com"}, // Tambahkan alamat server yang ingin dipantau contoh google
		Status:   make(map[string]string),
		Interval: 10 * time.Second, 
	}
}

// PingServer mengirimkan ping ke server dan mengembalikan status
func PingServer(server string) string {
	_, err := net.LookupHost(server)
	if err != nil {
		return "DOWN"
	}
	conn, err := net.DialTimeout("ip4:icmp", server, 2*time.Second)
	if err != nil {
		return "DOWN"
	}
	defer conn.Close()
	return "UP"
}

// StartMonitoring memulai proses monitoring server
func (m *Monitor) StartMonitoring() {
	for {
		for _, server := range m.Servers {
			status := PingServer(server)
			previousStatus := m.Status[server]

			if status == "DOWN" && previousStatus != "DOWN" {
				alert.SendEmailAlert(server)
			}

			m.Status[server] = status
			storage.SaveServerStatus(server, status)
			fmt.Printf("Server %s is %s\n", server, status)
		}
		time.Sleep(m.Interval)
	}
}
