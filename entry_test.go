package logger

import (
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestString(t *testing.T) {
	expected := "2020-02-10 13:51:12\tentry_test.go:17\tThis is a test\t"

	ts, _ := time.Parse(GoMySQLDateTime, "2020-02-10 13:51:12")

	actual := &Entry{
		Timestamp: ts,
		Location:  Here(),
		Message:   "This is a test",
	}

	if diff := cmp.Diff(actual.String(), expected); diff != "" {
		t.Error("TestString; (-got +want)", diff)
	}
}

func TestNew(t *testing.T) {
	actual := NewEntry("This is a test")
	actual.AppendContext(struct {
		Fizz string
	}{
		Fizz: "Buzz",
	})

	ts, _ := time.Parse(GoMySQLDateTime, "2020-02-10 13:51:12")
	actual.Timestamp = ts

	expected := "2020-02-10 13:51:12\tentry_test.go:27\tThis is a test\t[{\"Fizz\":\"Buzz\"}]"
	if diff := cmp.Diff(actual.String(), expected); diff != "" {
		t.Error("TestNew; (-got +want)", diff)
	}
}

func TestMarshalBin(t *testing.T) {
	actual := entry("This is a test", Here(), struct {
		Fizz string
	}{
		Fizz: "Buzz",
	})

	ts, _ := time.Parse(GoMySQLDateTime, "2020-02-10 13:51:12")
	actual.Timestamp = ts

	expected := []byte{
		20, 50, 48, 50, 48, 45, 48, 50, 45, 49, 48, 84, 49, 51, 58, 53, 49, 58, 49, 50, 90, 16, 101, 110, 116, 114, 121, 95, 116, 101, 115, 116, 46, 103, 111, 58, 52, 52, 14, 84, 104, 105, 115, 32, 105, 115, 32, 97, 32, 116, 101, 115, 116, 17, 91, 123, 34, 70, 105, 122, 122, 34, 58, 34, 66, 117, 122, 122, 34, 125, 93,
		// 20, 50, 48, 50, 48, 45, 48, 50, 45, 49, 48, 84, 49, 51, 58, 53, 49,
		// 58, 49, 50, 90, 16, 101, 110, 116, 114, 121, 95, 116, 101, 115, 116,
		// 46, 103, 111, 58, 52, 52, 14, 84, 104, 105, 115, 32, 105, 115, 32,
		// 97, 32, 116, 101, 115, 116, 15, 123, 34, 70, 105, 122, 122, 34, 58,
		// 34, 66, 117, 122, 122, 34, 125,
	}

	marshaled, _ := actual.MarshalBin()
	if diff := cmp.Diff(marshaled, expected); diff != "" {
		t.Error("TestMarshalBin; (-got +want)", diff)
	}
}

func TestUnmarshalBin(t *testing.T) {
	actual := []byte{
		20, 50, 48, 50, 48, 45, 48, 50, 45, 49, 48, 84, 49, 51, 58, 53, 49, 58, 49, 50, 90, 16, 101, 110, 116, 114, 121, 95, 116, 101, 115, 116, 46, 103, 111, 58, 52, 52, 14, 84, 104, 105, 115, 32, 105, 115, 32, 97, 32, 116, 101, 115, 116, 17, 91, 123, 34, 70, 105, 122, 122, 34, 58, 34, 66, 117, 122, 122, 34, 125, 93,
		// 20, 50, 48, 50, 48, 45, 48, 50, 45, 49, 48, 84, 49, 51, 58, 53, 49,
		// 58, 49, 50, 90, 16, 101, 110, 116, 114, 121, 95, 116, 101, 115, 116,
		// 46, 103, 111, 58, 52, 51, 14, 84, 104, 105, 115, 32, 105, 115, 32,
		// 97, 32, 116, 101, 115, 116, 15, 123, 34, 70, 105, 122, 122, 34, 58,
		// 34, 66, 117, 122, 122, 34, 125,
	}

	ts, _ := time.Parse(GoMySQLDateTime, "2020-02-10 13:51:12")

	type fb struct {
		Fizz string
	}

	expected := &Entry{
		Timestamp: ts,
		Location:  Location("entry_test.go:44"),
		Message:   "This is a test",
		Context: []interface{}{
			&fb{
				Fizz: "Buzz",
			},
		},
	}

	e := Entry{
		Context: []interface{}{&fb{}},
	}

	err := e.UnmarshalBin(actual)
	if err != nil {
		t.Error(err)
	}

	if diff := cmp.Diff(e, *expected); diff != "" {
		t.Error("TestUnmarshalBin; (-got +want)", diff)
	}
}
