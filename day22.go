package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

func init() {
	RegisterSolution("22-1", Solution22_1)
	// RegisterSolution("22-2", Solution22_2)
}

const (
	M22_VOID = iota
	M22_EMPTY
	M22_WALL
)

var (
	D22_UP    = Coord{0, -1}
	D22_DOWN  = Coord{0, 1}
	D22_LEFT  = Coord{-1, 0}
	D22_RIGHT = Coord{1, 0}

	D22 = []Coord{D22_RIGHT, D22_DOWN, D22_LEFT, D22_UP}
)

func ReadData22(r io.Reader) (m [][]int, moves []int, turns []byte) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		s := scanner.Text()
		if len(strings.TrimSpace(s)) == 0 {
			break
		}
		row := make([]int, len(s))
		for i := 0; i < len(s); i++ {
			switch s[i] {
			case '#':
				row[i] = M22_WALL
			case '.':
				row[i] = M22_EMPTY
			case ' ':
				row[i] = M22_VOID
			default:
				panic("unknown char")
			}
		}
		m = append(m, row)
	}

	if scanner.Scan() {
		f := strings.NewReader(scanner.Text())
		for {
			var n int
			fmt.Fscanf(f, "%d", &n)
			moves = append(moves, n)
			c, err := f.ReadByte()
			if errors.Is(err, io.EOF) {
				break
			} else if err != nil {
				panic(err)
			}
			turns = append(turns, c)
		}
	}
	turns = append(turns, 0)
	return
}

func Move22_1(m [][]int, pos Coord, dir int) (Coord, bool) {
	x, y := pos.x, pos.y
	switch dir {
	case 0:
		if x >= len(m[y])-1 {
			for x = 0; x < len(m[y]) && m[y][x] == M22_VOID; x++ {
			}
		} else {
			x++
		}
	case 1:
		if y >= len(m)-1 || len(m[y+1]) <= x || m[y+1][x] == M22_VOID {
			// wrap around
			for y = 0; y < len(m) && m[y][x] == M22_VOID; y++ {
			}
		} else {
			y++
		}
	case 2:
		if x == 0 || m[y][x-1] == M22_VOID {
			x = len(m[y]) - 1
		} else {
			x--
		}
	case 3:
		if y == 0 || len(m[y-1]) <= x || m[y-1][x] == M22_VOID {
			for y = len(m) - 1; len(m[y]) <= x || m[y][x] == M22_VOID; y-- {
			}
		} else {
			y--
		}
	}
	if m[y][x] == M22_WALL {
		return pos, false
	}
	return Coord{x, y}, true
}

func Solution22_1(r io.Reader) {
	m, moves, turns := ReadData22(r)
	pos := Coord{0, 0}
	dir := 0
	for i, c := range m[0] {
		if c == M22_EMPTY {
			pos.x = i
			break
		}
	}

	for round := 0; round < len(moves); round++ {
		for i := 0; i < moves[round]; i++ {
			newPos, ok := Move22_1(m, pos, dir)
			if !ok {
				break
			}
			pos = newPos
		}
		switch turns[round] {
		case 'L':
			dir = (dir + 3) % 4
		case 'R':
			dir = (dir + 1) % 4
		}
	}

	result := 1000*(pos.y+1) + 4*(pos.x+1) + dir
	fmt.Println(result)
}
