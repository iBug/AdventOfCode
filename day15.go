package main

import (
	"bufio"
	"fmt"
	"io"
	"sort"
)

func init() {
	RegisterSolution("15-1", Solution15_1)
	// RegisterSolution("15-2", Solution15_2)
}

const (
	M15_END = iota
	M15_START
	M15_BEACON
)

func Solution15_1(r io.Reader) {
	type Flag struct {
		mode, x int
	}
	beacons := make(map[string]bool)
	flags := make([]Flag, 0, 100)

	targetY := 2000000

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		var sx, sy, bx, by int
		fmt.Sscanf(scanner.Text(), "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
		beacons[Coord{bx, by}.String()] = true
		dist := Abs(sx-bx) + Abs(sy-by)
		diff := dist - Abs(sy-targetY)
		if diff >= 0 {
			if diff == 0 && by == targetY && bx == sx {
				continue
			}
			start := sx - diff
			if by == targetY && bx == start {
				start++
			}
			flags = append(flags, Flag{M15_START, start})
			end := sx + diff
			if by == targetY && bx == end {
				end--
			}
			flags = append(flags, Flag{M15_END, end + 1})
		}
	}

	sort.Slice(flags, func(i, j int) bool {
		if flags[i].x < flags[j].x {
			return true
		}
		if flags[i].x == flags[j].x {
			return flags[i].mode < flags[j].mode
		}
		return false
	})

	total := 0
	coverage := 0
	startX := 0
	for _, f := range flags {
		switch f.mode {
		case M15_START:
			if coverage == 0 {
				startX = f.x
			}
			coverage++
		case M15_END:
			coverage--
			if coverage == 0 {
				total += f.x - startX
			}
		}
	}
	fmt.Println(total)
}
