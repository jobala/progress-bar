package core

import "fmt"

type Terminal struct {
	CursorX int
	CursorY int
}

var Esc = "\x1b"

func escape(format string, args ...interface{}) string {
	return fmt.Sprintf("%s%s", Esc, fmt.Sprintf(format, args...))
}

func (terminal *Terminal) MoveCursorTo(line, col int) string {
	terminal.CursorX = line
	terminal.CursorY = col
	return escape("[%d;%dH", line, col)
}

func (terminal *Terminal) Write(msg string) {
	escaped := escape("[%d;%dH", terminal.CursorX, terminal.CursorY)
	fmt.Print(escaped, msg)
}
