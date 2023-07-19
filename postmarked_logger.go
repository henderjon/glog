package logger

// A PostmarkedLogger is a regular instance of Logger but will Postmark all entries.
type PostmarkedLogger struct {
	Logger
}

// NewPostmarkedLogger creates a new logger that assumes Postmark is true for all entries
func NewPostmarkedLogger(l Logger) *PostmarkedLogger {
	return &PostmarkedLogger{
		l,
	}
}

// Log fulfills the Logger interface. It writes the entry to the underlying destination
func (l PostmarkedLogger) Log(args ...interface{}) {
	args = append(args, Postmark)
	e := entry(args...)
	l.Log(e)
}

// Fatal fulfills the Logger interface. It writes the entry to the underlying destination then exits
func (l PostmarkedLogger) Fatal(args ...interface{}) {
	args = append(args, Postmark)
	e := entry(args...)
	l.Fatal(e)
}

// Write fulfills the io.Writer interface
func (l PostmarkedLogger) Write(p []byte) (n int, err error) {
	e := entry(p, Postmark)
	return l.Write([]byte(e.String()))
}
