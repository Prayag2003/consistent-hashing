package hash

import "hash/crc32"

func HashKey(key string) uint32 {
	return crc32.ChecksumIEEE([]byte(key))
}
