package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
	"sort"
	"strings"
)

func init() {
	RegisterSolution("13-1", Solution13_1)
	RegisterSolution("13-2", Solution13_2)
}

const (
	C13_EQUAL = iota
	C13_LEFT
	C13_RIGHT
)

func CompareSingle13_1(left_i, right_i any) int {
	left_f, l_ok := left_i.(float64)
	right_f, r_ok := right_i.(float64)
	if l_ok && r_ok {
		left, right := int(left_f), int(right_f)
		if left == right {
			return C13_EQUAL
		}
		if left < right {
			return C13_RIGHT
		}
		return C13_LEFT
	}

	if l_ok && !r_ok {
		return CompareSingle13_1([]any{left_i}, right_i)
	}
	if !l_ok && r_ok {
		return CompareSingle13_1(left_i, []any{right_i})
	}

	left_l, l_ok := left_i.([]any)
	right_l, r_ok := right_i.([]any)
	if !l_ok || !r_ok {
		panic(fmt.Sprintf("bad types: %T and %T", left_i, right_i))
	}
	for i := 0; i < len(left_l); i++ {
		if i >= len(right_l) {
			return C13_LEFT
		}
		res := CompareSingle13_1(left_l[i], right_l[i])
		if res != C13_EQUAL {
			return res
		}
	}
	if len(left_l) == len(right_l) {
		return C13_EQUAL
	}
	return C13_RIGHT
}

func Parse13_1(s string) any {
	var res any
	json.Unmarshal([]byte(s), &res)
	return res
}

func Compare13_1(left, right string) bool {
	return CompareSingle13_1(Parse13_1(left), Parse13_1(right)) == C13_RIGHT
}

func Solution13_1(r io.Reader) {
	scanner := bufio.NewScanner(r)
	total := 0
	n := 0
	for scanner.Scan() {
		n++
		left := scanner.Text()
		scanner.Scan()
		right := scanner.Text()
		scanner.Scan()
		if Compare13_1(left, right) {
			total += n
		}
	}
	fmt.Println(total)
}

func Solution13_2(r io.Reader) {
	scanner := bufio.NewScanner(r)
	p1 := Parse13_1("[[2]]")
	p2 := Parse13_1("[[6]]")
	packets := []any{p1, p2}
	for scanner.Scan() {
		s := scanner.Text()
		if strings.TrimSpace(s) == "" {
			continue
		}
		packets = append(packets, Parse13_1(s))
	}
	sort.Slice(packets, func(i, j int) bool {
		return CompareSingle13_1(packets[i], packets[j]) == C13_RIGHT
	})
	res := 1
	for i := 0; i < len(packets); i++ {
		if CompareSingle13_1(packets[i], p1) == C13_EQUAL {
			res *= i + 1
		}
		if CompareSingle13_1(packets[i], p2) == C13_EQUAL {
			res *= i + 1
		}
	}
	fmt.Println(res)
}
