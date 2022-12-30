package main

import (
	"bufio"
	"fmt"
	"io"
)

func init() {
	RegisterSolution("6-1", Solution6_1)
	// RegisterSolution("6-2", Solution6_2)
}

func Solution6_1(r io.Reader) {
	scanner := bufio.NewScanner(r)
	scanner.Scan()
	s := scanner.Text()
	b := make(map[byte]int)
	for i := 0; i < 4; i++ {
		b[s[i]-'a']++
	}
	for i := 4; i < len(s); i++ {
		if len(b) == 4 {
			fmt.Println(i)
			break
		}
		b[s[i]-'a']++
		b[s[i-4]-'a']--
		if b[s[i-4]-'a'] == 0 {
			delete(b, s[i-4]-'a')
		}
	}
}
