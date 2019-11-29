package main

import (
	"./core"
	"fmt"
)

func main() {
	fmt.Print(core.MoveCursorTo(2, 0), "Hello world")
	fmt.Print(core.MoveCursorTo(2, 15), "Hello world")
	fmt.Print(core.MoveCursorTo(2, 30), "Hello world")
}
