package logger

import "os"

type MultiLog struct {
	logs []Logger
}

// NewMultiLog creates a new MultiLog with the associated Loggers
func NewMultiLog(args ...Logger) *MultiLog {
	return &MultiLog{args}
}

// Log fulfills the Logger interface. It writes the entry to the underlying destination
func (l MultiLog) Log(args ...interface{}) {
	e := entry(args...)
	for _, log := range l.logs {
		log.Log(e)
	}
}

// Fatal fulfills the Logger interface. It writes the entry to the underlying destination then exits
func (l MultiLog) Fatal(args ...interface{}) {
	l.Log(args...)
	os.Exit(1)
}
