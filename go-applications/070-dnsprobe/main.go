package main

import (
	"encoding/json"
	"fmt"
	"net"
	"time"
)

type DNSResult struct {
	Domain  string
	Latency time.Duration
	Error   error
}

func measureDNSTime(domain string, results chan<- DNSResult) {
	start := time.Now()
	_, err := net.LookupHost(domain)
	latency := time.Since(start)

	results <- DNSResult{
		Domain:  domain,
		Latency: latency,
		Error:   err,
	}
}

func main() {
	// example JSON data containing a list of domains
	jsonData := `{"domains": ["google.com", "facebook.com", "twitter.com"]}`

	// parse the JSON data
	var data struct {
		Domains []string `json:"domains"`
	}
	err := json.Unmarshal([]byte(jsonData), &data)
	if err != nil {
		fmt.Println("JSON parsing error:", err)
		return
	}

	// create a channel for the results
	results := make(chan DNSResult, len(data.Domains))

	// measure DNS latency for each domain in a separate goroutine
	for _, domain := range data.Domains {
		go measureDNSTime(domain, results)
	}

	// wait for all goroutines to finish and collect the results
	for i := 0; i < len(data.Domains); i++ {
		result := <-results
		if result.Error != nil {
			fmt.Printf("DNS lookup error for %s: %v\n", result.Domain, result.Error)
		} else {
			fmt.Printf("DNS lookup took %v for %s\n", result.Latency, result.Domain)
		}
	}
}
