package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/henderjon/logger"
)

func main() {
	out := logger.NewStdLogger(os.Stderr)
	pout := logger.NewPostmarkedLogger(out)

	out.Log("postmarked", logger.Postmark, logger.Level(8))
	out.Log("true-d", true, logger.Level(8))

	pout.Log("first example")
	out.Log("second example with location", logger.Here(), logger.Flag(255))
	out.Log("third example with defaults (time/location) with an added time.Time in the Context", true, time.Now().Add(-time.Hour), logger.Level(15))
	ent := logger.NewEntry("fourth example with context").AppendContext(time.Now().Add(-time.Hour))
	out.Log(ent)

	s, e := json.Marshal(logger.Entry{
		Message: "This is a message",
		// Timestamp: time.Now(),
		Location: logger.Here(),
	})

	fmt.Println(ent.MarshalCSV(true))
	fmt.Println(ent.MarshalLV(true))

	fmt.Println(string(s), e)

	var e3 logger.Entry
	e = json.Unmarshal(s, &e3)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Printf("%+v\n", &e3)

	ent = logger.NewEntry("fifth example with flags").AppendContext(time.Now().Add(-time.Hour))
	// ent.Append(logger.Flag(34))
	// ent.Append(logger.Flag(54))
	// ent.Append(logger.Flag(2))
	fmt.Println(ent.MarshalCSV(true))
	fmt.Println(ent.MarshalLV(true))

}
