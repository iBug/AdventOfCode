package main

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func init() {
	RegisterSolution("20-1", Solution20_1)
	// RegisterSolution("20-2", Solution20_2)
}

type Node20 struct {
	v          int
	prev, next *Node20
}

func PrintList20(n *Node20, count int) {
	p := n
	fmt.Printf("[%d", p.v)
	for i := 1; i < count; i++ {
		p = p.next
		fmt.Printf(", %d", p.v)
	}
	fmt.Println("]")
}

func Solution20_1(r io.Reader) {
	scanner := bufio.NewScanner(r)
	n := make([]*Node20, 0, 5000)
	zeroPos := 0
	for scanner.Scan() {
		v, _ := strconv.Atoi(scanner.Text())
		n = append(n, &Node20{v: v})
		if v == 0 {
			zeroPos = len(n) - 1
		}
	}

	n[0].prev = n[len(n)-1]
	n[len(n)-1].next = n[0]
	for i := 0; i < len(n)-1; i++ {
		n[i].next = n[i+1]
		n[i+1].prev = n[i]
	}

	for i := 0; i < len(n); i++ {
		p := n[i]
		v := p.v
		if v == 0 {
			continue
		}
		p.prev.next = p.next
		p.next.prev = p.prev

		t := p
		if v < 0 {
			for v < 0 {
				t = t.prev
				v++
			}
			// insert p before t
			p.prev = t.prev
			p.next = t
			t.prev.next = p
			t.prev = p
		} else {
			for v > 0 {
				t = t.next
				v--
			}
			// insert p after t
			p.prev = t
			p.next = t.next
			t.next.prev = p
			t.next = p
		}
	}

	total := 0
	p := n[zeroPos]
	for i := 0; i < 3000; i++ {
		p = p.next
		if (i+1)%1000 == 0 {
			total += p.v
		}
	}

	// PrintList20(n[0], len(n))
	fmt.Println(total)
}
