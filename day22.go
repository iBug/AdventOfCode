package main

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"strings"
)

func init() {
	RegisterSolution("22-1", func(r io.Reader) { Solution22(r, 1) })
	RegisterSolution("22-2", func(r io.Reader) { Solution22(r, 2) })
	RegisterSolution("22-2e", func(r io.Reader) { Solution22(r, 3) })
}

const (
	M22_VOID = iota
	M22_EMPTY
	M22_WALL
)

const (
	D22_RIGHT = 0
	D22_DOWN  = 1
	D22_LEFT  = 2
	D22_UP    = 3
)

func ReadData22(r io.Reader) (m [][]int, pos Coord, moves []int, turns []byte) {
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

	for i, c := range m[0] {
		if c == M22_EMPTY {
			pos.x = i
			break
		}
	}
	return
}

func Move22_1(m [][]int, pos Coord, dir int) (Coord, int, bool) {
	x, y := pos.x, pos.y
	switch dir {
	case D22_RIGHT:
		if x >= len(m[y])-1 {
			for x = 0; x < len(m[y]) && m[y][x] == M22_VOID; x++ {
			}
		} else {
			x++
		}
	case D22_DOWN:
		if y >= len(m)-1 || len(m[y+1]) <= x || m[y+1][x] == M22_VOID {
			// wrap around
			for y = 0; y < len(m) && m[y][x] == M22_VOID; y++ {
			}
		} else {
			y++
		}
	case D22_LEFT:
		if x == 0 || m[y][x-1] == M22_VOID {
			x = len(m[y]) - 1
		} else {
			x--
		}
	case D22_UP:
		if y == 0 || len(m[y-1]) <= x || m[y-1][x] == M22_VOID {
			for y = len(m) - 1; len(m[y]) <= x || m[y][x] == M22_VOID; y-- {
			}
		} else {
			y--
		}
	}
	if m[y][x] == M22_WALL {
		return pos, dir, false
	}
	return Coord{x, y}, dir, true
}

func Move22_2(m [][]int, pos Coord, dir int) (Coord, int, bool) {
	const SIZE = 50
	x, y, newDir := pos.x, pos.y, dir
	switch dir {
	case D22_RIGHT:
		if x >= len(m[y])-1 {
			if y < SIZE {
				newDir = D22_LEFT
				y = 3*SIZE - y - 1
				x = 2*SIZE - 1
			} else if y < 2*SIZE {
				newDir = D22_UP
				x = SIZE + y
				y = SIZE - 1
			} else if y < 3*SIZE {
				newDir = D22_LEFT
				y = 3*SIZE - y - 1
				x = 3*SIZE - 1
			} else {
				newDir = D22_UP
				x = y - 2*SIZE
				y = 3*SIZE - 1
			}
		} else {
			x++
		}
	case D22_DOWN:
		if y >= len(m)-1 || len(m[y+1]) <= x {
			if x < SIZE {
				newDir = D22_DOWN
				x = x + 2*SIZE
				y = 0
			} else if x < 2*SIZE {
				newDir = D22_LEFT
				y = x + 2*SIZE
				x = SIZE - 1
			} else {
				newDir = D22_LEFT
				y = x - SIZE
				x = 2*SIZE - 1
			}
		} else {
			y++
		}
	case D22_LEFT:
		if x == 0 || m[y][x-1] == M22_VOID {
			if y < SIZE {
				newDir = D22_RIGHT
				y = 3*SIZE - y - 1
				x = 0
			} else if y < 2*SIZE {
				newDir = D22_DOWN
				x = y - SIZE
				y = 2 * SIZE
			} else if y < 3*SIZE {
				newDir = D22_RIGHT
				y = 3*SIZE - y - 1
				x = SIZE
			} else {
				newDir = D22_DOWN
				x = y - 2*SIZE
				y = 0
			}
		} else {
			x--
		}
	case D22_UP:
		if y == 0 || m[y-1][x] == M22_VOID {
			if x < SIZE {
				newDir = D22_RIGHT
				y = x + SIZE
				x = SIZE
			} else if x < 2*SIZE {
				newDir = D22_RIGHT
				y = x + 2*SIZE
				x = 0
			} else {
				newDir = D22_UP
				y = 4*SIZE - 1
				x = x - 2*SIZE
			}
		} else {
			y--
		}
	}
	if m[y][x] == M22_WALL {
		return pos, dir, false
	}
	return Coord{x, y}, newDir, true
}

func Move22_2e(m [][]int, pos Coord, dir int) (Coord, int, bool) {
	const SIZE = 4
	x, y, newDir := pos.x, pos.y, dir
	switch dir {
	case D22_RIGHT:
		if x >= len(m[y])-1 {
			if y < SIZE {
				newDir = D22_LEFT
				y = 3*SIZE - y - 1
				x = 4*SIZE - 1
			} else if y < 2*SIZE {
				newDir = D22_DOWN
				x = 5*SIZE - y - 1
				y = 2 * SIZE
			} else {
				newDir = D22_LEFT
				y = 3*SIZE - y - 1
				x = 3*SIZE - 1
			}
		} else {
			x++
		}
	case D22_DOWN:
		if y >= len(m)-1 || m[y+1][x] == M22_VOID {
			if x < SIZE {
				newDir = D22_UP
				x = 3*SIZE - x - 1
				y = 3*SIZE - 1
			} else if x < 2*SIZE {
				newDir = D22_RIGHT
				x = 2 * SIZE
				y = 4*SIZE - x - 1
			} else if x < 3*SIZE {
				newDir = D22_UP
				x = 3*SIZE - x - 1
				y = 2*SIZE - 1
			} else {
				newDir = D22_RIGHT
				x = 0
				y = 5*SIZE - x - 1
			}
		} else {
			y++
		}
	case D22_LEFT:
		if x == 0 || m[y][x-1] == M22_VOID {
			if y < SIZE {
				newDir = D22_DOWN
				y = SIZE
				x = SIZE + y
			} else if y < 2*SIZE {
				newDir = D22_UP
				y = 3*SIZE - 1
				x = 5*SIZE - y - 1
			} else {
				newDir = D22_UP
				y = 2*SIZE - 1
				x = 4*SIZE - y - 1
			}
		} else {
			x--
		}
	case D22_UP:
		if y == 0 || len(m[y-1]) <= x || m[y-1][x] == M22_VOID {
			if x < SIZE {
				newDir = D22_DOWN
				x = 3*SIZE - x - 1
				y = 0
			} else if x < 2*SIZE {
				newDir = D22_RIGHT
				y = x - SIZE
				x = 2 * SIZE
			} else if x < 3*SIZE {
				newDir = D22_DOWN
				x = 3*SIZE - x - 1
				y = SIZE
			} else {
				newDir = D22_LEFT
				y = 5*SIZE - x - 1
				x = 3*SIZE - 1
			}
		} else {
			y--
		}
	}
	if m[y][x] == M22_WALL {
		return pos, dir, false
	}
	return Coord{x, y}, newDir, true
}

func PrintMap22(m [][]int, pos Coord, dir int) {
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[y]); x++ {
			if y == pos.y && x == pos.x {
				switch dir {
				case D22_RIGHT:
					fmt.Print(">")
				case D22_DOWN:
					fmt.Print("v")
				case D22_LEFT:
					fmt.Print("<")
				case D22_UP:
					fmt.Print("^")
				}
				continue
			}
			switch m[y][x] {
			case M22_VOID:
				fmt.Print(" ")
			case M22_EMPTY:
				fmt.Print(".")
			case M22_WALL:
				fmt.Print("#")
			}
		}
		fmt.Println()
	}
}

func Solution22(r io.Reader, mode int) {
	moveFunc := Move22_1
	if mode == 2 {
		moveFunc = Move22_2
	} else if mode == 3 {
		// Part 2 but using example input
		moveFunc = Move22_2e
	}

	m, pos, moves, turns := ReadData22(r)
	dir := 0
	for round := 0; round < len(moves); round++ {
		for i := 0; i < moves[round]; i++ {
			newPos, newDir, ok := moveFunc(m, pos, dir)
			if !ok {
				break
			}
			pos = newPos
			dir = newDir
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
