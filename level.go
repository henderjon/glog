package logger

import "fmt"

type Level uint8

func (l Level) String() string {
	return fmt.Sprintf("l%d", l)
}
