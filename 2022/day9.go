package year

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func init() {
	RegisterSolution("9-1", Solution9_1)
	RegisterSolution("9-1a", func(r io.Reader) { Solution9(r, 2) })
	RegisterSolution("9-2", func(r io.Reader) { Solution9(r, 10) })
}

func Solution9_1(r io.Reader) {
	scanner := bufio.NewScanner(r)
	hx, hy, tx, ty := 0, 0, 0, 0
	visited := make(map[string]int)
	visited["0,0"] = 1
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		dir := fields[0]
		dist, _ := strconv.Atoi(fields[1])

		for i := 0; i < dist; i++ {
			switch dir {
			case "U":
				hy++
				if hy-ty > 1 {
					tx = hx
					ty = hy - 1
				}
			case "D":
				hy--
				if ty-hy > 1 {
					tx = hx
					ty = hy + 1
				}
			case "L":
				hx--
				if tx-hx > 1 {
					tx = hx + 1
					ty = hy
				}
			case "R":
				hx++
				if hx-tx > 1 {
					tx = hx - 1
					ty = hy
				}
			}
			visited[fmt.Sprintf("%d,%d", tx, ty)]++
		}
	}
	fmt.Println(len(visited))
}

func Move9(x1, y1, x2, y2 int) (int, int) {
	if x1-x2 < -1 {
		x2 = x1 + 1
		if y2 > y1 {
			y2--
		} else if y2 < y1 {
			y2++
		}
	}
	if x1-x2 > 1 {
		x2 = x1 - 1
		if y2 > y1 {
			y2--
		} else if y2 < y1 {
			y2++
		}
	}
	if y1-y2 < -1 {
		y2 = y1 + 1
		if x2 > x1 {
			x2--
		} else if x2 < x1 {
			x2++
		}
	}
	if y1-y2 > 1 {
		y2 = y1 - 1
		if x2 > x1 {
			x2--
		} else if x2 < x1 {
			x2++
		}
	}
	return x2, y2
}

func Solution9(r io.Reader, size int) {
	scanner := bufio.NewScanner(r)
	x, y := make([]int, size), make([]int, size)
	visited := make(map[string]int)
	visited["0,0"] = 1
	for scanner.Scan() {
		fields := strings.Fields(scanner.Text())
		dir := fields[0]
		dist, _ := strconv.Atoi(fields[1])

		for i := 0; i < dist; i++ {
			switch dir {
			case "U":
				y[0]++
			case "D":
				y[0]--
			case "L":
				x[0]--
			case "R":
				x[0]++
			}
			for j := 1; j < len(y); j++ {
				nx, ny := Move9(x[j-1], y[j-1], x[j], y[j])
				if nx != x[j] || ny != y[j] {
					x[j] = nx
					y[j] = ny
				} else {
					break
				}
			}
			visited[fmt.Sprintf("%d,%d", x[size-1], y[size-1])]++
		}
	}
	fmt.Println(len(visited))
}
