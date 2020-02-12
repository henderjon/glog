package logger

import (
	"encoding/binary"
	"encoding/json"
	"errors"
	"log"
	"strings"
	"time"
)

const (
	// TabSep is the seperator used when using MarshalPlain
	TabSep = "\t"
)

// Entry is a log entry
type Entry struct {
	Message   string        `json:",omitempty"` // func with brief note
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

// AppendContext is a func to add items to an Entry's Context
func (e *Entry) AppendContext(arg interface{}) *Entry {
	e.Context = append(e.Context, arg)
	return e
}

// see interface docs
func entry(args ...interface{}) *Entry {
	e := &Entry{}
	for _, arg := range args {
		switch val := arg.(type) {
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
		str strings.Builder
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
	return str.String(), nil
}

// MarshalBin is the byte/binary representation of an Entry
func (e *Entry) MarshalBin() ([]byte, error) {
	var marshaledBin []byte

	if e == nil {
		return marshaledBin, errors.New("empty entry")
	}

	marshaledBin = appendString(marshaledBin, e.Timestamp.Format(time.RFC3339))
	marshaledBin = appendString(marshaledBin, string(e.Location))
	marshaledBin = appendString(marshaledBin, e.Message)

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

	marshaledBin = appendString(marshaledBin, string(ctx))
	return marshaledBin, nil
}

// UnmarshalBin is the reverse of MarshalBin and populates an Entry from the byte/binary representation
func (e *Entry) UnmarshalBin(b []byte) error {
	if len(b) == 0 {
		return nil
	}

	var err error

	data, b := getBytes(b)
	if data != nil {
		e.Timestamp, err = time.Parse(time.RFC3339, string(data))
		if err != nil {
			return err
		}
	}

	data, b = getBytes(b)
	if data != nil {
		e.Location = Location(data) // @TODO double casting?
	}

	data, b = getBytes(b)
	if data != nil {
		e.Message = string(data)
	}

	data, b = getBytes(b)
	err = json.Unmarshal(data, &e.Context)
	if err != nil {
		return err
	}

	return nil
}

func appendString(b []byte, str string) []byte {
	var tmp [16]byte // For use by PutUvarint.
	N := binary.PutUvarint(tmp[:], uint64(len(str)))
	b = append(b, tmp[:N]...)
	b = append(b, str...)
	return b
}

func getBytes(b []byte) (data, remaining []byte) {
	u, N := binary.Uvarint(b)
	if len(b) < N+int(u) {
		log.Printf("Unmarshal error: bad encoding")
		return nil, nil
	}
	if N == 0 {
		log.Printf("Unmarshal error: bad encoding")
		return nil, b
	}
	return b[N : N+int(u)], b[N+int(u):]
}
