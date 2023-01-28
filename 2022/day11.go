package year

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strconv"
	"strings"
)

func init() {
	RegisterSolution("11-1", func(r io.Reader, w io.Writer) { Solution11(r, w, 1, 20) })
	RegisterSolution("11-2", func(r io.Reader, w io.Writer) { Solution11(r, w, 2, 10000) })
}

type Monkey11 struct {
	items []int
	op    byte
	opVal int // special value -1 = square
	div   int
	ifT   int
	ifF   int
	count int
}

func Gcd(a, b int) int {
	for b != 0 {
		a, b = b, a%b
	}
	return a
}

func Lcm(a, b int) int {
	return a * b / Gcd(a, b)
}

func Solution11(r io.Reader, w io.Writer, mode, roundsN int) {
	m := make([]Monkey11, 0, 8)
	n := 0

	scanner := bufio.NewScanner(r)
	for scanner.Scan() {
		parts := strings.SplitN(strings.TrimSpace(scanner.Text()), ": ", 2)
		if strings.HasPrefix(parts[0], "Monkey ") {
			n = len(m)
			m = append(m, Monkey11{})
			continue
		}

		switch parts[0] {
		case "Starting items":
			m[n].items = make([]int, 0, 8)
			for _, item := range strings.Split(parts[1], ", ") {
				val, _ := strconv.Atoi(item)
				m[n].items = append(m[n].items, val)
			}
		case "Operation":
			f := strings.Fields(parts[1])
			m[n].op = f[3][0]
			if f[4] == "old" {
				m[n].opVal = -1
			} else {
				m[n].opVal, _ = strconv.Atoi(f[4])
			}
		case "Test":
			m[n].div, _ = strconv.Atoi(strings.Fields(parts[1])[2])
		case "If true":
			m[n].ifT, _ = strconv.Atoi(strings.Fields(parts[1])[3])
		case "If false":
			m[n].ifF, _ = strconv.Atoi(strings.Fields(parts[1])[3])
		}
	}

	// reuse n as modulo
	n = Lcm(m[0].div, m[1].div)
	for i := 2; i < len(m); i++ {
		n = Lcm(n, m[i].div)
	}

	for i := 0; i < roundsN; i++ {
		for j := 0; j < len(m); j++ {
			m[j].count += len(m[j].items)
			for _, item := range m[j].items {
				switch m[j].op {
				case '+':
					if m[j].opVal == -1 {
						item += item
					} else {
						item += m[j].opVal
					}
				case '*':
					if m[j].opVal == -1 {
						item *= item
					} else {
						item *= m[j].opVal
					}
				}
				if mode == 1 {
					item /= 3
				} else if mode == 2 {
					item %= n
				}

				target := m[j].ifF
				if item%m[j].div == 0 {
					target = m[j].ifT
				}
				m[target].items = append(m[target].items, item)

			}
			m[j].items = make([]int, 0, 8)
		}
	}

	counts := make([]int, len(m))
	for i := 0; i < len(m); i++ {
		counts[i] = m[i].count
	}
	sort.Ints(counts)
	fmt.Fprintln(w, counts[len(m)-2]*counts[len(m)-1])
}
