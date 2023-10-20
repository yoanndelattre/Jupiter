package main

import (
	"crypto/tls"
	"log"
	"net/http"
	"net/url"
	"os"
	"time"

	"golang.org/x/net/proxy"
)

func main() {
	targetURL := os.Getenv("TARGET_URL")   // Target URL for the load balancer
	proxyAddr := os.Getenv("SOCKS5_PROXY") // e.g., 127.0.0.1:1080

	if targetURL == "" {
		log.Fatalf("TARGET_URL is empty")
	}

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

	// Wait for the proxy connection to become available with a retry mechanism.
	for dialer != nil {
		err := testProxyConnection(dialer)
		if err == nil {
			break
		}
		log.Printf("Proxy connection failed: %v. Retrying in 5 seconds...", err)
		time.Sleep(5 * time.Second)
	}

	// Create an HTTP client with or without the proxy dialer.
	httpClient := &http.Client{
		Timeout: time.Second * 5,
	}
	if dialer != nil {
		httpClient.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
			Dial:            dialer.Dial,
		}
		log.Printf("The SOCKS5 proxy is set\n")
	} else {
		httpClient.Transport = &http.Transport{
			TLSClientConfig: &tls.Config{InsecureSkipVerify: true},
		}
		log.Printf("The SOCKS5 proxy is not set\n")
	}

	// Start the test.
	for {
		sendRequest(httpClient, targetURL)
	}
}

func testProxyConnection(dialer proxy.Dialer) error {
	conn, err := dialer.Dial("tcp", "bitcoin.org:80")
	if err != nil {
		return err
	}
	conn.Close()
	return nil
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
