package year

import (
	"bufio"
	"fmt"
	"io"
)

func init() {
	RegisterSolution("25", Solution25)
	RegisterSolution("25-1", Solution25)
}

func FromSnafuDigit25(c byte) int {
	switch c {
	case '=':
		return -2
	case '-':
		return -1
	case '1':
		return 1
	case '2':
		return 2
	default:
		return 0
	}
}

func ToSnafuDigit25(n int) byte {
	switch n {
	case -2:
		return '='
	case -1:
		return '-'
	case 1:
		return '1'
	case 2:
		return '2'
	default:
		return '0'
	}
}

func ParseSnafu25(s string) int64 {
	result := int64(0)
	for _, c := range s {
		result = 5*result + int64(FromSnafuDigit25(byte(c)))
	}
	return result
}

func ToSnafu25(n int64) string {
	digits := make([]int, 0, 10)
	for n > 0 {
		digits = append(digits, int(n%5))
		n /= 5
	}
	for i := 0; i < len(digits)-1; i++ {
		if digits[i] >= 3 {
			digits[i] -= 5
			digits[i+1]++
		}
	}
	if digits[len(digits)-1] >= 3 {
		digits[len(digits)-1] -= 5
		digits = append(digits, 1)
	}
	s := make([]byte, 0, len(digits))
	for i := len(digits) - 1; i >= 0; i-- {
		s = append(s, ToSnafuDigit25(digits[i]))
	}
	return string(s)
}

func Solution25(r io.Reader, w io.Writer) {
	scanner := bufio.NewScanner(r)
	total := int64(0)
	for scanner.Scan() {
		total += ParseSnafu25(scanner.Text())
	}
	fmt.Fprintln(w, total, ToSnafu25(total))
}
