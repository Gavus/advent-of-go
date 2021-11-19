// Add new loggers which prepend the messages.
package log

import (
	"log"
	"os"
)

var (
	Info *log.Logger
	Warn *log.Logger
	Err  *log.Logger
)

// Initializes the loggers. Call this function once, preferably in the beginning of the main function.
func Init() {
	Info = log.New(os.Stderr, "INFO: ", 0)
	Warn = log.New(os.Stderr, "WARNING: ", 0)
	Err = log.New(os.Stderr, "ERROR: ", 0)
}
