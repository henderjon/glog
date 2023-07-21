package logger

import (
	"io"
	"log"
	"os"
)

// A DropLogger is logger that is tied to either STDOUT ot STDIN
type DropLogger struct {
	log  *log.Logger
	opts Opts
}

// NewDropLogger returns a logger that writes to w
func NewDropLogger(w io.Writer, opts ...SetOpt) *DropLogger {
	o := DefaultOpts
	for _, opt := range opts {
		o = opt(o)
	}

	return &DropLogger{
		log:  log.New(w, "", 0),
		opts: o,
	}
}

// Log fulfills the Logger interface. It writes the entry to the underlying destination
func (l DropLogger) Log(args ...any) {
	dropLog(3, l.opts, args)
}

// Fatal fulfills the Logger interface. It writes the entry to the underlying destination then exits
func (l DropLogger) Fatal(args ...any) {
	dropLog(3, l.opts, args)
	os.Exit(1)
}

// Write fulfills the io.Writer interface
func (l DropLogger) Write(p []byte) (n int, err error) {
	dropLog(4, l.opts, []any{string(p)})
	return len(p), nil
}
