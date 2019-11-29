package core

import (
	"go/types"
	"time"
)

type ProgressBar struct {
	value           int
	total           int
	payload         types.Nil
	startTime       time.Time
	lastDrawnString types.Nil
	active          bool
	lasRedraw       time.Time
}

func NewProgressBar() ProgressBar {
	return ProgressBar{
		value:           0,
		total:           100,
		payload:         nil,
		startTime:       nil,
		lastDrawnString: nil,
		active:          false,
		lasRedraw:       time.Now(),
	}
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
