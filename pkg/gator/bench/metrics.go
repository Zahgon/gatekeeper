package bench

import (
	"time"
)

// calculateLatencies computes latency statistics from a slice of durations.
func calculateLatencies(durations []time.Duration) Latencies {
	_ = "STUB: not implemented"
	return *new(Latencies)
}

// Sort for percentile calculation

// percentile calculates the p-th percentile from a sorted slice of durations.
// The input slice must be sorted in ascending order.
func percentile(sorted []time.Duration, p float64) time.Duration {
	_ = "STUB: not implemented"
	return *new(time.Duration)
}

// Calculate the index using the nearest-rank method

// Linear interpolation between the two nearest ranks

// calculateThroughput computes reviews per second.
func calculateThroughput(reviewCount int, duration time.Duration) float64 {
	_ = "STUB: not implemented"
	return 0
}
