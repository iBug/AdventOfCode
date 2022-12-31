package main

import (
	"bufio"
	"fmt"
	"io"
)

func init() {
	RegisterSolution("18-1", Solution18_1)
	// RegisterSolution("18-2", Solution18_2)
}

func Solution18_1(r io.Reader) {
	scanner := bufio.NewScanner(r)
	m := make(map[string]bool)
	total := 0
	for scanner.Scan() {
		c := Coord3{}
		fmt.Sscanf(scanner.Text(), "%d,%d,%d", &c.x, &c.y, &c.z)
		m[c.String()] = true
		total += 6
		for _, d := range []Coord3{{1, 0, 0}, {-1, 0, 0}, {0, 1, 0}, {0, -1, 0}, {0, 0, 1}, {0, 0, -1}} {
			if m[c.Add(d).String()] {
				total -= 2
			}
		}
	}
	fmt.Println(total)
}
