package logger

import (
	"bytes"
	"encoding/base64"
	"encoding/csv"
	"encoding/json"
	"strconv"
	"strings"
	"time"
)

const (
	// Postmark is a readable way of adding both the current Timestampa nd Location to an Entry
	Postmark = true
	// TabSep is the seperator used when using MarshalPlain
	TabSep = "\t"
)

// Entry is a log entry
type Entry struct {
	Message   string        `json:",omitempty"` // string message
	Location  Location      `json:",omitempty"` // path/file.ext:line
	Timestamp time.Time     `json:",omitempty"` // time.Time; omit doesn't work on empty time.Time but does on an empty *time.Time
	Level     Level         `json:",omitempty"` // uint8
	Context   []interface{} `json:",omitempty"` // additional structured information to be JSON serialized
}

// NewEntry create a new Entry
func NewEntry(msg string) *Entry {
	return &Entry{
		Timestamp: time.Now().UTC(),
		Location:  here(2),
		Message:   msg,
	}
}

// see interface docs
func entry(args ...interface{}) *Entry {
	e := &Entry{}
	for _, arg := range args {
		val, ok := arg.(*Entry)
		if ok {
			e.append(val.String())
		} else {
			e.append(arg)
		}
	}
	return e
}

// Write fulfills the io.Writer interface
func (e *Entry) Write(p []byte) (n int, err error) {
	e.setMessage(string(p))
	return len(p), nil
}

// Append is a func to add items to an Entry's Context
func (e *Entry) Append(arg interface{}) *Entry {
	return e.append(arg)
}

// AppendContext is a func to add items to an Entry's Context
func (e *Entry) AppendContext(arg interface{}) *Entry {
	e.Context = append(e.Context, arg)
	return e
}

func (e *Entry) append(arg interface{}) *Entry {
	switch val := arg.(type) {
	case []byte:
		e.setMessage(string(val))
	case string:
		e.setMessage(val)
	case error:
		e.setMessage(val.Error())
	case time.Time:
		e.setTimestamp(val)
	case Location:
		e.setLocation(val)
	case bool:
		if val {
			e.setTimestamp(time.Now().UTC())
			e.setLocation(here(4))
		}
	case Level:
		e.Level = val
	default:
		e.AppendContext(val)
	}
	return e
}

// setTimestamp will check and add the provided the Timestamp
func (e *Entry) setTimestamp(t time.Time) {
	if t.IsZero() {
		t = time.Now().UTC()
	}

	if e.Timestamp.IsZero() {
		e.Timestamp = t
	} else {
		e.AppendContext(t)
	}
}

// setTimestamp will check and add the provided the Location
func (e *Entry) setLocation(l Location) {
	if l == "" {
		l = here(4)
	}

	if e.Location == "" {
		e.Location = l
	} else {
		e.AppendContext(l)
	}
}

// setTimestamp will check and add the provided the Message
func (e *Entry) setMessage(m string) {
	if len(m) < 1 {
		return
	}

	if e.Message == "" {
		e.Message = m
	} else {
		e.AppendContext(m)
	}
}

// String satisfies the fmt.Stringer interface
func (e *Entry) String() string {
	s, _ := e.MarshalText()
	return string(s)
}

// MarshalText implements the encoding.TextMarshaler interface.
func (e *Entry) MarshalText() ([]byte, error) {
	var (
		str bytes.Buffer
		ctx []byte
		err error
	)

	if !e.Timestamp.IsZero() {
		str.WriteString(e.Timestamp.Format(GoSimpleDateTimeZone))
		str.WriteString(TabSep)
	}

	if e.Location != "" {
		str.WriteString(string(e.Location))
		str.WriteString(TabSep)
	}

	if e.Level != 0 {
		str.WriteString(e.Level.String())
		str.WriteString(TabSep)
	}

	str.WriteString(e.Message)
	str.WriteString(TabSep)

	if e.Context != nil {
		ctx, err = json.Marshal(e.Context)
		if err != nil {
			ctx = nil
		}
	}

	str.WriteString(string(ctx))
	return bytes.TrimRight(str.Bytes(), TabSep), nil
}

// MarshalJSON implements the json.Marshaler interface and let's me fix time.Time over JSON
// Using an alias with an embedded Entry let's us control the time.Time un/marshaling
// choly.ca/post/go-json-marshalling/
func (e Entry) MarshalJSON() ([]byte, error) {
	type tmp Entry
	e2 := &struct {
		Timestamp *time.Time `json:",omitempty"`
		*tmp
	}{
		tmp: (*tmp)(&e),
	}

	if !e.Timestamp.IsZero() {
		e2.Timestamp = &e.Timestamp
	}

	return json.Marshal(e2)
}

// MarshalJSON implements the json.Unmarshaler interface.
func (e *Entry) UnmarshalJSON(data []byte) error {
	type tmp Entry
	e2 := &struct {
		Timestamp *time.Time `json:",omitempty"`
		*tmp
	}{
		// cast our actual Entry to our alias and embed it so that the empty
		tmp: (*tmp)(e),
	}

	if err := json.Unmarshal(data, &e2); err != nil {
		return err
	}

	if e2.Timestamp != nil && !e2.Timestamp.IsZero() {
		e.Timestamp = time.Time(*e2.Timestamp)
	}

	return nil
}

// MarshalFlat creates an ordered []string out of the entry for use in CSVs and
// other text/log storage. Optionally, the keys can be interpolated to allow for
// breaking part large files and the json encoded context can be base64 encoded
// to provide for easier CSV parsing.
func (e *Entry) MarshalFlat(keys bool, b64 bool) []string {
	var (
		ctx []byte
		err error
	)

	record := make([]string, 0)

	if !e.Timestamp.IsZero() {
		if keys {
			record = append(record, `timestamp`)
		}
		record = append(record, e.Timestamp.Format(time.RFC3339))
	}

	if e.Location != "" {
		if keys {
			record = append(record, `location`)
		}
		record = append(record, string(e.Location))
	}

	if e.Level != 0 {
		if keys {
			record = append(record, `level`)
		}
		record = append(record, e.Level.String())
	}

	if keys {
		record = append(record, `message`)
	}
	record = append(record, e.Message)

	if e.Context != nil {
		ctx, err = json.Marshal(e.Context)
		if err == nil {
			if keys {
				record = append(record, `context`)
			}

			ctxStr := string(ctx)
			if b64 {
				ctxStr = base64.RawStdEncoding.EncodeToString(ctx)
			}

			record = append(record, ctxStr)
		}
	}

	return record
}

// MarshalCSV encodes the entry as a CSV
func (e *Entry) MarshalCSV(keys bool) (string, error) {
	var (
		str bytes.Buffer
	)

	c := csv.NewWriter(&str)
	// c.Comma = rune(31)

	record := e.MarshalFlat(keys, true)
	c.Write(record)
	c.Flush()

	return str.String(), nil
}

// MarshalLV encodes the entry as a `n:length:value;[length:value;...]` tuple
func (e *Entry) MarshalLV(keys bool) (string, error) {
	record := e.MarshalFlat(keys, false)
	return marshalRecord(record), nil
}

func marshalUnit(v string) string {
	var s strings.Builder
	l := len(v)
	s.WriteString(strconv.Itoa(l))
	s.WriteString(`:`)
	s.WriteString(v)
	s.WriteString(`;`)
	return s.String()
}

func marshalRecord(vs []string) string {
	var s strings.Builder
	l := len(vs)
	if l < 1 {
		return ""
	}

	s.WriteString(strconv.Itoa(l))
	s.WriteString(`:`)
	for i := range vs {
		s.WriteString(marshalUnit(vs[i]))
	}
	return s.String()
}
