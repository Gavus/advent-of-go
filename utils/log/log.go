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

// Initializes the loggers.
func init() {
	Info = log.New(os.Stderr, "INFO: ", 0)
	Warn = log.New(os.Stderr, "WARNING: ", 0)
	Err = log.New(os.Stderr, "ERROR: ", log.Lshortfile)
}
