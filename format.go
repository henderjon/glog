package logger

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
	"time"
)

func dropLogFormat(depth int, o Opts, args []any) {
	fmt.Printf("\n")
	fmt.Printf("%s %s\n", o.Prefix, flatLoc(depth))
	for _, v := range args {
		switch v := reflect.ValueOf(v); v.Kind() {
		// case reflect.String:
		// 	fmt.Printf("string: %q\n", v.String())
		// case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
		// 	fmt.Printf("int: %q\n", v.Int())
		// case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64:
		// 	fmt.Printf("uint: %q\n", v.Uint())
		default:
			// fmt.Printf("%s\t(%s) [%s] %+v\n", o.Prefix, v.Type(), v.Kind(), v)
			fmt.Printf("%s\t(%s) %+v\n", o.Prefix, v.Type(), v)
		}
	}
	fmt.Printf("\n")
}

func rowLogFormat(depth int, o Opts, args []any) string {
	var s strings.Builder

	if o.Timestamp {
		s.WriteString(time.Now().UTC().Format(time.RFC3339))
		s.WriteString(o.UnitSep)
	}

	if o.Location {
		s.WriteString(flatLoc(depth))
		s.WriteString(o.UnitSep)
	}

	for _, v := range args {
		switch v := reflect.ValueOf(v); v.Kind() {
		default:
			s.WriteString(fmt.Sprintf("%+v", v))
			s.WriteString(o.UnitSep)
		}
	}

	return strings.TrimSpace(s.String())
}

func flatLoc(depth int) string {
	if _, file, line, ok := runtime.Caller(depth); ok {
		return fmt.Sprintf("%s:%d", filepath.Base(file), line)
	}
	return ""
}
