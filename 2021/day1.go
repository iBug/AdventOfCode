package year

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"strconv"
)

func init() {
	RegisterSolution("1-1", Solution1_1)
	RegisterSolution("1-2", Solution1_2)
}

func Solution1_1(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	last := math.MaxInt
	count := 0
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		if n > last {
			count++
		}
		last = n
	}
	fmt.Fprintln(w, count)
}

func Solution1_2(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	last := []int{math.MaxInt, math.MaxInt, math.MaxInt}
	count := 0
	i := 0
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		if n > last[i%3] {
			count++
		}
		last[i%3] = n
		i++
	}
	fmt.Fprintln(w, count)
}
