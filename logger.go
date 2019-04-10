package logger

import (
	"io"
	"io/ioutil"
	"log"
	"os"
)

// Error, Warning, Info, and Debug define prefixes for logged output that signal a level of severity.
const (
	Error   = "! error: "   // errors are when both the operation and the application can't proceed
	Warning = "# warning: " // warnings (alerts) are when the operation cannot proceed but the application can
	Debug   = ": debug: "   // debug is used for debugging messages that will most likely be silenced/removed in production
	Info    = ". info: "    // info is used for informational purposes like status messages for long running processed
	None    = ""            // None is used when the information is intentional often intended for stdout
	devnull = 0             // ioutil.Discard
	stdout  = 1             // os.Stdout
	stderr  = 2             // os.Stderr
)

// New creates a new logger to the given writer with the given prefix. This
// constructor serves to wrap the flags that are commonly used.
func New(w io.Writer, prefix string) *log.Logger {
	return log.New(w, prefix, log.Lshortfile|log.LUTC|log.LstdFlags)
}

// errWriter decides to be loud or not on STDERR
func errWriter(verbose bool) io.Writer {
	var w io.Writer
	w = ioutil.Discard
	if verbose {
		w = os.Stderr
	}
	return w
}

// outWriter decides to be loud or not on STDOUT
func outWriter(verbose bool) io.Writer {
	var w io.Writer
	w = ioutil.Discard
	if verbose {
		w = os.Stdout
	}
	return w
}

// NewErrorLogger returns a new logger of level: Error. These errors are when both the operation and the application can't proceed.
func NewErrorLogger(verbose bool) *log.Logger {
	return New(errWriter(verbose), Error)
}

// NewWarningLogger returns a new logger of level: Warning. These warnings (alerts) are when the operation cannot proceed but the application can.
func NewWarningLogger(verbose bool) *log.Logger {
	return New(errWriter(verbose), Warning)
}

// NewDebugLogger returns a new logger of level: Debug. Debug is used for debugging messages that will most likely be silenced/removed in production.
func NewDebugLogger(verbose bool) *log.Logger {
	return New(errWriter(verbose), Debug)
}

// NewInfoLogger returns a new logger of level: Info. Info is used for informational purposes like status messages for long running processed.
func NewInfoLogger(verbose bool) *log.Logger {
	return New(errWriter(verbose), Info)
}

// NewStdoutLogger returns a new logger of level: Stdout. This is used when the information is intentional; often intended for stdout.
func NewStdoutLogger(verbose bool) *log.Logger {
	return log.New(outWriter(verbose), None, 0)
}
