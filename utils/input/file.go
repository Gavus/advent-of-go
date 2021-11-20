// Functionalities related to file handling, such as reading or writing.
package input

import (
	"errors"
	"fmt"
	"github.com/Gavus/advent-of-code/utils/conv"
	"github.com/Gavus/advent-of-code/utils/date"
	"github.com/Gavus/advent-of-code/utils/log"
	"os"
)

const (
	inputPath = "./%s/day%s/input.txt"
	fileMode  = 0600
)

// Get Advent of code input from a certain date into strings. If the file isn't found will a download be attempted.
func GetInput(timeStr string, delimiter string) []string {
	time := date.ParseDate(timeStr)
	year, day := date.ToDayYearString(time)
	filepath := fmt.Sprintf(inputPath, year, day)

	if !fileExists(filepath) {
		log.Info.Print("Attempt to download input from website")
		bytes, err := DownloadInput(time.Year(), time.Day())
		if err != nil {
			log.Err.Fatal("Could not download input: ", err)
		}

		err = os.WriteFile(filepath, bytes, fileMode)
		if err != nil {
			log.Warn.Print(err)
		}
	}

	log.Info.Print("Attempt read input from file")
	input, err := getInputFromFile(filepath, delimiter)
	if err == nil {
		return input
	}

	return input
}

// Read input file and returns strings separated by delimiter.
func getInputFromFile(filepath string, delimiter string) ([]string, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		log.Warn.Print(err)
		return make([]string, 0), err
	}

	input := conv.ToStrings(data, delimiter)

	return input, nil
}

// Check if a file exists or not.
func fileExists(filepath string) bool {
	_, err := os.Stat(filepath)
	if errors.Is(err, os.ErrNotExist) {
		return false
	} else if err != nil {
		log.Err.Fatal(err)
	}

	return true
}
