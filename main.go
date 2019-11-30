package main

import (
	"./core"
)

func main() {
	bar := core.NewProgressBar()
	bar.Start(100, 0)
}
