package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func init() {
	RegisterSolution("9-1", Solution9_1)
	// RegisterSolution("9-2", Solution9_2)
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
