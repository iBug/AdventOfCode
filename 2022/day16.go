package year

import (
	"bufio"
	"fmt"
	"io"
	"sort"
	"strings"
)

func init() {
	RegisterSolution("16-1", func(r io.Reader) { Solution16(r, 1) })
	RegisterSolution("16-2", func(r io.Reader) { Solution16(r, 2) })
}

type Valve16 struct {
	flow int
	adj  []string
}

func Recurse16_1(v map[string]Valve16, d map[string]int, node string, time int, remaining []string) int {
	thisValue := v[node].flow * time
	next := make([]string, 0, 16)
	for _, c := range remaining {
		if d[node+c] < time {
			next = append(next, c)
		}
	}
	maxNext := 0
	for i, c := range next {
		newNext := make([]string, 0, len(next)-1)
		newNext = append(newNext, next[:i]...)
		newNext = append(newNext, next[i+1:]...)

		nextValue := Recurse16_1(v, d, c, time-d[node+c], newNext)
		if nextValue > maxNext {
			maxNext = nextValue
		}
	}
	return thisValue + maxNext
}

func Recurse16_2(v map[string]Valve16, d map[string]int, node string, time, value int, remaining, visited []string, resultMap map[string]int) {
	value += v[node].flow * time
	visited = append(visited, node)

	next := make([]string, 0, 16)
	for _, c := range remaining {
		if d[node+c] < time {
			next = append(next, c)
		}
	}
	for i, c := range next {
		newNext := make([]string, 0, len(next)-1)
		newNext = append(newNext, next[:i]...)
		newNext = append(newNext, next[i+1:]...)

		Recurse16_2(v, d, c, time-d[node+c], value, newNext, visited, resultMap)
	}

	visited2 := make([]string, len(visited))
	copy(visited2, visited)
	sort.Strings(visited2)
	if resultMap[strings.Join(visited2, ",")] < value {
		resultMap[strings.Join(visited2, ",")] = value
	}
}

func Solution16(r io.Reader, mode int) {
	scanner := bufio.NewScanner(r)
	v := make(map[string]Valve16)
	for scanner.Scan() {
		f := strings.Fields(scanner.Text())
		s := f[1]
		var flow int
		fmt.Sscanf(f[4], "rate=%d;", &flow)
		adj := make([]string, 0, 8)
		for i := 9; i < len(f); i++ {
			adj = append(adj, strings.ReplaceAll(f[i], ",", ""))
		}
		v[s] = Valve16{flow, adj}
	}

	d := make(map[string]int)
	for i := range v {
		for j := range v {
			d[i+j] = 999
		}
	}
	for valve := range v {
		// +1 for the cost of opening the valve
		d[valve+valve] = 1
		for _, adj := range v[valve].adj {
			d[valve+adj] = 2
		}
	}
	for k := range v {
		for i := range v {
			for j := range v {
				if d[i+k]+d[k+j]-1 < d[i+j] {
					d[i+j] = d[i+k] + d[k+j] - 1
				}
			}
		}
	}

	allValves := make([]string, 0, len(v)/2)
	for name, valve := range v {
		if valve.flow > 0 {
			allValves = append(allValves, name)
		} else {
			delete(v, name)
		}
	}
	if mode == 1 {
		fmt.Println(Recurse16_1(v, d, "AA", 30, allValves))
	} else if mode == 2 {
		m := make(map[string]int)
		Recurse16_2(v, d, "AA", 26, 0, allValves, make([]string, 0, 16), m)
		max := 0
		for res1 := range m {
		res2:
			for res2 := range m {
				for _, part := range strings.Split(res1, ",") {
					if part != "AA" && strings.Contains(res2, part) {
						continue res2
					}
				}
				if m[res1]+m[res2] > max {
					max = m[res1] + m[res2]
				}
			}
		}
		for _, value := range m {
			if value > max {
				max = value
			}
		}
		fmt.Println(max)
	}
}
