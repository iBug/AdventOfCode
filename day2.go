package main

import (
	"bufio"
	"fmt"
	"io"
)

func init() {
	RegisterSolution("2-1", Solution2_1)
	RegisterSolution("2-2", Solution2_2)
}

func Solution2_1(r io.Reader) {
	scanner := bufio.NewScanner(r)
	total := 0
	for scanner.Scan() {
		s := scanner.Text()
		total += int(s[2] - 'W')
		total += (4 + int(s[2]-'X') - int(s[0]-'A')) % 3 * 3
	}
	fmt.Println(total)
}

func Solution2_2(r io.Reader) {
	scanner := bufio.NewScanner(r)
	total := 0
	for scanner.Scan() {
		s := scanner.Text()
		total += 1 + (int(s[0]-'A')+int(s[2]-'X')+2)%3
		total += int(s[2]-'X') * 3
	}
	fmt.Println(total)
}
