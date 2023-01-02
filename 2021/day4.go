package year

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func init() {
	RegisterSolution("4-1", func(r io.Reader) { Solution4(r, 1) })
	RegisterSolution("4-2", func(r io.Reader) { Solution4(r, 2) })
}

func Solution4(r io.Reader, mode int) {
	scanner := bufio.NewScanner(r)
	draws := make([]int, 0)
	scanner.Scan()
	for _, c := range strings.Split(scanner.Text(), ",") {
		n, _ := strconv.Atoi(c)
		draws = append(draws, n)
	}
	boards := make([][]int, 0)
	board := make([]int, 0, 25)
	for scanner.Scan() {
		if scanner.Text() == "" {
			continue
		}
		for _, c := range strings.Fields(scanner.Text()) {
			n, _ := strconv.Atoi(c)
			board = append(board, n)
		}
		if len(board) == 25 {
			boards = append(boards, board)
			board = make([]int, 0, 25)
		}
	}

	won := make([]bool, len(boards))
	boardsWon := 0
	for _, draw := range draws {
		for b, board := range boards {
			if won[b] {
				continue
			}
			for i, n := range board {
				if n == draw {
					board[i] = -1

					rowFull, colFull := true, true
					for j := 0; j < 5; j++ {
						if board[i/5*5+j] != -1 {
							rowFull = false
						}
						if board[i%5+j*5] != -1 {
							colFull = false
						}
					}

					if rowFull || colFull {
						boardsWon++
						won[b] = true

						if (mode == 1) ||
							(mode == 2 && boardsWon == len(boards)) {
							sum := 0
							for _, n := range board {
								if n != -1 {
									sum += n
								}
							}
							fmt.Println(n * sum)
							return
						}
					}
					break
				}
			}
		}
	}
}
