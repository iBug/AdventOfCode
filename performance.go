package main

import (
	"fmt"
	"io"
	"time"
)

type PerformanceState struct {
	Time   time.Time
	Memory int64
}

type PerformanceInfo struct {
	Time   time.Duration
	Memory int64
}

func GetPerformanceState() PerformanceState {
	return PerformanceState{
		Time:   time.Now(),
		Memory: GetCurrentMemoryUsage(),
	}
}

func DiffPerformanceState(s PerformanceState) PerformanceInfo {
	return PerformanceInfo{
		Time:   time.Since(s.Time),
		Memory: GetCurrentMemoryUsage(),
	}
}

func PrintPerformanceInfo(w io.Writer, info PerformanceInfo) {
	fmt.Fprintf(w, "Time: %s, Memory: %s\n", info.Time.Truncate(time.Microsecond), FormatSize(info.Memory))
}
