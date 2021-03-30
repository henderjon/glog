package logger

import (
	"fmt"
	"strings"
)

// Flags will not render with a zero value, a special case
type Flags []Flag

type Flag int

func (f Flags) String() string {
	var s strings.Builder
	for _, i := range f {
		s.WriteString(i.String())
		s.WriteString(" ")
	}
	return strings.TrimRight(s.String(), " ")
}

func (f Flag) String() string {
	return fmt.Sprintf("f%d", f)
}
