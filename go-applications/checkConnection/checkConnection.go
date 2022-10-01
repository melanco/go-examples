package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	// var wg sync.WaitGroup
	ports := []string{"80", "443"}
	hosts := []string{"google.com", "192.168.64.6"}
	//fonctionne aussi avec DNS
	connect(hosts, ports)
}

func connect(hosts []string, ports []string) {
	conFailedCounter := 0
	for _, host := range hosts {
		for _, port := range ports {
			timeout := time.Second
			conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
			if err != nil {
				fmt.Println("Error:", err)
				conFailedCounter++
				fmt.Println("Connection failed:", conFailedCounter)
			}
			if conn != nil {
				defer conn.Close()
				fmt.Println("Opened", net.JoinHostPort(host, port))
			}
		}
	}
}
