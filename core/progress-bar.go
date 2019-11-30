package core

import (
	"time"
)

type ProgressBar struct {
	value           int
	total           int
	startTime       time.Time
	lastDrawnString string
	active          bool
	lasRedraw       time.Time
	terminal        Terminal
}

func NewProgressBar() ProgressBar {
	return ProgressBar{
		value:           0,
		total:           100,
		startTime:       time.Now(),
		lastDrawnString: "",
		active:          false,
		lasRedraw:       time.Now(),
		terminal:        Terminal{0, 0},
	}
}

func (progressBar *ProgressBar) Start(total int, startValue int) {
	progressBar.value = startValue
	progressBar.total = total

	progressBar.startTime = time.Now()
	progressBar.lastDrawnString = ""

	progressBar.active = true

	progressBar.Render()

	for progressBar.value < progressBar.total {
		progressBar.increment(10)
		time.Sleep(time.Second)
		progressBar.Render()
	}
}

func (progressBar *ProgressBar) Render() {
	config := GetConfig()
	s := Format(progressBar, config)

	if progressBar.lastDrawnString != s {
		progressBar.terminal.MoveCursorTo(2, 0)
		progressBar.terminal.Write(s)
		progressBar.lastDrawnString = s
		progressBar.lasRedraw = time.Now()
	}
}

func (progressBar *ProgressBar) Update(current int) {
	progressBar.value = current

	if progressBar.value >= progressBar.getTotal() {
		progressBar.stop()
	}
}

func (progressBar *ProgressBar) increment(step int) {
	progressBar.Update(progressBar.value + step)
}

func (progressBar *ProgressBar) setTotal(total int) {
	progressBar.total = total
}

func (progressBar *ProgressBar) getTotal() int {
	return progressBar.total
}

func (progressBar *ProgressBar) stop() {
	progressBar.active = false
}
