package main

import (
	"fmt"
	"sort"
	"strconv"
)

type HashRing struct {
	replicas int

	// 	Ring (map):
	// {
	//   100  → "192.168.0.1"
	//   200  → "192.168.0.2"
	//   300  → "192.168.0.3"
	// }
	ring map[int32]string

	// keys (sorted): [100, 200, 300]
	keys []uint32
}

func NewHashRing(replicas int) *HashRing {
	return &HashRing{
		replicas: replicas,
		ring:     make(map[int32]string),
	}
}

// ------------------------
// - - - Add Server - - -
// ------------------------

// server IP : 192.168.0.1
// virtual Nodes: 192.168.0.1#vn0, 192.168.0.1#vn1 ..., 192.168.0.1#vn99

func (h *HashRing) AddServer(serverIP string) {
	fmt.Printf("\nAdding server %s with Virtal nodes %d", serverIP, h.replicas)
	for i := 0; i < h.replicas; i++ {
		// 192.168.0.1#vn1
		virtualNode := serverIP + "#vn" + strconv.Itoa(i)

		// hashKey(192.168.0.1#vn1) => 100
		hash := hashKey(virtualNode)

		// h.ring[100] = 192.168.0.1#vn1
		h.ring[int32(hash)] = serverIP

		// h.keys.append(100)
		h.keys = append(h.keys, hash)

		// sorting clockwise
		sort.Slice(h.keys, func(i, j int) bool {
			return h.keys[i] < h.keys[j]
		})
	}
}

// -----------------------------
// - - - Remove Server - - -
// -----------------------------

func (h *HashRing) RemoveServer(serverIP string) {
	fmt.Printf("Removing Server %s\n", serverIP)
	newKeys := h.keys[:0]

	for _, key := range h.keys {
		if h.ring[int32(key)] != serverIP {
			newKeys = append(newKeys, key)
		} else {
			delete(h.ring, int32(key))
		}
	}
	h.keys = newKeys
}

// -----------------------------
// - - - Lookup - - -
// -----------------------------

func (h *HashRing) GetNearestServer(requestIP string) string {
	if len(h.keys) == 0 {
		return " "
	}

	ipHash := hashKey(requestIP)

	idx := sort.Search(len(h.keys), func(i int) bool {
		return h.keys[i] >= ipHash
	})

	if idx == len(h.keys) {
		idx = 0
	}

	return h.ring[int32(h.keys[idx])]
}
