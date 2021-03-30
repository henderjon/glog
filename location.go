package logger

import (
	"path/filepath"
	"runtime"
	"strconv"
	"strings"
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

// CurrFunc returns the name of the function in which CurrFunc is called
func CurrFunc() string {
	pc, _, _, ok := runtime.Caller(1)
	if !ok {
		return "unable to determine function"
		// file = "?"
		// line = 0
	}

	fn := runtime.FuncForPC(pc)
	dotName := filepath.Ext(fn.Name())
	return strings.TrimLeft(dotName, ".") + "()"
}
