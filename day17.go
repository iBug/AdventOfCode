package main

import (
	"bufio"
	"fmt"
	"io"
)

func init() {
	RegisterSolution("17-1", Solution17_1)
	// RegisterSolution("17-2", Solution17_2)
}

type Rock17 struct {
	w, h int
	s    []Coord
}

const (
	J17_LEFT = iota
	J17_RIGHT
)

var rocks17 = []Rock17{
	{w: 4, h: 1, s: []Coord{{0, 0}, {1, 0}, {2, 0}, {3, 0}}},
	{w: 3, h: 3, s: []Coord{{1, 0}, {0, 1}, {1, 1}, {2, 1}, {1, 2}}},
	{w: 3, h: 3, s: []Coord{{2, 0}, {2, 1}, {0, 2}, {1, 2}, {2, 2}}},
	{w: 1, h: 4, s: []Coord{{0, 0}, {0, 1}, {0, 2}, {0, 3}}},
	{w: 2, h: 2, s: []Coord{{0, 0}, {1, 0}, {0, 1}, {1, 1}}},
}

func Solution17_1(r io.Reader) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	jets := make([]int, len(scanner.Text()))
	for i, c := range scanner.Text() {
		if c == '<' {
			jets[i] = J17_LEFT
		} else {
			jets[i] = J17_RIGHT
		}
	}

	m := make([][7]int, 3)
	top := -1
	tick := 0
	for n := 0; n < 2022; n++ {
		rock := rocks17[n%len(rocks17)]
		r := Coord{2, top + rock.h + 3}
		for len(m) <= r.y {
			m = append(m, [7]int{})
		}

		for {
			if jets[tick] == J17_LEFT {
				if r.x > 0 {
					r.x--
					for _, s := range rock.s {
						if m[r.y-s.y][r.x+s.x] > 0 {
							r.x++
							break
						}
					}
				}
			} else {
				if r.x < 7-rock.w {
					r.x++
					for _, s := range rock.s {
						if m[r.y-s.y][r.x+s.x] > 0 {
							r.x--
							break
						}
					}
				}
			}
			tick++
			if tick >= len(jets) {
				tick = 0
			}

			ok := true
			if r.y < rock.h {
				ok = false
			} else {
				r.y--
				for _, s := range rock.s {
					if m[r.y-s.y][r.x+s.x] > 0 {
						r.y++
						ok = false
						break
					}
				}
			}
			if !ok {
				for _, s := range rock.s {
					m[r.y-s.y][r.x+s.x] = n + 1
				}
				if top < r.y {
					top = r.y
				}
				break
			}
		}
	}
	fmt.Println(top + 1)
}
