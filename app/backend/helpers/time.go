package helpers

import "time"

// Returns its execution time.
func MeasureLatency(next func()) time.Duration {
	start := time.Now()
	next()
	return time.Since(start)
}
