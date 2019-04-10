package logger

import (
	"log"
	"testing"
)

func TestLogger(t *testing.T) {
	var l *log.Logger

	l = NewErrorLogger(false)
	l.Println("Hello World")

	l = NewWarningLogger(false)
	l.Println("Hello World")

	l = NewDebugLogger(false)
	l.Println("Hello World")

	l = NewInfoLogger(false)
	l.Println("Hello World")

	l = NewStdoutLogger(false)
	l.Println("Hello World")
}
