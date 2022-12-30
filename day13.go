package main

import (
	"bufio"
	"encoding/json"
	"fmt"
	"io"
)

func init() {
	RegisterSolution("13-1", Solution13_1)
	// RegisterSolution("13-2", Solution13_2)
}

const (
	C13_EQUAL = iota
	C13_LEFT
	C13_RIGHT
)

func CompareSingle13_1(left_i, right_i interface{}) int {
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
		return CompareSingle13_1([]interface{}{left_i}, right_i)
	}
	if !l_ok && r_ok {
		return CompareSingle13_1(left_i, []interface{}{right_i})
	}

	left_l, l_ok := left_i.([]interface{})
	right_l, r_ok := right_i.([]interface{})
	if !l_ok || !r_ok {
		panic("bad type")
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

func Compare13_1(left_s, right_s string) bool {
	var left, right interface{}
	json.Unmarshal([]byte(left_s), &left)
	json.Unmarshal([]byte(right_s), &right)
	return CompareSingle13_1(left, right) == C13_RIGHT
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
