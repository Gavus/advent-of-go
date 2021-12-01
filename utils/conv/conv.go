// Converting functionalities.
package conv

import (
	"crypto/md5"
	"encoding/hex"
	"errors"
	"fmt"
	"github.com/Gavus/advent-of-code/utils/types"
	"strings"
)

// Convert strings into Presents.
func ToPresents(input []string) ([]types.Present, error) {
	presents := make([]types.Present, len(input))

	for index, value := range input {
		var l, w, h int
		_, err := fmt.Sscanf(value, "%dx%dx%d", &l, &w, &h)
		if err != nil {
			return []types.Present{}, err
		}
		presents[index].Set(l, w, h)
	}

	return presents, nil
}

// Convert bytes into strings.
func ToStrings(bytes []byte, delimiter string) []string {
	str := string(bytes)
	input := strings.Split(str, delimiter)

	// TODO: Figure out why this part is needed.
	if input[len(input)-1] == "" {
		input = input[:len(input)-1]
	}

	return input
}

// Convert strings to Directions.
func ToDirections(input []string) ([]types.Direction, error) {
	dirs := make([]types.Direction, len(input))
	for i, v := range input {
		switch v {
		case "<":
			dirs[i] = types.DirLeft
		case ">":
			dirs[i] = types.DirRight
		case "^":
			dirs[i] = types.DirUp
		case "v":
			dirs[i] = types.DirDown
		default:
			return []types.Direction{}, errors.New("Invalid character found in input")
		}
	}
	return dirs, nil
}

// Convert bytes to string representation of the hex.
func ToString(bytes []byte) string {
	return hex.EncodeToString(bytes)
}

// Convert string and int to MD5Sum.
func Md5Sum(str string, num int) string {
	var combined string
	fmt.Sscanf(combined, "%s%d", str, num)
	bytes := []byte(combined)
	hex := md5.Sum(bytes)
	return ToString(hex[:])
}
