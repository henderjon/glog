package logger

import (
	"bytes"
	"encoding/json"
	"time"
)

const (
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

func (e *Entry) Write(p []byte) (n int, err error) {
	e.Message = string(p)
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
		if e.Message == "" {
			e.Message = string(val)
		} else {
			e.AppendContext(val)
		}
	case string:
		if e.Message == "" {
			e.Message = val
		} else {
			e.AppendContext(val)
		}
	case error:
		if e.Message == "" {
			e.Message = val.Error()
		} else {
			e.AppendContext(val)
		}
	case time.Time:
		if e.Timestamp.IsZero() {
			e.Timestamp = val
		} else {
			e.AppendContext(val)
		}
	case Location:
		if e.Location == "" {
			e.Location = val
		} else {
			e.AppendContext(val)
		}
	case bool:
		if val == true {
			e.Timestamp = time.Now().UTC()
			e.Location = here(4)
		}
	case Level:
		e.Level = val
	default:
		e.AppendContext(val)
	}
	return e
}

func (e *Entry) String() string {
	s, _ := e.MarshalText()
	return string(s)
}

// MarshalText is the plain text representation of an Entry
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

// MarshalJSON satisfies the Marshal interface and let's me fix time.Time over JSON
// Using an alias with an embedded Entry let's us control the time.Time un/marshaling
// choly.ca/post/go-json-marshalling/
func (e *Entry) MarshalJSON() ([]byte, error) {
	type tmp Entry
	e2 := &struct {
		Timestamp *time.Time `json:",omitempty"`
		*tmp
	}{
		tmp: (*tmp)(e),
	}

	if !e.Timestamp.IsZero() {
		e2.Timestamp = &e.Timestamp
	}

	return json.Marshal(e2)
}

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

	return nil
}
