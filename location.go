package logger

import (
	"path/filepath"
	"runtime"
	"strconv"
)

// Location is the name:line of a file. Ideally returned by Here(). In usage
// it'll give you the file:line of the invocation of Here() to be passed as part
// of the error.
type Location string

// Here returns the file:line at the point of invocation. This is purely sugar.
func Here() Location {
	return here(2)
}

// here returns the file:line at the point of invocation
func here(depth int) Location {
	var l Location
	_, file, line, ok := runtime.Caller(depth)
	if ok {
		path := filepath.Base(file)
		l = Location(path + ":" + strconv.Itoa(line))
	}
	return l
}
