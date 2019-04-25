package logger

import (
	"log"
	"testing"
)

func TestLogger(t *testing.T) {
	var l *log.Logger

	l = NewStderrLogger(true, Error.String(), 0)
	l.Println("Hello World")

	l = NewStdoutLogger(true, Warning.String(), 0)
	l.Println("Hello World")

	l = NewStderrLogger(true, Info.String(), 0)
	l.Println("Hello World")

	l = NewStdoutLogger(true, Debug.String(), 0)
	l.Println("Hello World")
}
