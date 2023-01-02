package year

import (
	"bufio"
	"fmt"
	"io"
)

func init() {
	RegisterSolution("8-1", Solution8_1)
	RegisterSolution("8-2", Solution8_2)
}

func Solution8_1(r io.Reader) {
	scanner := bufio.NewScanner(r)
	m := make([][]int, 0)
	visible := make([][]bool, 0)
	for scanner.Scan() {
		s := scanner.Text()
		row := make([]int, 0, len(s))
		for _, c := range s {
			row = append(row, int(c-'0'))
		}
		m = append(m, row)
		visible = append(visible, make([]bool, len(s)))
	}

	// vertical
	for x := 0; x < len(m[0]); x++ {
		seen := -1
		for y := 0; y < len(m); y++ {
			if m[y][x] > seen {
				seen = m[y][x]
				visible[y][x] = true
				if seen == 9 {
					break
				}
			}
		}

		seen = -1
		for y := len(m) - 1; y >= 0; y-- {
			if m[y][x] > seen {
				seen = m[y][x]
				visible[y][x] = true
				if seen == 9 {
					break
				}
			}
		}
	}

	// horizontal
	for y := 0; y < len(m); y++ {
		seen := -1
		for x := 0; x < len(m[0]); x++ {
			if m[y][x] > seen {
				seen = m[y][x]
				visible[y][x] = true
				if seen == 9 {
					break
				}
			}
		}

		seen = -1
		for x := len(m[0]) - 1; x >= 0; x-- {
			if m[y][x] > seen {
				seen = m[y][x]
				visible[y][x] = true
				if seen == 9 {
					break
				}
			}
		}
	}

	total := 0
	for y := 0; y < len(m); y++ {
		for x := 0; x < len(m[0]); x++ {
			if visible[y][x] {
				total++
			}
		}
	}
	fmt.Println(total)
}

func Solution8_2(r io.Reader) {
	scanner := bufio.NewScanner(r)
	m := make([][]int, 0)
	for scanner.Scan() {
		s := scanner.Text()
		row := make([]int, 0, len(s))
		for _, c := range s {
			row = append(row, int(c-'0'))
		}
		m = append(m, row)
	}

	max := 0
	for y := 1; y < len(m)-1; y++ {
		for x := 1; x < len(m[0])-1; x++ {
			limit := m[y][x]
			var i, j, k, l int
			for i = 1; i < x; i++ {
				if m[y][x-i] >= limit {
					break
				}
			}
			for j = 1; j < y; j++ {
				if m[y-j][x] >= limit {
					break
				}
			}
			for k = 1; k < len(m[0])-1-x; k++ {
				if m[y][x+k] >= limit {
					break
				}
			}
			for l = 1; l < len(m)-1-y; l++ {
				if m[y+l][x] >= limit {
					break
				}
			}

			score := i * j * k * l
			if score > max {
				max = score
			}
		}
	}
	fmt.Println(max)
}
