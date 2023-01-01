//go:build !linux

package main

import "runtime"

func GetCurrentMemoryUsage() int64 {
	var m runtime.MemStats
	runtime.ReadMemStats(&m)
	return int64(m.Alloc)
}
