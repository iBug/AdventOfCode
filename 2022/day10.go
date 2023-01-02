package year

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
	"strings"
)

func init() {
	RegisterSolution("10-1", Solution10_1)
	RegisterSolution("10-2", Solution10_2)
}

func Solution10_1(r io.Reader) {
	scanner := bufio.NewScanner(r)
	total := 0
	tick := 0
	x := 1
	for scanner.Scan() {
		f := strings.Fields(scanner.Text())
		switch f[0] {
		case "noop":
			tick += 1
			if tick%40 == 20 {
				total += tick * x
			}
		case "addx":
			tick += 2
			value, _ := strconv.Atoi(f[1])
			if tick%40 == 20 {
				total += tick * x
			} else if tick%40 == 21 {
				total += (tick - 1) * x
			}
			x += value
		}
		if tick >= 220 {
			break
		}
	}
	fmt.Println(total)
}

func Solution10_2(r io.Reader) {
	scanner := bufio.NewScanner(r)
	tick := 0
	x := 1

	advanceTick := func() {
		c := ' ' // better visibility than '.'
		if x-1 <= tick%40 && tick%40 <= x+1 {
			c = '#'
		}
		fmt.Printf("%c", c)
		tick += 1
		if tick%40 == 0 {
			fmt.Printf("\n")
		}
	}

	for scanner.Scan() {
		f := strings.Fields(scanner.Text())
		switch f[0] {
		case "noop":
			advanceTick()
		case "addx":
			advanceTick()
			advanceTick()
			value, _ := strconv.Atoi(f[1])
			x += value
		}
		if tick >= 240 {
			break
		}
	}
}
