package logger

const (
	// GoMySQLIdxTimestamp et al are go format strings for mysql representations
	GoMySQLIdxTimestamp = "2006-01-02 15:04:00"
	GoMySQLDateTime     = "2006-01-02 15:04:05"
	GoMySQLDate         = "2006-01-02"
	GoMySQLTime         = "15:04:05"

	// PrefixFormatMinutes     = "200601021504"     // time.Now().UTC().Format(...) month-day-hour-min // Mon Jan 2 15:04:05 -0700 MST 2006
	// PrefixFormatHours       = "2006010215"       // time.Now().UTC().Format(...) month-day-hour     // Mon Jan 2 15:04:05 -0700 MST 2006
	// PrefixFormatDays        = "20060102"         // time.Now().UTC().Format(...) month-day          // Mon Jan 2 15:04:05 -0700 MST 2006
	// FilePrefixFormatMinutes = "2006/01/02/15/04" // time.Now().UTC().Format(...) month-day-hour-min // Mon Jan 2 15:04:05 -0700 MST 2006
	// FilePrefixFormatHours   = "2006/01/02/15"    // time.Now().UTC().Format(...) month-day-hour     // Mon Jan 2 15:04:05 -0700 MST 2006
	// FilePrefixFormatDays    = "2006/01/02/"      // time.Now().UTC().Format(...) month-day          // Mon Jan 2 15:04:05 -0700 MST 2006

	// MySQLFmtDateTime et al mysql datetime format strings
	MySQLFmtDateTime = "%Y-%m-%d %H:%i:%s"
	MySQLFmtMinute   = "%Y-%m-%d %H:%i"
	MySQLFmtHour     = "%Y-%m-%d %H"
	MySQLFmtDate     = "%Y-%m-%d"

	// FileSep et al are ASCII control chars for data
	FileSep   = "\034" // byte(28); \x1c – FS – File separator
	GroupSep  = "\035" // byte(29); \x1d – GS – Group separator
	RecordSep = "\036" // byte(30); \x1e – RS – Record separator
	UnitSep   = "\037" // byte(31); \x1f – US – Unit separator

)

// Logger is a simple interfac for logging a message and possibly exiting your program
type Logger interface {
	Log(args ...interface{})
	Fatal(args ...interface{})
}
