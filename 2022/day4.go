package year

import (
	"bufio"
	"fmt"
	"io"
)

func init() {
	RegisterSolution("4-1", Solution4_1)
	RegisterSolution("4-2", Solution4_2)
}

func Solution4_1(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	total := 0
	for scanner.Scan() {
		var a, b, c, d int
		fmt.Sscanf(scanner.Text(), "%d-%d,%d-%d", &a, &b, &c, &d)
		if (a <= c && b >= d) || (c <= a && d >= b) {
			total++
		}
	}
	fmt.Fprintln(w, total)
}

func Solution4_2(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	total := 0
	for scanner.Scan() {
		var a, b, c, d int
		fmt.Sscanf(scanner.Text(), "%d-%d,%d-%d", &a, &b, &c, &d)
		if (a <= c && c <= b) || (c <= a && a <= d) {
			total++
		}
	}
	fmt.Fprintln(w, total)
}
