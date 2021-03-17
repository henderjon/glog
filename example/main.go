package main

import (
	"encoding/json"
	"fmt"
	"os"
	"time"

	"github.com/henderjon/logger"
)

func main() {
	var out logger.Logger
	out = logger.NewStdLogger(os.Stderr)
	out.Log("first example")
	out.Log("second example with location", logger.Here(), logger.Level(255))
	out.Log("third example with defaults (time/location) with an added time.Time in the Context", true, time.Now().Add(-time.Hour), logger.Level(15))
	ent := logger.NewEntry("fourth example with context").AppendContext(time.Now().Add(-time.Hour))
	out.Log(ent)
	out.Log(logger.Postmark)
	out = logger.NewMultiLog(
		logger.NewStderrLogger(true),
		logger.NewStdoutLogger(true),
	)
	out.Log(ent)
	fmt.Fprintf(out, "%d", logger.Level(5))
	fmt.Fprintf(out, "%s", logger.Level(5))
	s, e := json.Marshal(logger.Entry{
		Message: "This is a message",
		// Timestamp: time.Now(),
		Level:    logger.Level(51),
		Location: logger.Here(),
	})

	fmt.Println(string(s), e)

	var e3 logger.Entry
	e = json.Unmarshal(s, &e3)
	if e != nil {
		fmt.Println(e)
	}
	fmt.Printf("%+v\n", &e3)

}
