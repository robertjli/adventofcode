package util

import (
	"fmt"
	"time"
)

// StartTiming begins a timer and returns the function to stop the timer.
// The elapsed time is printed when the stop function is called.
func StartTiming() (stop func()) {
	start := time.Now()
	return func() {
		fmt.Printf("Timing: %s\n", time.Since(start))
	}
}
