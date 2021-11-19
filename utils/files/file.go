package file

import (
	"fmt"
	"github.com/Gavus/advent-of-code/utils/adventofcode"
	"github.com/Gavus/advent-of-code/utils/date"
	"github.com/Gavus/advent-of-code/utils/log"
	"os"
	"strings"
)

const (
	inputPath = "./%s/day%s/input.txt"
	fileMode  = 0600
)

func GetInput(timeStr string, delimiter string) []string {
	time := date.ParseDate(timeStr)
	year, day := date.TimeToDayYearString(time)
	filepath := fmt.Sprintf(inputPath, year, day)

	input, err := getInputFromFile(filepath, delimiter)
	if err == nil {
		return input
	}

	input, err = adventofcode.GetInput(year, day, delimiter)
	if err != nil {
		log.Err.Fatal(err)
	}

	err = SaveInputToFile(filepath, input, delimiter)
	if err != nil {
		log.Warn.Print(err)
	}

	return input
}

func getInputFromFile(filepath string, delimiter string) ([]string, error) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		log.Warn.Print(err)
		return make([]string, 0), err
	}

	dataStr := string(data)
	input := strings.Split(dataStr, delimiter)

	log.Info.Print("Read input from file successfully.")
	return input, nil
}

func SaveInputToFile(filepath string, input []string, delimiter string) error {
	dataStr := strings.Join(input, delimiter)
	data := []byte(dataStr)
	return os.WriteFile(filepath, data, fileMode)
}
