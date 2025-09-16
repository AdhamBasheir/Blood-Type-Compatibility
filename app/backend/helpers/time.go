package helpers

import "time"

// Returns its execution time.
func MeasureLatency(fn func()) time.Duration {
	start := time.Now()
	fn()
	return time.Since(start)
}
