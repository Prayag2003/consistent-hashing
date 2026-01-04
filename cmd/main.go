package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/Prayag2003/consistent-hashing/internal/ring"
	"github.com/Prayag2003/consistent-hashing/utils"
)

func main() {
	replicas := 100
	if len(os.Args) > 1 {
		if r, err := strconv.Atoi(os.Args[1]); err == nil {
			replicas = r
		}
	}
	fmt.Println("Virtual nodes per server: ", replicas)

	ring := ring.NewHashRing(replicas)
	servers := []string{
		"10.0.0.1",
		"10.0.0.2",
		"10.0.0.3",
	}

	for _, s := range servers {
		ring.AddServer(s)
	}

	var requests []string
	for i := range 10000 {
		requests = append(requests, "192.168.1."+strconv.Itoa(i))
	}

	fmt.Println("\n\nLoad Distribution before adding server:")
	utils.PrintLoad(utils.AssignRequests(ring, requests))

	ring.AddServer("10.0.0.4")
	fmt.Println("\n\nLoad Distribution after adding server:")
	utils.PrintLoad(utils.AssignRequests(ring, requests))
}
