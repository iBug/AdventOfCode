package year

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func init() {
	RegisterSolution("2-1", Solution2_1)
	RegisterSolution("2-2", Solution2_2)
}

func Solution2_1(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	x, y := 0, 0
	for scanner.Scan() {
		f := strings.Fields(scanner.Text())
		n, _ := strconv.Atoi(f[1])
		switch f[0] {
		case "forward":
			x += n
		case "down":
			y += n
		case "up":
			y -= n
		}
	}
	fmt.Fprintln(w, x*y)
}

func Solution2_2(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	x, y, aim := 0, 0, 0
	for scanner.Scan() {
		f := strings.Fields(scanner.Text())
		n, _ := strconv.Atoi(f[1])
		switch f[0] {
		case "forward":
			x += n
			y += aim * n
		case "down":
			aim += n
		case "up":
			aim -= n
		}
	}
	fmt.Fprintln(w, x*y)
}
