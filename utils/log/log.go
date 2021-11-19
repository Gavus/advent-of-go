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

func Init() {
	Info = log.New(os.Stderr, "INFO: ", 0)
	Warn = log.New(os.Stderr, "WARNING: ", 0)
	Err = log.New(os.Stderr, "ERROR", 0)
}
