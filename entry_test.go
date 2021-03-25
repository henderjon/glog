package logger

import (
	"encoding/json"
	"testing"
	"time"

	"github.com/google/go-cmp/cmp"
)

func TestString(t *testing.T) {
	expected := "2020-02-10 13:51:12Z\tentry_test.go:18\tl5\tThis is a test"

	ts, _ := time.Parse(GoSimpleDateTimeZone, "2020-02-10 13:51:12Z")

	actual := &Entry{
		Timestamp: ts,
		Location:  Here(),
		Message:   "This is a test",
		Level:     Level(5),
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

	ts, _ := time.Parse(GoSimpleDateTimeZone, "2020-02-10 13:51:12Z")
	actual.Timestamp = ts

	expected := "2020-02-10 13:51:12Z\tentry_test.go:29\tThis is a test\t[{\"Fizz\":\"Buzz\"}]"
	if diff := cmp.Diff(actual.String(), expected); diff != "" {
		t.Error("TestNew; (-got +want)", diff)
	}
}

func TestJSONMarshal(t *testing.T) {
	// expected := `{"Message":"This is a test","Location":"main.go:30","Level":51}`
	expected := `{"Message":"This is a test","Level":51,"Context":[{"Fizz":"Buzz"}]}`

	s, e := json.Marshal(&Entry{
		Message: "This is a test",
		Level:   Level(51),
		Context: []interface{}{
			struct {
				Fizz string
			}{
				Fizz: "Buzz",
			},
		},
	})

	if e != nil {
		t.Error("TestJSONMarshal;", e)
	}

	if diff := cmp.Diff(string(s), expected); diff != "" {
		t.Error("TestJSONMarshal; (-got +want)", diff)
	}
}
func TestJSONUnmarshal(t *testing.T) {
	// expected := `{"Message":"This is a test","Location":"main.go:30","Level":51}`
	expected := &Entry{
		Message: "This is a test",
		Level:   Level(51),
		// Timestamp: time.Now().UTC(),
		Context: []interface{}{ // interface{} and JSONUnmarshal don't play nice
			map[string]interface{}{"Fizz": string("Buzz")},
		},
	}

	var actual Entry
	e := json.Unmarshal([]byte(`{"Message":"This is a test","Level":51,"Context":[{"Fizz":"Buzz"}]}`), &actual)

	if e != nil {
		t.Error("TestJSONUnmarshal;", e)
	}

	if diff := cmp.Diff(&actual, expected); diff != "" {
		t.Error("TestJSONMarshal; (-got +want)", diff)
	}
}

func TestWrite(t *testing.T) {
	expected := `This is a test`

	actual := &Entry{}

	i, _ := actual.Write([]byte(`This is a test`))

	if i != 14 {
		t.Error("TestWrite; (-got +want)", i, 14)
	}

	if diff := cmp.Diff(actual.String(), expected); diff != "" {
		t.Error("TestString; (-got +want)", diff)
	}
}

func TestCSV(t *testing.T) {
	var (
		actual   string
		expected string
	)

	e := &Entry{
		Message: "This is a test",
		Level:   Level(51),
		// Timestamp: time.Now().UTC(),
		Context: []interface{}{
			struct {
				Fizz string
			}{
				Fizz: "Buzz",
			},
		},
	}

	expected = "level,l51,message,This is a test,context,W3siRml6eiI6IkJ1enoifV0\n"
	actual, _ = e.MarshalCSV(true)

	if diff := cmp.Diff(actual, expected); diff != "" {
		t.Error("MarshalCSV; (-got +want)", diff)
	}

	expected = "l51,This is a test,W3siRml6eiI6IkJ1enoifV0\n"
	actual, _ = e.MarshalCSV(false)

	if diff := cmp.Diff(actual, expected); diff != "" {
		t.Error("MarshalCSV; (-got +want)", diff)
	}
}

func TestLV(t *testing.T) {
	var (
		actual   string
		expected string
	)

	e := &Entry{
		Message: "This is a test",
		Level:   Level(51),
		// Timestamp: time.Now().UTC(),
		Context: []interface{}{
			struct {
				Fizz string
			}{
				Fizz: "Buzz",
			},
		},
	}

	expected = `6:5:level;3:l51;7:message;14:This is a test;7:context;23:W3siRml6eiI6IkJ1enoifV0;`
	actual, _ = e.MarshalLV(true)

	if diff := cmp.Diff(actual, expected); diff != "" {
		t.Error("MarshalLV; (-got +want)", diff)
	}

	expected = `3:3:l51;14:This is a test;23:W3siRml6eiI6IkJ1enoifV0;`
	actual, _ = e.MarshalLV(false)

	if diff := cmp.Diff(actual, expected); diff != "" {
		t.Error("MarshalLV; (-got +want)", diff)
	}
}
