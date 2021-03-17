[![GoDoc](https://godoc.org/github.com/henderjon/logger?status.svg)](https://godoc.org/github.com/henderjon/logger)
[![License: BSD-3](https://img.shields.io/badge/license-BSD--3-blue.svg)](https://img.shields.io/badge/license-BSD--3-blue.svg)
![tag](https://img.shields.io/github/tag/henderjon/logger.svg)
![release](https://img.shields.io/github/release/henderjon/logger.svg)
[![Go Report Card](https://goreportcard.com/badge/github.com/henderjon/logger)](https://goreportcard.com/report/github.com/henderjon/logger)



# logger
This is a simple logger interface.

```golang
type Logger interface {
	Log(args ...interface{})
	Fatal(args ...interface{})
	io.Writer
}
```


