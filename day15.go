package main

import (
	"bufio"
	"fmt"
	"io"
	"sort"
)

func init() {
	RegisterSolution("15-1", Solution15_1)
	RegisterSolution("15-2", Solution15_2)
}

const (
	M15_END = iota
	M15_START
	M15_BEACON
)

type Flag15 struct {
	mode, x int
}

func Solution15_1(r io.Reader) {
	beacons := make(map[string]bool)
	flags := make([]Flag15, 0, 100)

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
			flags = append(flags, Flag15{M15_START, start})
			end := sx + diff
			if by == targetY && bx == end {
				end--
			}
			flags = append(flags, Flag15{M15_END, end + 1})
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

type Line15 struct {
	// x + by = c, delimited by x1 and x2
	// assume x1 < x2 and Abs(b) == 1
	x1, x2, b, c int
}

func LineIntersect15(l1, l2 Line15) (Coord, bool) {
	det := l2.b - l1.b
	if det == 0 || (l2.c-l1.c)%2 == 1 {
		return Coord{0, 0}, false
	}
	x := (l2.b*l1.c - l1.b*l2.c) / det
	if x < l1.x1 || x < l2.x1 || x >= l1.x2 || x >= l2.x2 {
		return Coord{0, 0}, false
	}
	y := (l2.c - l1.c) / det
	return Coord{x, y}, true
}

func Solution15_2(r io.Reader) {
	type Sensor struct {
		Coord
		r int // range or radius
	}

	sensors := make([]Sensor, 0, 50)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		var sx, sy, bx, by int
		fmt.Sscanf(scanner.Text(), "Sensor at x=%d, y=%d: closest beacon is at x=%d, y=%d", &sx, &sy, &bx, &by)
		r := Abs(sx-bx) + Abs(sy-by)
		sensors = append(sensors, Sensor{Coord{sx, sy}, r})
	}

	lines := make([]Line15, 0, 200)
	intersects := make(map[string]int)
	for _, s := range sensors {
		lines = append(lines, Line15{s.x - s.r - 1, s.x + 1, 1, s.x + s.y - s.r - 1})
		lines = append(lines, Line15{s.x - s.r - 1, s.x + 1, -1, s.x - s.y - s.r - 1})
		lines = append(lines, Line15{s.x, s.x + s.r + 2, 1, s.x + s.y + s.r + 1})
		lines = append(lines, Line15{s.x, s.x + s.r + 2, -1, s.x - s.y + s.r + 1})
	}
	for i := 0; i < len(lines); i++ {
		for j := i + 1; j < len(lines); j++ {
			if p, ok := LineIntersect15(lines[i], lines[j]); ok {
				intersects[p.String()]++
			}
		}
	}

	var x, y int
	maxK, maxV := "", 0
	for k, v := range intersects {
		fmt.Sscanf(k, "%d,%d", &x, &y)
		for _, s := range sensors {
			if Abs(x-s.x)+Abs(y-s.y) <= s.r {
				v = 0
				break
			}
		}
		if v > maxV {
			maxK = k
			maxV = v
		}
	}
	fmt.Sscanf(maxK, "%d,%d", &x, &y)
	fmt.Println(maxK, 4000000*int64(x)+int64(y))
}
