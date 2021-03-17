package logger

import (
	"io"
	"log"
	"os"
)

type StdLogger struct {
	log *log.Logger
}

func NewStdLogger(w io.Writer) *StdLogger {
	return &StdLogger{
		log.New(w, "", 0),
	}
}

// Log fulfills the Logger interface. It writes the entry to the underlying destination
func (l StdLogger) Log(args ...interface{}) {
	e := entry(args...)
	l.log.Println(e)
}

// Fatal fulfills the Logger interface. It writes the entry to the underlying destination then exits
func (l StdLogger) Fatal(args ...interface{}) {
	e := entry(args...)
	l.log.Fatalln(e)
}

// Write fulfills the io.Writer interface
func (l StdLogger) Write(p []byte) (n int, err error) {
	e := entry(p)
	l.log.Println(e)
	return len(p), nil
}

// NewStderrLogger creates a new debuglogger that can write to STDERR.
// By default this logger logs to `ioutil.Discard` which is an alias for
// /dev/null. Passing `true` to this constructor causes the output to go
// to stderr. This behavior allows the Log and Fatal invocations to be
// silenced and therefore left in place.
func NewStderrLogger(stderr bool) *StdLogger {
	if stderr {
		return NewStdLogger(os.Stderr)
	}
	return NewStdLogger(io.Discard)
}

// NewStdoutLogger creates a new debuglogger that can write to STDOUT.
// By default this logger logs to `io.Discard` which is an alias for
// /dev/null. Passing `true` to this constructor causes the output to go
// to stdout. This behavior allows the Log and Fatal invocations to be
// silenced and therefore left in place.
func NewStdoutLogger(stdout bool) *StdLogger {
	if stdout {
		return NewStdLogger(os.Stdout)
	}
	return NewStdLogger(io.Discard)
}
