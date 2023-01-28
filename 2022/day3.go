package year

import (
	"bufio"
	"fmt"
	"io"
)

func init() {
	RegisterSolution("3-1", Solution3_1)
	RegisterSolution("3-2", Solution3_2)
}

func Priority3(c byte) int {
	if c <= 'Z' {
		return int(26 + c - 'A')
	} else {
		return int(c - 'a')
	}
}

func Solution3_1(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	total := 0
outer:
	for scanner.Scan() {
		s := scanner.Text()
		b := make([]bool, 52)
		mid := len(s) / 2
		for i := 0; i < mid; i++ {
			b[Priority3(s[i])] = true
		}
		for i := mid; i < len(s); i++ {
			prio := Priority3(s[i])
			if b[prio] {
				total += 1 + prio
				continue outer
			}
		}
	}
	fmt.Fprintln(w, total)
}

func Solution3_2(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	total := 0
outer:
	for scanner.Scan() {
		b := make([]byte, 52)
		for _, c := range scanner.Text() {
			b[Priority3(byte(c))] = 1
		}
		scanner.Scan()
		for _, c := range scanner.Text() {
			b[Priority3(byte(c))] |= 2
		}
		scanner.Scan()
		for _, c := range scanner.Text() {
			if b[Priority3(byte(c))] == 3 {
				total += 1 + Priority3(byte(c))
				continue outer
			}
		}
	}
	fmt.Fprintln(w, total)
}
