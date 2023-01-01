package main

import (
	"fmt"
	"strconv"
)

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

func Max[T int](x T, xs ...T) T {
	for _, y := range xs {
		if y > x {
			x = y
		}
	}
	return x
}

func FormatSize(sizeInt int64) string {
	size := float64(sizeInt)
	units := []string{"B", "KiB", "MiB", "GiB"}
	for _, unit := range units {
		if size < 1000 {
			return fmt.Sprintf("%.1f %s", size, unit)
		}
		size /= 1024
	}
	return fmt.Sprintf("%.1f %s", size, "TiB")
}
