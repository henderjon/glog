package logger

import (
	"io"
	"log"
)

type RowLog struct {
	log  *log.Logger
	opts Opts
}

func NewRowLogger(w io.Writer, opts ...SetOpt) Logger {
	o := DefaultOpts
	for _, opt := range opts {
		o = opt(o)
	}

	return &RowLog{
		log:  log.New(w, "", 0),
		opts: o,
	}
}

// Log fulfills the Logger interface. It writes the entry to the underlying destination
func (l RowLog) Log(args ...interface{}) {
	l.log.Print(rowLogFormat(3, l.opts, args), l.opts.RecordSep)
}

// Fatal fulfills the Logger interface. It writes the entry to the underlying destination then exits
func (l RowLog) Fatal(args ...interface{}) {
	l.log.Fatal(rowLogFormat(3, l.opts, args), l.opts.RecordSep)
}

// Write fulfills the io.Writer interface
func (l RowLog) Write(p []byte) (n int, err error) {
	l.log.Print(rowLogFormat(4, l.opts, []any{string(p)}), l.opts.RecordSep)
	return len(p), nil
}
