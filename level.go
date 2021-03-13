package logger

import "fmt"

// Level will not render with a zero value, a special case
type Level int

func (l Level) String() string {
	return fmt.Sprintf("l%d", l)
}
