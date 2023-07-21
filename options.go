package logger

const (
	GoSimpleDateTimeZone = "2006-01-02 15:04:05Z0700"
	GoSimpleDateTime     = "2006-01-02 15:04:05"
	GoSimpleDate         = "2006-01-02"
	GoSimpleTime         = "15:04:05"

	// FileSep et al are ASCII control chars for data
	FS = "\034" // byte(28); \x1c – FS – File separator
	GS = "\035" // byte(29); \x1d – GS – Group separator
	RS = "\036" // byte(30); \x1e – RS – Record separator
	US = "\037" // byte(31); \x1f – US – Unit separator

	// Postmark is a readable way of adding both the current Timestamp and Location to an Entry
	Postmark = true

	TabSep          = "\t"
	LineSep         = "\n"
	RecordPrefix    = `#`
	UnitSeperator   = `:`
	RecordSeperator = `;`
)

type Opts struct {
	UnitSep   string
	RecordSep string
	Prefix    string
	Timestamp bool
	Location  bool
}

var DefaultOpts = Opts{
	UnitSep:   TabSep,
	RecordSep: LineSep,
	Prefix:    RecordPrefix,
	Timestamp: false,
	Location:  true,
}

type SetOpt func(o Opts) Opts

func SetUnitSep(s string) SetOpt {
	return SetOpt(func(o Opts) Opts {
		o.UnitSep = s
		return o
	})
}

func SetRecordSep(s string) SetOpt {
	return SetOpt(func(o Opts) Opts {
		o.RecordSep = s
		return o
	})
}

func SetPrefix(s string) SetOpt {
	return SetOpt(func(o Opts) Opts {
		o.Prefix = s
		return o
	})
}

func LogTimestamp(b bool) SetOpt {
	return SetOpt(func(o Opts) Opts {
		o.Timestamp = b
		return o
	})
}

func LogLocation(b bool) SetOpt {
	return SetOpt(func(o Opts) Opts {
		o.Location = b
		return o
	})
}
