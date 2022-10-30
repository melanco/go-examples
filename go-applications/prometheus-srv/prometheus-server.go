package main

import (
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"time"

	"github.com/prometheus/client_golang/prometheus"
	"github.com/prometheus/client_golang/prometheus/promhttp"
)

var (

	//Create a new prometheus gauge
	errored = prometheus.NewGaugeVec(
		prometheus.GaugeOpts{
			Name: "validate_errored",
		},
		[]string{"module", "error"},
	)
)

func main() {

	rand.Seed(time.Now().Unix())
	//Create a new prometheus registry
	http.Handle("/metrics", promhttp.Handler())

	prometheus.MustRegister(errored)

	// List the hosts to test and the ports
	ports := []string{"80", "443"}
	hosts := []string{"google.com", "192.168.64.6"}

	conFailedCounter := connect(hosts, ports)
	go func() {
		for {
			//send to prometheus metrics
			errored.WithLabelValues("checkConnection", "connection failed").Set(float64(conFailedCounter))
			time.Sleep(time.Second)
		}
	}()

	//Start http server
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func connect(hosts []string, ports []string) int {
	//Count number of failed connection
	conFailedCounter := 0

	//Loop through hosts
	for _, host := range hosts {
		for _, port := range ports {
			//Set timeout to 1 second
			timeout := time.Second
			//Try to connect to host:port
			conn, err := net.DialTimeout("tcp", net.JoinHostPort(host, port), timeout)
			//If error, print error and increment failed connection counter
			if err != nil {
				fmt.Println("Error:", err)
				conFailedCounter++
				fmt.Println("Connection failed:", conFailedCounter)
			}
			//If connection is open, close it and print host:port
			if conn != nil {
				defer conn.Close()
				fmt.Println("Opened", net.JoinHostPort(host, port))
			}
		}
	}
	return conFailedCounter
}
