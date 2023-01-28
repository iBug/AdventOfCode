package year

import (
	"bufio"
	"fmt"
	"io"
)

func init() {
	RegisterSolution("17-1", func(r io.Reader, w io.Writer) { Solution17(r, w, 1) })
	RegisterSolution("17-2", func(r io.Reader, w io.Writer) { Solution17(r, w, 2) })
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

func Solution17(r io.Reader, w io.Writer, mode int) {
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

	// for part2
	tops := make([]int, 0)
	sampleN := 20
	sampleInterval := 1000

	for n := 0; ; n++ {
		if mode == 1 && n >= 2022 {
			fmt.Fprintln(w, top+1)
			break
		} else if mode == 2 {
			if n%sampleInterval == 0 {
				tops = append(tops, top+1)
				if len(tops) >= sampleN {
					break
				}
			}
		}

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

	if mode == 2 {
		interval := 0
		ok := true
		for interval = 1; interval < len(tops)-1; interval++ {
			diff := tops[1+interval] - tops[1]
			ok = true
			for i := 2; i < len(tops)-interval; i++ {
				if tops[i]+diff != tops[i+interval] {
					ok = false
					break
				}
			}
			if ok {
				target := int(1000000000000 / int64(sampleInterval))
				rem := target % interval
				if rem == 0 {
					rem = interval
				}
				fmt.Fprintln(w, tops[rem]+(target/interval)*diff)
				break
			}
		}
		if !ok {
			fmt.Printf("No pattern found, perhaps the sample interval %d is too small?\n", sampleInterval)
		}
	}
}
