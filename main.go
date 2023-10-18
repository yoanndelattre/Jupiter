package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	// Define the list of backend servers (replace with your own backend server URLs).
	backendServers := []string{
		"https://www.google.com",
	}

	// Number of requests to send per client.
	numRequestsPerClient := 10

	// Number of concurrent clients.
	numClients := 5

	var wg sync.WaitGroup

	for i := 0; i < numClients; i++ {
		wg.Add(1)
		go func(clientID int) {
			defer wg.Done()

			for j := 0; j < numRequestsPerClient; j++ {
				backendURL := backendServers[j%len(backendServers)]
				resp, err := http.Get(backendURL)
				if err != nil {
					fmt.Printf("Client %d: Error making request to %s: %v\n", clientID, backendURL, err)
					continue
				}
				defer resp.Body.Close()
				fmt.Printf("Client %d: Request to %s returned status code %d\n", clientID, backendURL, resp.StatusCode)
			}
		}(i)
	}

	wg.Wait()
}
