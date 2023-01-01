package main

import (
	"syscall"
)

func GetCurrentMemoryUsage() int64 {
	var rusage syscall.Rusage
	if err := syscall.Getrusage(syscall.RUSAGE_SELF, &rusage); err != nil {
		return 0
	}
	return rusage.Maxrss * 1024
}
