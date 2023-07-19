package logger

import (
	"io"
	"log"
	"os"
)

// A StdLogger is logger that is tied to either STDOUT ot STDIN
type StdLogger struct {
	Postmark bool
	log      *log.Logger
}

// NewStdLogger returns a logger that writes to w
func NewStdLogger(w io.Writer) *StdLogger {
	return &StdLogger{
		false,
		log.New(w, "", 0),
	}
}

// Log fulfills the Logger interface. It writes the entry to the underlying destination
func (l StdLogger) Log(args ...interface{}) {
	args = append(args, l.Postmark)
	e := entry(args...)
	l.log.Println(e)
}

// Fatal fulfills the Logger interface. It writes the entry to the underlying destination then exits
func (l StdLogger) Fatal(args ...interface{}) {
	args = append(args, l.Postmark)
	e := entry(args...)
	l.log.Fatalln(e)
}

// Write fulfills the io.Writer interface
func (l StdLogger) Write(p []byte) (n int, err error) {
	e := entry(p)
	e.append(l.Postmark)
	l.log.Println(e)
	return len(p), nil
}

// NewStderrLogger creates a new logger that writes to STDERR or /dev/null.
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

// NewStdoutLogger creates a new logger that writes to STDOUT or /dev/null.
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
