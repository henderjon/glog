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

// Write fulfills the io.Writer interface
func (l StdLog) Write(p []byte) (n int, err error) {
	e := entry(p)
	l.log.Println(e)
	return len(p), nil
}

// NewStderrLogger creates a new debuglogger that can write to STDERR.
// By default this logger logs to `ioutil.Discard` which is an alias for
// /dev/null. Passing `true` to this constructor causes the output to go
// to stderr. This behavior allows the Log and Fatal invocations to be
// silenced and therefore left in place.
func NewStderrLogger(stderr bool) *StdLog {
	l := log.New(ioutil.Discard, "null ", log.Lshortfile|log.LUTC|log.LstdFlags)
	if stderr {
		l = log.New(os.Stderr, "", 0)
	}
	return &StdLog{l}
}

// NewStdoutLogger creates a new debuglogger that can write to STDOUT.
// By default this logger logs to `ioutil.Discard` which is an alias for
// /dev/null. Passing `true` to this constructor causes the output to go
// to stdout. This behavior allows the Log and Fatal invocations to be
// silenced and therefore left in place.
func NewStdoutLogger(stdout bool) *StdLog {
	l := log.New(ioutil.Discard, "null ", log.Lshortfile|log.LUTC|log.LstdFlags)
	if stdout {
		l = log.New(os.Stdout, "", 0)
	}
	return &StdLog{l}
}
