package main

import (
	"time"

	"github.com/henderjon/logger"
)

func main() {
	out := logger.NewStdoutLogger(true)
	out.Log("first example")
	out.Log("second example with location", logger.Here())
	out.Log("third example with defaults (time/location) with an added time.Time in the Context", true, time.Now().Add(-time.Hour))
	ent := logger.NewEntry("fourth example with context").AppendContext(time.Now().Add(-time.Hour))
	out.Log(ent)
}
