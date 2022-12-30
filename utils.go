package main

import "fmt"

type Coord struct {
	x, y int
}

func (c Coord) Equal(o Coord) bool {
	return c.x == o.x && c.y == o.y
}

func (c Coord) String() string {
	return fmt.Sprintf("%d,%d", c.x, c.y)
}
