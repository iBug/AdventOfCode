package main

import (
	"bufio"
	"fmt"
	"io"
)

func init() {
	RegisterSolution("24-1", Solution24_1)
	// RegisterSolution("24-2", Solution24_2)
}

const (
	B24_EMPTY = iota
	B24_UP
	B24_DOWN
	B24_LEFT
	B24_RIGHT
)

type State24 struct {
	Coord
	round int
}

func (s State24) String() string {
	return fmt.Sprintf("%d: %v", s.round, s.Coord)
}

func HasBlizzardAtRound24(m [][]int, size Coord, round int, c Coord) bool {
	width, height := size.x, size.y
	if c.y < 0 || c.y >= height {
		return false
	}
	if m[(c.y+height-(round%height))%height][c.x] == B24_DOWN {
		return true
	}
	if m[(c.y+round)%height][c.x] == B24_UP {
		return true
	}
	if m[c.y][(c.x+width-(round%width))%width] == B24_RIGHT {
		return true
	}
	if m[c.y][(c.x+round)%width] == B24_LEFT {
		return true
	}
	return false
}

func Solution24_1(r io.Reader) {
	scanner := bufio.NewScanner(r)
	m := make([][]int, 0)
	size := Coord{}
	scanner.Scan()
	size.x = len(scanner.Text()) - 2
	for scanner.Scan() {
		s := scanner.Text()
		if s[1] == '#' {
			size.y = len(m)
			break
		}
		row := make([]int, size.x)
		for i := 0; i < size.x; i++ {
			switch s[i+1] {
			case '^':
				row[i] = B24_UP
			case 'v':
				row[i] = B24_DOWN
			case '<':
				row[i] = B24_LEFT
			case '>':
				row[i] = B24_RIGHT
			default:
				row[i] = B24_EMPTY
			}
		}
		m = append(m, row)
	}

	queue := []State24{{Coord{0, -1}, 0}}
	seen := make(map[State24]bool)
	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]
		if seen[head] {
			continue
		}
		seen[head] = true
		head.round++

		if head.x == size.x-1 && head.y == size.y-1 {
			fmt.Println(head.round)
			break
		}

		if head.y < 0 {
			if !HasBlizzardAtRound24(m, size, head.round, head.Coord) {
				queue = append(queue, State24{Coord{head.x, head.y}, head.round})
			}
			if !HasBlizzardAtRound24(m, size, head.round, head.Coord.Add(Coord{0, 1})) {
				queue = append(queue, State24{Coord{head.x, head.y + 1}, head.round})
			}
			continue
		}

		// move right
		if head.x < size.x-1 && !HasBlizzardAtRound24(m, size, head.round, head.Coord.Add(Coord{1, 0})) {
			queue = append(queue, State24{Coord{head.x + 1, head.y}, head.round})
		}
		// move down
		if head.y < size.y-1 && !HasBlizzardAtRound24(m, size, head.round, head.Coord.Add(Coord{0, 1})) {
			queue = append(queue, State24{Coord{head.x, head.y + 1}, head.round})
		}
		// move left
		if head.x > 0 && !HasBlizzardAtRound24(m, size, head.round, head.Coord.Add(Coord{-1, 0})) {
			queue = append(queue, State24{Coord{head.x - 1, head.y}, head.round})
		}
		// move up
		if head.y > 0 && !HasBlizzardAtRound24(m, size, head.round, head.Coord.Add(Coord{0, -1})) {
			queue = append(queue, State24{Coord{head.x, head.y - 1}, head.round})
		}
		// stay
		if !HasBlizzardAtRound24(m, size, head.round, head.Coord) {
			queue = append(queue, State24{Coord{head.x, head.y}, head.round})
		}
	}
}
