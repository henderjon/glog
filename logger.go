package logger

import "io"

// Logger is a simple interface for logging a message and possibly exiting your program.
//
// Log creates an entry from an unspecified group of params. Note that
// the first string or error will be considered the Message and all
// others will be appended to the Context. The same is true for the
// first Location and the first time.Time as they will be assigned to
// the correct field and subsequent values of that type will be appended
// to the context. A bool(true) will automatically populate the
// Timestamp and Location of the created entry. Fatal works the same way
// but will call os.Exit(1) when it's done. Passing an *Entry will
// simply return that *Entry as opposed to wrapping it in the Context.
type Logger interface {
	Log(args ...interface{})
	Fatal(args ...interface{})
	io.Writer
}
