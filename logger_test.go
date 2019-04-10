package logger

import (
	"log"
	"os"
	"testing"
)

func TestLogger(t *testing.T) {
	var l *log.Logger

	l = New(os.Stderr, Error)
	l.Println("Hello World")

	l = New(os.Stderr, Warning)
	l.Println("Hello World")

	l = New(os.Stderr, Debug)
	l.Println("Hello World")

	l = New(os.Stderr, Info)
	l.Println("Hello World")

	l = New(os.Stderr, None)
	l.Println("Hello World")
}
