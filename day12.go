package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"text/tabwriter"
)

func init() {
	RegisterSolution("12-1", func(r io.Reader) { Solution12(r, 1) })
	RegisterSolution("12-2", func(r io.Reader) { Solution12(r, 2) })
}

type Coord struct {
	x, y int
}

func (c Coord) Equal(o Coord) bool {
	return c.x == o.x && c.y == o.y
}

func Solution12(r io.Reader, mode int) {
	h := make([][]int, 0)
	var s, e Coord
	d := make([][]int, 0)
	q := make([]Coord, 0)

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		str := scanner.Text()
		row := make([]int, 0, len(str))
		for x, c := range str {
			if c == 'S' {
				s.y = len(h)
				s.x = x
				c = 'a'
			} else if c == 'E' {
				e.y = len(h)
				e.x = x
				c = 'z'
			}
			row = append(row, int(c-'a'))
		}
		h = append(h, row)
		d = append(d, make([]int, len(str)))
	}

	q = append(q, e)
	for len(q) > 0 {
		c := q[0]
		dist := d[c.y][c.x]
		q = q[1:]
		if mode == 1 && c.Equal(s) {
			fmt.Println(dist)
			break
		} else if mode == 2 && h[c.y][c.x] == 0 {
			fmt.Println(dist)
			break
		}

		for _, n := range []Coord{
			{c.x, c.y - 1},
			{c.x, c.y + 1},
			{c.x - 1, c.y},
			{c.x + 1, c.y},
		} {
			if n.x >= 0 && n.x < len(h[0]) && n.y >= 0 && n.y < len(h) && d[n.y][n.x] == 0 {
				if h[n.y][n.x]-h[c.y][c.x] >= -1 {
					d[n.y][n.x] = dist + 1
					q = append(q, n)
				}
			}
		}
	}

	if false {
		// DEBUG: prints the distance map
		writer := tabwriter.NewWriter(os.Stdout, 0, 8, 1, ' ', tabwriter.AlignRight)
		for _, row := range d {
			for _, v := range row {
				fmt.Fprintf(writer, "%d\t", v)
			}
			fmt.Fprintln(writer)
		}
		writer.Flush()
	}
}
