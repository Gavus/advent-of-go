package ints

import (
	"errors"
	"strconv"
)

type Ints []int

func (s Ints) Contains(e int) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (s Ints) Sum() int {
	sum := 0
	for _, v := range s {
		sum += v
	}

	return sum
}

func (a Ints) Equal(b Ints) bool {
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

// Convert strings to ints
func ToInts(input []string) (Ints, error) {
	ret := Ints{}
	for _, valStr := range input {
		valInt, err := strconv.Atoi(valStr)
		if err != nil {
			return nil, errors.New("Failed to convert string into int")
		}
		ret = append(ret, valInt)
	}
	return ret, nil
}
