package logger

import (
	"log"
	"testing"
)

func TestLogger(t *testing.T) {
	var l *log.Logger

	l = NewStderrLogger(false, Info, 0)
	l.Println("Hello World")

	l = NewStdoutLogger(false, Debug, 0)
	l.Println("Hello World")
}
