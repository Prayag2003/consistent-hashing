package main

import "fmt"

func assignRequests(ring *HashRing, requests []string) map[string]int {
	result := make(map[string]int)

	for _, r := range requests {
		serverIP := ring.GetNearestServer(r)
		result[serverIP]++
	}
	return result
}

func printLoad(dist map[string]int) {
	for server, count := range dist {
		fmt.Printf("Server %s serving -> %d requests\n", server, count)
	}
}
