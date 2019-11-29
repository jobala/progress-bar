package core

import "fmt"

var Esc = "\x1b"

func escape(format string, args ...interface{}) string {
	return fmt.Sprintf("%s%s", Esc, fmt.Sprintf(format, args...))
}

func MoveCursorTo(line, col int) string {
	return escape("[%d;%dH", line, col)
}
