package main

import "strconv"

type Coord struct {
	x, y int
}

func (c Coord) Add(o Coord) Coord {
	return Coord{c.x + o.x, c.y + o.y}
}

func (c Coord) Equal(o Coord) bool {
	return c.x == o.x && c.y == o.y
}

func (c Coord) String() string {
	return strconv.Itoa(c.x) + "," + strconv.Itoa(c.y)
}

type Coord3 struct {
	x, y, z int
}

func (c Coord3) Add(o Coord3) Coord3 {
	return Coord3{c.x + o.x, c.y + o.y, c.z + o.z}
}

func (c Coord3) Equal(o Coord3) bool {
	return c.x == o.x && c.y == o.y && c.z == o.z
}

func (c Coord3) String() string {
	return strconv.Itoa(c.x) + "," + strconv.Itoa(c.y) + "," + strconv.Itoa(c.z)
}

func Abs[T int](x T) T {
	if x < 0 {
		return -x
	}
	return x
}
