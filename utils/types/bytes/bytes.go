package bytes

import (
	"strings"
)

type Bytes []byte

// Convert bytes into strings.
func ToStrings(bytes Bytes, delimiter string) []string {
	str := string(bytes)
	input := strings.Split(str, delimiter)

	// TODO: Figure out why this part is needed.
	if input[len(input)-1] == "" {
		input = input[:len(input)-1]
	}

	return input
}
