package year

import (
	"bufio"
	"fmt"
	"io"
)

func init() {
	RegisterSolution("18-1", Solution18_1)
	RegisterSolution("18-2", Solution18_2)
}

var Adjacent18 = []Coord3{{1, 0, 0}, {-1, 0, 0}, {0, 1, 0}, {0, -1, 0}, {0, 0, 1}, {0, 0, -1}}

func Solution18_1(r io.Reader) {
	scanner := bufio.NewScanner(r)
	m := make(map[string]bool)
	total := 0
	for scanner.Scan() {
		c := Coord3{}
		fmt.Sscanf(scanner.Text(), "%d,%d,%d", &c.x, &c.y, &c.z)
		m[c.String()] = true
		total += 6
		for _, d := range Adjacent18 {
			if m[c.Add(d).String()] {
				total -= 2
			}
		}
	}
	fmt.Println(total)
}

func InCubeRange18(c, min, max Coord3) bool {
	return c.x >= min.x && c.x <= max.x &&
		c.y >= min.y && c.y <= max.y &&
		c.z >= min.z && c.z <= max.z
}

func Solution18_2(r io.Reader) {
	scanner := bufio.NewScanner(r)
	m := make(map[string]int, 8192)
	min, max := Coord3{999, 999, 999}, Coord3{-999, -999, -999}
	for scanner.Scan() {
		c := Coord3{}
		fmt.Sscanf(scanner.Text(), "%d,%d,%d", &c.x, &c.y, &c.z)
		m[c.String()] = 1
		if c.x < min.x {
			min.x = c.x
		}
		if c.y < min.y {
			min.y = c.y
		}
		if c.z < min.z {
			min.z = c.z
		}
		if c.x > max.x {
			max.x = c.x
		}
		if c.y > max.y {
			max.y = c.y
		}
		if c.z > max.z {
			max.z = c.z
		}
	}
	min = min.Add(Coord3{-1, -1, -1})
	max = max.Add(Coord3{1, 1, 1})

	stack := make([]Coord3, 0, 128)
	total := 0
	stack = append(stack, min)
	for len(stack) > 0 {
		c := stack[len(stack)-1]
		stack = stack[:len(stack)-1]
		if m[c.String()] != 0 {
			continue
		}
		m[c.String()] = 2
		total += 6
		for _, d := range Adjacent18 {
			adj := c.Add(d)
			if !InCubeRange18(adj, min, max) {
				continue
			}
			t := m[adj.String()]
			if t == 0 {
				stack = append(stack, adj)
			} else if t == 2 {
				total -= 2
			}
		}
	}
	xSize := max.x - min.x + 1
	ySize := max.y - min.y + 1
	zSize := max.z - min.z + 1
	total -= 2 * (xSize*ySize + xSize*zSize + ySize*zSize)
	fmt.Println(total)
}
