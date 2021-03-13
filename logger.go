package logger

import "io"

const (
	GoSimpleDateTimeZone = "2006-01-02 15:04:05Z0700"
	GoSimpleDateTime     = "2006-01-02 15:04:05"
	GoSimpleDate         = "2006-01-02"
	GoSimpleTime         = "15:04:05"

	// FileSep et al are ASCII control chars for data
	FileSep   = "\034" // byte(28); \x1c – FS – File separator
	GroupSep  = "\035" // byte(29); \x1d – GS – Group separator
	RecordSep = "\036" // byte(30); \x1e – RS – Record separator
	UnitSep   = "\037" // byte(31); \x1f – US – Unit separator

)

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
