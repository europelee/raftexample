package main

import (
	"crypto/sha1"
	"encoding/binary"
)

func newID(name, clusterName string) uint64 {
	var b []byte
	b = append(b, []byte(name)...)
	b = append(b, []byte(clusterName)...)
	hash := sha1.Sum(b)
	return binary.BigEndian.Uint64(hash[:8])
}
