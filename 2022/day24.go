package year

import (
	"bufio"
	"fmt"
	"io"
)

func init() {
	RegisterSolution("24", Solution24)
	RegisterSolution("24-1", Solution24)
	RegisterSolution("24-2", Solution24)
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

var (
	D24_UP    = Coord{0, -1}
	D24_DOWN  = Coord{0, 1}
	D24_LEFT  = Coord{-1, 0}
	D24_RIGHT = Coord{1, 0}
	D24_STAY  = Coord{0, 0}
)

func TryMove24(m [][]int, size Coord, s State24, c Coord) (State24, bool) {
	if !HasBlizzardAtRound24(m, size, s.round, s.Coord.Add(c)) {
		return State24{s.Coord.Add(c), s.round}, true
	}
	return s, false
}

func Solution24(r io.Reader, w io.Writer) {
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
	stage := 0
	for len(queue) > 0 {
		head := queue[0]
		queue = queue[1:]
		if seen[head] {
			continue
		}
		seen[head] = true
		head.round++

		if head.x == size.x-1 && head.y == size.y-1 {
			if stage == 0 {
				fmt.Fprintln(w, head.round)
				stage++
				// reset data, start moving back
				queue = []State24{{Coord{size.x - 1, size.y}, head.round}}
				seen = make(map[State24]bool)
				continue
			}
			if stage == 2 {
				fmt.Fprintln(w, head.round)
				break
			}
		}
		if head.x == 0 && head.y == 0 {
			if stage == 1 {
				stage++
				// reset data, start moving forward again
				queue = []State24{{Coord{0, -1}, head.round}}
				seen = make(map[State24]bool)
				continue
			}
		}

		if head.y < 0 {
			if s, ok := TryMove24(m, size, head, D24_DOWN); ok {
				queue = append(queue, s)
			}
			if s, ok := TryMove24(m, size, head, D24_STAY); ok {
				queue = append(queue, s)
			}
			continue
		}
		if head.y == size.y {
			if s, ok := TryMove24(m, size, head, D24_UP); ok {
				queue = append(queue, s)
			}
			if s, ok := TryMove24(m, size, head, D24_STAY); ok {
				queue = append(queue, s)
			}
			continue
		}

		if head.x < size.x-1 {
			if s, ok := TryMove24(m, size, head, D24_RIGHT); ok {
				queue = append(queue, s)
			}
		}
		if head.y < size.y-1 {
			if s, ok := TryMove24(m, size, head, D24_DOWN); ok {
				queue = append(queue, s)
			}
		}
		if head.x > 0 {
			if s, ok := TryMove24(m, size, head, D24_LEFT); ok {
				queue = append(queue, s)
			}
		}
		if head.y > 0 {
			if s, ok := TryMove24(m, size, head, D24_UP); ok {
				queue = append(queue, s)
			}
		}
		if s, ok := TryMove24(m, size, head, D24_STAY); ok {
			queue = append(queue, s)
		}
	}
}
