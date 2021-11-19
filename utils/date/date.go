// Functionality related to date and time.
package date

import (
	"github.com/Gavus/advent-of-code/utils/log"
	"strconv"
	"time"
)

const (
	layoutISO = "2006-01-02"
)

// Parse date from string in format "yyyy-mm-dd" to time.Time.
func ParseDate(date string) time.Time {
	t, err := time.Parse(layoutISO, date)
	if err != nil {
		log.Err.Fatal(err)
	}

	return t
}

// Convert time.Time into day and year as strings.
func TimeToDayYearString(time time.Time) (string, string) {
	year := strconv.Itoa(time.Year())
	day := strconv.Itoa(time.Day())
	if len(day) == 1 {
		return year, "0" + day
	}

	return year, day
}
