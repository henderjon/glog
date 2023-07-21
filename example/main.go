package main

import (
	"fmt"
	"os"
	"time"

	"github.com/henderjon/logger"
)

func main() {

	l := logger.NewRowLogger(os.Stdout,
		logger.SetPrefix("!!!"),
		logger.LogTimestamp(false),
		logger.LogLocation(false),
	)

	fmt.Fprintf(l, "%s", "this is a new message")

	l.Log("tags", logger.Tag(15), logger.Tag(16), "timestamp", time.Now().Add(-time.Hour))

	l.Log(func(i int) int {
		return i
	})

	l.Log(struct {
		one   int
		two   int
		three int
	}{
		4, 5, 6,
	})

}
