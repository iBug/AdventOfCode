package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func init() {
	RegisterSolution("14-1", func(r io.Reader) { Solution14(r, 1) })
	RegisterSolution("14-2", func(r io.Reader) { Solution14(r, 2) })
}

const (
	E14_EMPTY = iota
	E14_ROCK
	E14_SAND
)

func Solution14(r io.Reader, mode int) {
	m := make(map[string]int)
	maxY := 0

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		parts := strings.Split(scanner.Text(), " -> ")
		var c, n Coord
		fmt.Sscanf(parts[0], "%d,%d", &c.x, &c.y)

		if c.y >= maxY {
			maxY = c.y + 1
		}

		for i := 1; i < len(parts); i++ {
			fmt.Sscanf(parts[i], "%d,%d", &n.x, &n.y)
			if n.y >= maxY {
				maxY = n.y + 1
			}

			if n.x == c.x {
				dir := -1
				if n.y > c.y {
					dir = 1
				}
				for y := c.y; y != n.y; y += dir {
					m[Coord{c.x, y}.String()] = E14_ROCK
				}
			} else if n.y == c.y {
				dir := -1
				if n.x > c.x {
					dir = 1
				}
				for x := c.x; x != n.x; x += dir {
					m[Coord{x, c.y}.String()] = E14_ROCK
				}
			} else {
				panic("invalid input")
			}
			c = n
		}
		m[parts[len(parts)-1]] = E14_ROCK
	}

	count := 0
outer:
	for {
		s := Coord{500, 0}
		if mode == 2 && m[s.String()] == E14_SAND {
			break
		}
	inner:
		for {
			if s.y >= maxY {
				if mode == 1 {
					break outer
				} else if mode == 2 {
					m[s.String()] = E14_SAND
					break inner
				}
			}
			for _, n := range []Coord{
				{s.x, s.y + 1},
				{s.x - 1, s.y + 1},
				{s.x + 1, s.y + 1},
			} {
				if _, ok := m[n.String()]; !ok {
					s = n
					continue inner
				}
			}
			m[s.String()] = E14_SAND
			break
		}
		count++
	}
	fmt.Println(count)
}
