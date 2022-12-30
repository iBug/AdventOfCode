package main

import (
	"bufio"
	"fmt"
	"io"
	"strings"
)

func init() {
	RegisterSolution("16-1", Solution16_1)
	// RegisterSolution("16-2", Solution16_2)
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

func Solution16_1(r io.Reader) {
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
		}
	}
	fmt.Println(Recurse16_1(v, d, "AA", 30, allValves))
}
