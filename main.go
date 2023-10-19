package main

import (
	"log"
	"net/http"
	"net/url"
	"os"

	"golang.org/x/net/proxy"
)

func main() {
	targetURL := os.Getenv("TARGET_URL")   // Target URL for the load balancer
	proxyAddr := os.Getenv("SOCKS5_PROXY") // e.g., 127.0.0.1:1080

	var dialer proxy.Dialer
	if proxyAddr != "" {
		// Initialize SOCKS5 proxy dialer if proxy is set.
		proxyURL, err := url.Parse("socks5://" + proxyAddr)
		if err != nil {
			log.Fatalf("Error parsing proxy address: %v", err)
		}
		dialer, err = proxy.FromURL(proxyURL, proxy.Direct)
		if err != nil {
			log.Fatalf("Error creating proxy dialer: %v", err)
		}
	}

	// Create an HTTP client with or without the proxy dialer.
	httpClient := &http.Client{}
	if dialer != nil {
		httpClient.Transport = &http.Transport{
			Dial: dialer.Dial,
		}
		log.Printf("The SOCKS5 proxy is set\n")
	} else {
		log.Printf("The SOCKS5 proxy is not set\n")
	}

	// Start the test.
	for {
		sendRequest(httpClient, targetURL)
	}
}

func sendRequest(client *http.Client, targetURL string) {
	resp, err := client.Get(targetURL)
	if err != nil {
		log.Printf("Error sending request to %s: %v\n", targetURL, err)
		return
	}
	defer resp.Body.Close()
	log.Printf("Request to %s completed with status: %s\n", targetURL, resp.Status)
}
