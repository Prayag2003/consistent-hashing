package main

import "hash/crc32"

func hashKey(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}
