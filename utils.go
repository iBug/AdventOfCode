package main

import "strconv"

type Coord struct {
	x, y int
}

func (c Coord) Equal(o Coord) bool {
	return c.x == o.x && c.y == o.y
}

func (c Coord) String() string {
	return strconv.Itoa(c.x) + "," + strconv.Itoa(c.y)
}

func Abs[T int](x T) T {
	if x < 0 {
		return -x
	}
	return x
}
