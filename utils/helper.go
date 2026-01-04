package utils

import (
	"fmt"

	"github.com/Prayag2003/consistent-hashing/internal/ring"
)

func AssignRequests(ring *ring.HashRing, requests []string) map[string]int {
	result := make(map[string]int)

	for _, r := range requests {
		serverIP := ring.GetNearestServer(r)
		result[serverIP]++
	}
	return result
}

func PrintLoad(dist map[string]int) {
	for server, count := range dist {
		fmt.Printf("Server %s serving -> %d requests\n", server, count)
	}
}
