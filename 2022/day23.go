package year

import (
	"bufio"
	"fmt"
	"io"
)

func init() {
	RegisterSolution("23-1", func(r io.Reader) { Solution23(r, 1) })
	RegisterSolution("23-2", func(r io.Reader) { Solution23(r, 2) })
}

func DirectionOk23(e map[Coord]int, c Coord) [4]bool {
	_, ok1 := e[Coord{c.x - 1, c.y - 1}]
	_, ok2 := e[Coord{c.x + 0, c.y - 1}]
	_, ok3 := e[Coord{c.x + 1, c.y - 1}]
	_, ok4 := e[Coord{c.x - 1, c.y + 0}]
	_, ok6 := e[Coord{c.x + 1, c.y + 0}]
	_, ok7 := e[Coord{c.x - 1, c.y + 1}]
	_, ok8 := e[Coord{c.x + 0, c.y + 1}]
	_, ok9 := e[Coord{c.x + 1, c.y + 1}]
	northOk := !ok1 && !ok2 && !ok3
	southOk := !ok7 && !ok8 && !ok9
	westOk := !ok1 && !ok4 && !ok7
	eastOk := !ok3 && !ok6 && !ok9
	if northOk && southOk && westOk && eastOk {
		return [4]bool{false, false, false, false}
	}
	return [4]bool{northOk, southOk, westOk, eastOk}
}

func MapBounds23(e map[Coord]int) (top, bottom, left, right int) {
	const LARGENUM = 999999
	top, bottom, left, right = LARGENUM, -LARGENUM, LARGENUM, -LARGENUM
	for c := range e {
		if c.y < top {
			top = c.y
		}
		if c.y > bottom {
			bottom = c.y
		}
		if c.x < left {
			left = c.x
		}
		if c.x > right {
			right = c.x
		}
	}
	return
}

func PrintMap23(e map[Coord]int) {
	top, bottom, left, right := MapBounds23(e)
	fmt.Printf("(%d,%d):\n", left, top)
	for y := top; y <= bottom; y++ {
		for x := left; x <= right; x++ {
			if _, ok := e[Coord{x, y}]; ok {
				fmt.Print("#")
			} else {
				fmt.Print(".")
			}
		}
		fmt.Println()
	}
}

var directions23 = []Coord{{0, -1}, {0, 1}, {-1, 0}, {1, 0}}

type Propose19 struct {
	Coord
	dir int
}

func Solution23(r io.Reader, mode int) {
	e := make(map[Coord]int)
	scanner := bufio.NewScanner(r)

	y := 0
	for scanner.Scan() {
		for x, s := range scanner.Text() {
			if s == '#' {
				e[Coord{x, y}] = 0
			}
		}
		y++
	}

	for round := 0; ; round++ {
		if mode == 1 && round == 10 {
			break
		}
		hasMoved := false

		proposal := make(map[Coord]int, len(e))
		proposed := make([]Propose19, 0, len(e))
		for c := range e {
			oks := DirectionOk23(e, c)
			i := 0
			for i = 0; i < 4; i++ {
				if oks[(i+round)%4] {
					newPos := c.Add(directions23[(i+round)%4])
					proposal[newPos]++
					proposed = append(proposed, Propose19{c, i})
					break
				}
			}
		}

		for _, p := range proposed {
			c, i := p.Coord, p.dir
			newPos := c.Add(directions23[(i+round)%4])
			if proposal[newPos] == 1 {
				hasMoved = true
				e[newPos] = 0
				delete(e, c)
			}
		}

		if mode == 2 && !hasMoved {
			fmt.Println(round + 1)
			break
		}
	}

	if mode == 1 {
		top, bottom, left, right := MapBounds23(e)
		result := (bottom-top+1)*(right-left+1) - len(e)
		fmt.Println(result)
	}
}
