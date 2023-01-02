package year

import (
	"bufio"
	"fmt"
	"io"
)

func init() {
	RegisterSolution("3-1", func(r io.Reader) { Solution3(r, 1) })
	RegisterSolution("3-2", func(r io.Reader) { Solution3(r, 2) })
}

func AddOnes3(count []int, s string) {
	for i, c := range s {
		if c == '1' {
			count[i]++
		}
	}
}

func CountToInt3(count []int) int {
	result := 0
	for _, c := range count {
		result *= 2
		if c > 0 {
			result++
		}
	}
	return result
}

func Solution3(r io.Reader, mode int) {
	scanner := bufio.NewScanner(r)

	scanner.Scan()
	count := make([]int, len(scanner.Text()))
	AddOnes3(count, scanner.Text())
	numbers := []string{scanner.Text()}
	for scanner.Scan() {
		AddOnes3(count, scanner.Text())
		numbers = append(numbers, scanner.Text())
	}

	if mode == 1 {
		for i, c := range count {
			if c >= len(numbers)/2 {
				count[i] = 1
			} else {
				count[i] = 0
			}
		}
		result := CountToInt3(count)
		result *= (1<<len(count) - 1) ^ result
		fmt.Println(result)
		return
	}

	numbersA, numbersB := numbers, numbers
	countA, countB := count, count
	for i := range count {
		wantA, wantB := byte('0'), byte('0')
		if countA[i] >= (len(numbersA)+1)/2 {
			wantA = '1'
		}
		if countB[i] < (len(numbersB)+1)/2 {
			wantB = '1'
		}

		if len(numbersA) > 1 {
			newA := make([]string, 0, len(numbersA))
			newCountA := make([]int, len(countA))
			for _, s := range numbersA {
				if s[i] == wantA {
					newA = append(newA, s)
					AddOnes3(newCountA, s)
				}
			}
			numbersA, countA = newA, newCountA
		}

		if len(numbersB) > 1 {
			newB := make([]string, 0, len(numbersB))
			newCountB := make([]int, len(countB))
			for _, s := range numbersB {
				if s[i] == wantB {
					newB = append(newB, s)
					AddOnes3(newCountB, s)
				}
			}
			numbersB, countB = newB, newCountB
		}
	}

	resultA := CountToInt3(countA)
	resultB := CountToInt3(countB)
	fmt.Println(resultA * resultB)
}
