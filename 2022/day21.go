package year

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func init() {
	RegisterSolution("21-1", Solution21_1)
	RegisterSolution("21-2", Solution21_2)
}

type Node21 struct {
	// y = ax + v
	a, v            float64
	ok              bool
	opLeft, opRight string
	op              byte
}

var registry21 = make(map[string]Node21)

func Eval21(name string) float64 {
	n := registry21[name]
	if n.ok {
		return n.v
	}
	vLeft := Eval21(n.opLeft)
	vRight := Eval21(n.opRight)
	aLeft := registry21[n.opLeft].a
	aRight := registry21[n.opRight].a
	a, v := 0.0, 0.0
	switch n.op {
	case '+':
		a = aLeft + aRight
		v = vLeft + vRight
	case '-':
		a = aLeft - aRight
		v = vLeft - vRight
	case '*':
		a = aLeft*vRight + aRight*vLeft
		v = vLeft * vRight
	case '/':
		a = aLeft / vRight
		v = vLeft / vRight
	}
	registry21[name] = Node21{a, v, true, n.opLeft, n.opRight, n.op}
	return v
}

func Solution21_1(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		f := strings.Fields(scanner.Text())
		name := strings.TrimSuffix(f[0], ":")
		if len(f) == 2 {
			v, _ := strconv.ParseFloat(f[1], 64)
			registry21[name] = Node21{0.0, v, true, "", "", 0}
		} else {
			registry21[name] = Node21{0.0, 0.0, false, f[1], f[3], f[2][0]}
		}
	}
	fmt.Printf("%.0f\n", Eval21("root"))
}

func Solution21_2(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		f := strings.Fields(scanner.Text())
		name := strings.TrimSuffix(f[0], ":")
		if name == "humn" {
			registry21[name] = Node21{1.0, 0.0, true, "", "", 0}
		} else if len(f) == 2 {
			v, _ := strconv.ParseFloat(f[1], 64)
			registry21[name] = Node21{0.0, v, true, "", "", 0}
		} else {
			registry21[name] = Node21{0.0, 0.0, false, f[1], f[3], f[2][0]}
		}
	}
	Eval21("root")
	rootLeft := registry21["root"].opLeft
	rootRight := registry21["root"].opRight

	// very bold assumption: "root" is a linear function
	// and the RHS has no coefficient for the variable
	fmt.Printf("%.0f\n", (registry21[rootRight].v-registry21[rootLeft].v)/registry21[rootLeft].a)
}
