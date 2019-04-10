package logger

import (
	"io"
	"log"
)

// Error, Warning, Info, and Debug define prefixes for logged output that signal a level of severity.
const (
	Error   = "! error; "   // errors are when both the operation and the application can't proceed
	Warning = "# warning; " // warnings (alerts) are when the operation cannot proceed but the application can
	Debug   = ": debug; "   // debug is used for debugging messages that will most likely be silenced/removed in production
	Info    = ". info; "    // info is used for informational purposes like status messages for long running processed
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
