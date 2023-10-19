package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
	"net/url"
	"time"

	"golang.org/x/net/proxy"
)

func main() {
	targetURL := flag.String("target", "http://localhost:8080", "Target URL for the load balancer")
	proxyAddr := flag.String("proxy", "", "SOCKS5 proxy address (e.g., 127.0.0.1:1080)")
	requestsPerSecond := flag.Int("rps", 10, "Requests per second")
	duration := flag.Duration("duration", 10*time.Second, "Test duration")
	flag.Parse()

	fmt.Printf("Testing Load Balancer with %d requests per second for %v\n", *requestsPerSecond, *duration)

	var dialer proxy.Dialer
	if *proxyAddr != "" {
		// Initialize SOCKS5 proxy (if provided).
		proxyURL, err := url.Parse("socks5://" + *proxyAddr)
		if err != nil {
			log.Fatalf("Error parsing proxy address: %v", err)
		}
		dialer, err = proxy.FromURL(proxyURL, proxy.Direct)
		if err != nil {
			log.Fatalf("Error creating proxy dialer: %v", err)
		}
	}

	// Create an HTTP client with the proxy dialer if it's set.
	httpClient := &http.Client{
		Transport: &http.Transport{
			Dial: dialer.Dial,
		},
	}

	// Start the test.
	startTime := time.Now()
	requestCount := 0
	for {
		select {
		case <-time.After(time.Second / time.Duration(*requestsPerSecond)):
			go func() {
				sendRequest(httpClient, *targetURL)
				requestCount++
			}()
		case <-time.After(*duration):
			elapsedTime := time.Since(startTime)
			fmt.Printf("Test completed. Sent %d requests in %v\n", requestCount, elapsedTime)
			return
		}
	}
}

func sendRequest(client *http.Client, targetURL string) {
	resp, err := client.Get(targetURL)
	if err != nil {
		log.Printf("Error sending request to %s: %v\n", targetURL, err)
		return
	}
	defer resp.Body.Close()
	fmt.Printf("Request to %s completed with status: %s\n", targetURL, resp.Status)
}
