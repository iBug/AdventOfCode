package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func init() {
	RegisterSolution("5-1", Solution5_1)
	// RegisterSolution("5-2", Solution5_2)
}

func Solution5_1(r io.Reader) {
	scanner := bufio.NewScanner(r)
	stacks_s := make([]string, 0)
	n := 0
	for scanner.Scan() {
		s := scanner.Text()
		if !strings.Contains(s, "[") {
			ss := strings.Fields(s)
			n, _ = strconv.Atoi(ss[len(ss)-1])
			break
		}
		stacks_s = append(stacks_s, s)
	}
	stacks := make([][]byte, n)
	for i := len(stacks_s) - 1; i >= 0; i-- {
		for j := 0; j < n; j++ {
			idx := 4*j + 1
			if stacks_s[i][idx] == ' ' {
				continue
			}
			stacks[j] = append(stacks[j], stacks_s[i][idx])
		}
	}

	scanner.Scan() // Empty line
	for scanner.Scan() {
		var count, from, to int
		fmt.Sscanf(scanner.Text(), "move %d from %d to %d", &count, &from, &to)
		from--
		to--
		for i := 0; i < count; i++ {
			stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-1])
			stacks[from] = stacks[from][:len(stacks[from])-1]
		}
	}

	for i := 0; i < n; i++ {
		fmt.Printf("%d:", i+1)
		for j := 0; j < len(stacks[i]); j++ {
			fmt.Printf(" %c", stacks[i][j])
		}
		fmt.Println()
	}

	s := ""
	for i := 0; i < n; i++ {
		s += string(stacks[i][len(stacks[i])-1])
	}
	fmt.Println(s)
}
