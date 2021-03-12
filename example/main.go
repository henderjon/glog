package main

import (
	"time"

	"github.com/henderjon/logger"
)

func main() {
	var out logger.Logger
	out = logger.NewStdoutLogger(true)
	out.Log("first example", "LL5")
	out.Log("second example with location", logger.Here(), logger.Level(255))
	out.Log("third example with defaults (time/location) with an added time.Time in the Context", true, time.Now().Add(-time.Hour), logger.Level(15))
	ent := logger.NewEntry("fourth example with context").AppendContext(time.Now().Add(-time.Hour))
	out.Log(ent)
	out.Log(true)
	out = logger.NewMultiLog(
		logger.NewStderrLogger(true),
		logger.NewStdoutLogger(true),
	)
	out.Log(ent)
}
