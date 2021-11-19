// Converting functionalities.
package conv

import (
	"fmt"
	"github.com/Gavus/advent-of-code/utils/log"
	"github.com/Gavus/advent-of-code/utils/types"
)

// Convert strings into Presents.
func ToPresents(input []string) []types.Present {
	presents := make([]types.Present, len(input))

	for index, value := range input {
		var l, w, h int
		_, err := fmt.Sscanf(value, "%dx%dx%d", &l, &w, &h)
		if err != nil {
			log.Err.Fatal(err)
		}
		presents[index].Set(l, w, h)
	}

	return presents
}
