package main

import (
	"bufio"
	"fmt"
	"io"
)

func init() {
	RegisterSolution("3-1", Solution3_1)
	// RegisterSolution("3-2", Solution3_2)
}

func Solution3_1(r io.Reader) {
	scanner := bufio.NewScanner(r)
	total := 0
outer:
	for scanner.Scan() {
		s := scanner.Text()
		b := make([]bool, 52)
		mid := len(s) / 2
		for i := 0; i < mid; i++ {
			if s[i] <= 'Z' {
				b[26+s[i]-'A'] = true
			} else {
				b[s[i]-'a'] = true
			}
		}
		for i := mid; i < len(s); i++ {
			if s[i] <= 'Z' {
				if b[26+s[i]-'A'] {
					total += int(27 + s[i] - 'A')
					continue outer
				}
			} else {
				if b[s[i]-'a'] {
					total += int(1 + s[i] - 'a')
					continue outer
				}
			}
		}
	}
	fmt.Println(total)
}
