package main

import (
	"./core"
	"fmt"
	"time"
)

func main() {
	i := 0
	col := 0

	for i < 3 {
		fmt.Print(core.MoveCursorTo(2, col), "Hello world")
		time.Sleep(2 * time.Second)
		col += 15
		i++
	}
}
