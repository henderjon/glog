package logger

import (
	"io"
	"io/ioutil"
	"log"
	"os"
	"time"
)

// Level is a typed string for representing the level of a log
type Level string

// Error, Warning, Info, and Debug define prefixes for logged output that signal a level of severity.
const (
	Error            Level = "! error; "                               // errors are when both the operation and the application can't proceed
	Warning          Level = "! warning; "                             // warnings (alerts) are when the operation cannot proceed but the application can
	Debug            Level = "< debug; "                               // debug is used for debugging messages that will most likely be silenced/removed in production
	Info             Level = "< info; "                                // info is used for informational purposes like status messages for long running processed
	None             Level = ""                                        // None is used when the information is intentional often intended for stdout
	devnull                = 0                                         // ioutil.Discard
	stdout                 = 1                                         // os.Stdout
	stderr                 = 2                                         // os.Stderr
	DateTimeFileLine       = log.Lshortfile | log.LUTC | log.LstdFlags // my usual flags set
)

func (l Level) String() string {
	return string(l)
}

// Now returns a
func Now() string {
	return time.Now().UTC().Format(time.RFC3339)
}

// NewStderrLogger returns a new logger that logs to stderr. The advantage of this sugar is that given a bool, this logger can be toggled on/off.
func NewStderrLogger(verbose bool, prefix string, flag int) *log.Logger {
	var w io.Writer
	w = ioutil.Discard
	if verbose {
		w = os.Stderr
	}
	return log.New(w, prefix, flag)
}

// NewStdoutLogger returns a new logger that logs to stdout. The advantage of this sugar is that given a bool, this logger can be toggled on/off.
func NewStdoutLogger(verbose bool, prefix string, flag int) *log.Logger {
	var w io.Writer
	w = ioutil.Discard
	if verbose {
		w = os.Stdout
	}
	return log.New(w, prefix, flag)
}
