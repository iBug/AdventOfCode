package year

import (
	"bufio"
	"fmt"
	"io"
	"strconv"
)

func init() {
	RegisterSolution("20-1", func(r io.Reader, w io.Writer) { Solution20(r, w, 1) })
	RegisterSolution("20-2", func(r io.Reader, w io.Writer) { Solution20(r, w, 2) })
}

type Node20 struct {
	v          int64
	prev, next *Node20
}

const MAGIC_KEY_20 = 811589153

func MixList20(n []*Node20) {
	for i := 0; i < len(n); i++ {
		p := n[i]
		v := p.v % int64(len(n)-1)
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

func Solution20(r io.Reader, w io.Writer, mode int) {
	mixTimes := 1
	multiplyVBy := int64(1)
	if mode == 2 {
		mixTimes = 10
		multiplyVBy = MAGIC_KEY_20
	}

	scanner := bufio.NewScanner(r)
	n := make([]*Node20, 0, 5000)
	zeroPos := 0
	for scanner.Scan() {
		v, _ := strconv.ParseInt(scanner.Text(), 10, 64)
		v *= multiplyVBy
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

	for i := 0; i < mixTimes; i++ {
		MixList20(n)
	}

	total := int64(0)
	p := n[zeroPos]
	for i := 0; i < 3000; i++ {
		p = p.next
		if (i+1)%1000 == 0 {
			total += p.v
		}
	}

	// PrintList20(n[0], len(n))
	fmt.Fprintln(w, total)
}
