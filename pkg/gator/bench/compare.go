package bench

import (
	"time"
)

// SaveResults saves benchmark results to a file in JSON or YAML format.
// The format is determined by the file extension (.json or .yaml/.yml).
func SaveResults(results []Results, path string) error { _ = "STUB: not implemented"; return nil }

// Default to JSON

// LoadBaseline loads baseline results from a file.
// The format is determined by the file extension (.json or .yaml/.yml).
func LoadBaseline(path string) ([]Results, error) { _ = "STUB: not implemented"; return nil, nil }

// Default to JSON

// Compare compares current results against baseline results and returns comparison data.
// The threshold is the percentage change considered a regression (e.g., 10 means 10%).
// The minThreshold is the minimum absolute difference to consider a regression.
// For latency metrics, positive change = regression. For throughput, negative change = regression.
func Compare(baseline, current []Results, threshold float64, minThreshold time.Duration) []ComparisonResult {
	_ = "STUB: not implemented"
	return nil
}

// Create a map of baseline results by engine for easy lookup

// Compare each current result against its baseline

// No baseline for this engine, skip comparison

func compareResults(baseline, current *Results, threshold float64, minThreshold time.Duration) ComparisonResult {
	_ = "STUB: not implemented"
	return *new(ComparisonResult)
}

// Compare latency metrics (higher is worse, so positive delta = regression)

// For latency, check both percentage threshold AND minimum absolute threshold
// If minThreshold is set, ignore regressions smaller than the absolute minimum

// Compare throughput (lower is worse, so negative delta = regression)

// For throughput, we invert the logic: negative delta is a regression
// If minThreshold is set, convert it to a throughput difference threshold
// A latency increase of minThreshold corresponds to a throughput change that we should ignore

// Calculate the absolute throughput difference

// Convert minThreshold to an equivalent throughput tolerance
// If we tolerate minThreshold latency change, we should tolerate proportional throughput change
// Use baseline throughput to derive a reasonable tolerance from the latency threshold
// throughput ≈ 1/latency, so tolerance should be proportional to baseline throughput

// Compare memory stats if available
// Note: minThreshold is a time.Duration and applies only to latency/throughput metrics.
// Memory metrics are evaluated strictly against the percentage threshold.

// calculateDelta calculates the percentage change from baseline to current.
// Returns positive value if current > baseline (regression for latency metrics).
func calculateDelta(baseline, current float64) float64 { _ = "STUB: not implemented"; return 0 }

// Infinite increase represented as 100%
