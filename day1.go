package main

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
)

func init() {
	RegisterSolution("1-1", Solution1_1)
	RegisterSolution("1-2", Solution1_2)
}

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

func Solution1_2(r io.Reader) {
	s := bufio.NewScanner(r)
	max := []int{0, 0, 0, 0}
	now := 0
	for s.Scan() {
		line := s.Text()
		if len(line) == 0 {
			max[0] = now
			sort.Ints(max)
			now = 0
			continue
		}
		this, _ := strconv.Atoi(line)
		now += this
	}
	max[0] = now
	sort.Ints(max)
	fmt.Println(max[1] + max[2] + max[3])
}
