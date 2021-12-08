// Assumes string contains only ones and zeroes.
package bstrings

import (
	"strconv"
)

// Find gamma in all numbers.
func findGamma(bs []string) string {
	gamma := ""
	for i := 0; i < len(bs[0]); i++ {
		zeroes := []string{}
		ones := []string{}
		for _, v := range bs {
			switch v[i] {
			case '0':
				zeroes = append(zeroes, v)
			case '1':
				ones = append(ones, v)
			}
		}
		if len(zeroes) < len(ones) {
			gamma += "1"
		} else {
			gamma += "0"

		}
	}

	return gamma
}

// Reverse the string (^).
func reverse(bs string) string {
	ret := ""
	for _, bit := range bs {
		if bit == '1' {
			ret += "0"
		} else {
			ret += "1"
		}
	}
	return ret
}

// Parse BinaryString to uint.
func parse(bs string) uint {
	num, err := strconv.ParseUint(bs, 2, len(bs))
	if err != nil {
		panic(err)
	}

	return uint(num)
}

// Return gamma and epsilon.
func CalcGammaEpsilon(bs []string) (uint, uint) {
	gamma := findGamma(bs)
	epsilon := reverse(gamma)

	return parse(gamma), parse(epsilon)
}
// Calculate Oxygen Generator Rating.
func CalcOxygenGeneratorRating(input []string) uint {
	bs := input
	if len(bs) == 0 {
		return 0
	}

	length := len(bs[0])
	for i := 0; i < length; i++ {
		ones := []string{}
		zeroes := []string{}
		for _, v := range bs {
			switch v[i] {
			case '0':
				zeroes = append(zeroes, v)
			case '1':
				ones = append(ones, v)
			}
		}
		if len(zeroes) == len(ones) {
			bs = ones
		} else if len(zeroes) < len(ones) {
			bs = ones
		} else {
			bs = zeroes
		}
	}

	return parse(bs[0])
}

// Calculate CO2 Scrubber Rating.
func CO2ScrubberRating(input []string) uint {
	bs := input
	if len(bs) == 0 {
		return 0
	}

	length := len(bs[0])
	for i := 0; i < length; i++ {
		ones := []string{}
		zeroes := []string{}
		for _, v := range bs {
			switch v[i] {
			case '0':
				zeroes = append(zeroes, v)
			case '1':
				ones = append(ones, v)
			}
		}
		if len(bs) == 1 {
			break
		}
		if len(zeroes) == len(ones) {
			bs = zeroes
		} else if len(zeroes) < len(ones) {
			bs = zeroes
		} else {
			bs = ones
		}
	}

	return parse(bs[0])
}
