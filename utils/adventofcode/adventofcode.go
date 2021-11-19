// Functionality directly related to advent of code webpage.
package adventofcode

import ()

const (
	inputUrl = "https://adventofcode.com/%s/day/%s/input"
	authUrl  = "https://adventofcode.com/auth/github"
)

// Get input directly from Advent of code webpage.
// Year, valid values are "2015" to "2020".
// Day, valid values are "1" to "25".
// Delimiter, how the data should be spliced.
// Returns the input in string splices.
func GetInput(year string, day string, delimiter string) ([]string, error) {
	panic("Not implemented")
}
