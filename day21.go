package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func init() {
	RegisterSolution("21-1", Solution21_1)
	// RegisterSolution("21-2", Solution21_2)
}

type Node21 struct {
	v               int64
	ok              bool
	opLeft, opRight string
	op              byte
}

var registry21 = make(map[string]Node21)

func Eval21(name string) int64 {
	n := registry21[name]
	if n.ok {
		return n.v
	}
	vLeft := Eval21(n.opLeft)
	vRight := Eval21(n.opRight)
	v := int64(0)
	switch n.op {
	case '+':
		v = vLeft + vRight
	case '-':
		v = vLeft - vRight
	case '*':
		v = vLeft * vRight
	case '/':
		v = vLeft / vRight
	}
	registry21[name] = Node21{v, true, "", "", 0}
	return v
}

func Solution21_1(r io.Reader) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		f := strings.Fields(scanner.Text())
		name := strings.TrimSuffix(f[0], ":")
		if len(f) == 2 {
			v, _ := strconv.ParseInt(f[1], 10, 64)
			registry21[name] = Node21{v, true, "", "", 0}
		} else {
			registry21[name] = Node21{0, false, f[1], f[3], f[2][0]}
		}
	}

	fmt.Println(Eval21("root"))
}
