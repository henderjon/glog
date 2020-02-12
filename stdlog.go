package logger

import (
	"io/ioutil"
	"log"
	"os"
)

type StdLog struct {
	log *log.Logger
}

// Log fulfills the Logger interface. It writes the entry to the underlying destination
func (l StdLog) Log(args ...interface{}) {
	e := entry(args...)
	l.log.Println(e)
}

// Fatal fulfills the Logger interface. It writes the entry to the underlying destination then exits
func (l StdLog) Fatal(args ...interface{}) {
	e := entry(args...)
	l.log.Fatalln(e)
}

// NewStderrLogger creates a new debuglogger that can write to STDERR
func NewStderrLogger(stderr bool) *StdLog {
	l := log.New(ioutil.Discard, "null ", log.Lshortfile|log.LUTC|log.LstdFlags)
	if stderr {
		l = log.New(os.Stderr, "", 0)
	}
	return &StdLog{l}
}

// NewStdoutLogger creates a new debuglogger that can write to STDOUT
func NewStdoutLogger(stdout bool) *StdLog {
	l := log.New(ioutil.Discard, "null ", log.Lshortfile|log.LUTC|log.LstdFlags)
	if stdout {
		l = log.New(os.Stdout, "", 0)
	}
	return &StdLog{l}
}
