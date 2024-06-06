package logger

import (
	"log"
	"os"
)

// This logger handles all data we got from the clients
// It should be sanitized before printing to stdout to avoid
// messing with terminals control characters
var DataLogger *log.Logger

// This logger is used to explain to the user what goyessir is doing
// (saving request, saving files...)
var InfoLogger *log.Logger

// This logger is used whenever an error occurs
var ErrorLogger *log.Logger

func InitLoggers() {
	DataLogger = log.New(os.Stdout, "", 0)

	InfoLogger = log.New(os.Stderr, "INFO ", log.LstdFlags|log.Lmsgprefix)

	ErrorLogger = log.New(os.Stderr, "ERROR ", log.Ldate|log.Ltime|log.Lmicroseconds|log.Llongfile|log.Lmsgprefix)
}
