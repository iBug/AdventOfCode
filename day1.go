package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func Solution1_1(r io.Reader) {
	s := bufio.NewScanner(r)
	max := 0
	now := 0
	for s.Scan() {
		line := s.Text()
		if len(line) == 0 {
			if now > max {
				max = now
			}
			now = 0
			continue
		}
		this, _ := strconv.Atoi(line)
		now += this
	}
	if now > max {
		max = now
	}
	fmt.Println(max)
}

func init() {
	RegisterSolution("1-1", Solution1_1)
}
