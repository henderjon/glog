package logger

import "fmt"

type Level int

func (l Level) String() string {
	return fmt.Sprintf("ll%d", l)
}
