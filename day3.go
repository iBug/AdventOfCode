package main

import (
	"bufio"
	"fmt"
	"io"
)

func init() {
	RegisterSolution("3-1", Solution3_1)
	RegisterSolution("3-2", Solution3_2)
}

func Day3_priority(c byte) int {
	if c <= 'Z' {
		return int(26 + c - 'A')
	} else {
		return int(c - 'a')
	}
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
			b[Day3_priority(s[i])] = true
		}
		for i := mid; i < len(s); i++ {
			prio := Day3_priority(s[i])
			if b[prio] {
				total += 1 + prio
				continue outer
			}
		}
	}
	fmt.Println(total)
}

func Solution3_2(r io.Reader) {
	scanner := bufio.NewScanner(r)
	total := 0
outer:
	for scanner.Scan() {
		b := make([]byte, 52)
		for _, c := range scanner.Text() {
			b[Day3_priority(byte(c))] = 1
		}
		scanner.Scan()
		for _, c := range scanner.Text() {
			b[Day3_priority(byte(c))] |= 2
		}
		scanner.Scan()
		for _, c := range scanner.Text() {
			if b[Day3_priority(byte(c))] == 3 {
				total += 1 + Day3_priority(byte(c))
				continue outer
			}
		}
	}
	fmt.Println(total)
}
