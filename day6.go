package main

import (
	"bufio"
	"fmt"
	"io"
)

func init() {
	RegisterSolution("6-1", func(r io.Reader) { Solution6(r, 4) })
	RegisterSolution("6-2", func(r io.Reader) { Solution6(r, 14) })
}

func Solution6(r io.Reader, thresh int) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	s := scanner.Text()
	b := make(map[byte]int)
	for i := 0; i < thresh; i++ {
		b[s[i]-'a']++
	}
	for i := thresh; i < len(s); i++ {
		if len(b) == thresh {
			fmt.Println(i)
			break
		}
		b[s[i]-'a']++
		b[s[i-thresh]-'a']--
		if b[s[i-thresh]-'a'] == 0 {
			delete(b, s[i-thresh]-'a')
		}
	}
}
