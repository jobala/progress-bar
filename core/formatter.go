package core

import "fmt"

func Format(progressBar *ProgressBar, config Config) string {
	complete := (progressBar.value * config.barsize) / progressBar.total
	incomplete := config.barsize - complete

	formatted := ""

	for i := 0; i < complete; i++ {
		formatted += config.barCompleteString
	}

	for i := 0; i < incomplete; i++ {
		formatted += config.barIncompleteString
	}
	return fmt.Sprintf("[%s]\n", formatted)
}
