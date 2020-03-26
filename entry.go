package logger

import (
	"bytes"
	"encoding/binary"
	"encoding/json"
	"errors"
	"io"
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
	Timestamp time.Time     `json:",omitempty"` // time.Time
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
		e.append(arg)
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
			e.Location = here(3)
		}
	case *Entry:
		return val // we're not allowing wrapping at this time
	default:
		e.AppendContext(val)
	}
	return e
}

func (e *Entry) String() string {
	s, _ := e.MarshalPlain()
	return s
}

// MarshalPlain is the plain text representation of an Entry
func (e *Entry) MarshalPlain() (string, error) {
	var (
		str bytes.Buffer
		ctx []byte
		err error
	)

	if !e.Timestamp.IsZero() {
		str.WriteString(e.Timestamp.Format(GoMySQLDateTime))
		str.WriteString(TabSep)
	}

	if e.Location != "" {
		str.WriteString(string(e.Location))
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
	return string(bytes.TrimRight(str.Bytes(), TabSep)), nil
}

// MarshalBin is the byte/binary representation of an Entry
func (e *Entry) MarshalBin() ([]byte, error) {
	var marshaledBin bytes.Buffer

	if e == nil {
		return marshaledBin.Bytes(), errors.New("empty entry")
	}

	appendString(&marshaledBin, e.Timestamp.Format(time.RFC3339))
	appendString(&marshaledBin, string(e.Location))
	appendString(&marshaledBin, e.Message)

	var (
		ctx []byte
		err error
	)

	if e.Context != nil {
		ctx, err = json.Marshal(e.Context)
		if err != nil {
			ctx = nil
		}
	}

	appendString(&marshaledBin, string(ctx))
	return marshaledBin.Bytes(), nil
}

// UnmarshalBin is the reverse of MarshalBin and populates an Entry from the byte/binary representation
func (e *Entry) UnmarshalBin(b []byte) error {
	if len(b) == 0 {
		return nil
	}

	var err error
	buf := bytes.NewBuffer(b)
	data := getBytes(buf)
	if data != nil {
		e.Timestamp, err = time.Parse(time.RFC3339, string(data))
		if err != nil {
			return err
		}
	}

	data = getBytes(buf)
	if data != nil {
		e.Location = Location(data) // @TODO double casting?
	}

	data = getBytes(buf)
	if data != nil {
		e.Message = string(data)
	}

	data = getBytes(buf)
	err = json.Unmarshal(data, &e.Context)
	if err != nil {
		return err
	}

	return nil
}

func appendString(w io.Writer, str string) uint64 {
	l := uint64(len(str))
	binary.Write(w, binary.BigEndian, l)
	binary.Write(w, binary.BigEndian, []byte(str))
	return l
}

func getBytes(r io.Reader) []byte {
	var tmpL uint64
	binary.Read(r, binary.BigEndian, &tmpL)
	tmpB := make([]byte, int(tmpL))
	binary.Read(r, binary.BigEndian, &tmpB)
	return tmpB
}
