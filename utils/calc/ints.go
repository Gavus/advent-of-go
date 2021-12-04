package calc

import (
	"fmt"
)

func Contains(s []int, e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func Sum(s []int) int {
	sum := 0
	for _, v := range s {
		sum += v
	}

	return sum
}

func Equal(a []int, b []int) bool {
	if len(a) != len(b) {
		return false
	}

	for i, v := range a {
		if b[i] != v {
			return false
		}
	}

	return true
}

func String(a []int) string {
	str := ""
	for _, v := range a {
		str += fmt.Sprintf("%d, ", v)
	}
	return str
}
