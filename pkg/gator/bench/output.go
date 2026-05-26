package bench

import (
	"io"
	"time"
)

// OutputFormat represents the output format for benchmark results.
type OutputFormat string

const (
	// OutputFormatTable outputs results as a human-readable table.
	OutputFormatTable OutputFormat = "table"
	// OutputFormatJSON outputs results as JSON.
	OutputFormatJSON OutputFormat = "json"
	// OutputFormatYAML outputs results as YAML.
	OutputFormatYAML OutputFormat = "yaml"
)

// ParseOutputFormat parses a string into an OutputFormat.
func ParseOutputFormat(s string) (OutputFormat, error) {
	_ = "STUB: not implemented"
	return *new(OutputFormat), nil
}

// FormatResults formats benchmark results according to the specified format.
func FormatResults(results []Results, format OutputFormat) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

// FormatComparison formats comparison results for display.
func FormatComparison(comparisons []ComparisonResult, threshold float64) string {
	_ = "STUB: not implemented"
	return ""
}

func writeComparisonResult(w io.Writer, comp *ComparisonResult, threshold float64) {
	_ = "STUB: not implemented"
	return
}

// Header

// Format values based on metric type

func formatJSON(results []Results) (string, error) {
	_ = "STUB: not implemented"
	// Convert to JSON-friendly format with string durations
	return "", nil
}

func formatYAML(results []Results) (string, error) {
	_ = "STUB: not implemented"
	// Convert to YAML-friendly format with string durations
	return "", nil
}

func formatTable(results []Results) string { _ = "STUB: not implemented"; return "" }

// Write individual result tables

// Write comparison table if multiple engines

func writeResultTable(w io.Writer, r *Results) { _ = "STUB: not implemented"; return }

// Configuration section

// Skipped templates/constraints/data warning

// Show first few objects if not too many

// Informational note about engine limitations (not a warning)

// Timing section with breakdown

// Latency section

// Results section

// Memory section (if available)

// Stats section (if available)

// Include StatsFor to identify which constraint/template produced the stat

// writeComparisonTable writes a side-by-side comparison of engine results.
func writeComparisonTable(w io.Writer, results []Results) { _ = "STUB: not implemented"; return }

// Header row

// Separator

// Templates

// Constraints

// Setup Duration

// Throughput

// Mean Latency

// P95 Latency

// P99 Latency

// Violations

// Memory stats (if available)

// Show performance difference if exactly 2 engines

// writePerfDiff writes a performance comparison between two engines.
func writePerfDiff(w io.Writer, r1, r2 *Results) {
	_ = "STUB: not implemented"
	// Calculate throughput ratio
	return
}

// formatDuration formats a duration in a human-readable way.
func formatDuration(d time.Duration) string { _ = "STUB: not implemented"; return "" }

// formatBytes formats bytes in a human-readable way.
func formatBytes(b uint64) string { _ = "STUB: not implemented"; return "" }

// JSONResults is a JSON/YAML-friendly version of Results with string durations.
type JSONResults struct {
	Engine                   string             `json:"engine" yaml:"engine"`
	TemplateCount            int                `json:"templateCount" yaml:"templateCount"`
	ConstraintCount          int                `json:"constraintCount" yaml:"constraintCount"`
	ObjectCount              int                `json:"objectCount" yaml:"objectCount"`
	Iterations               int                `json:"iterations" yaml:"iterations"`
	Concurrency              int                `json:"concurrency,omitempty" yaml:"concurrency,omitempty"`
	TotalReviews             int                `json:"totalReviews" yaml:"totalReviews"`
	SetupDuration            string             `json:"setupDuration" yaml:"setupDuration"`
	SetupBreakdown           JSONSetupBreakdown `json:"setupBreakdown" yaml:"setupBreakdown"`
	TotalDuration            string             `json:"totalDuration" yaml:"totalDuration"`
	Latencies                JSONLatency        `json:"latencies" yaml:"latencies"`
	ViolationCount           int                `json:"violationCount" yaml:"violationCount"`
	ReviewsPerSecond         float64            `json:"reviewsPerSecond" yaml:"reviewsPerSecond"`
	MemoryStats              *JSONMemoryStats   `json:"memoryStats,omitempty" yaml:"memoryStats,omitempty"`
	StatsEntries             []JSONStatsEntry   `json:"statsEntries,omitempty" yaml:"statsEntries,omitempty"`
	SkippedTemplates         []string           `json:"skippedTemplates,omitempty" yaml:"skippedTemplates,omitempty"`
	SkippedConstraints       []string           `json:"skippedConstraints,omitempty" yaml:"skippedConstraints,omitempty"`
	SkippedDataObjects       []string           `json:"skippedDataObjects,omitempty" yaml:"skippedDataObjects,omitempty"`
	ReferentialDataSupported bool               `json:"referentialDataSupported" yaml:"referentialDataSupported"`
}

// JSONSetupBreakdown is a JSON/YAML-friendly version of SetupBreakdown with string durations.
type JSONSetupBreakdown struct {
	ClientCreation      string `json:"clientCreation" yaml:"clientCreation"`
	TemplateCompilation string `json:"templateCompilation" yaml:"templateCompilation"`
	ConstraintLoading   string `json:"constraintLoading" yaml:"constraintLoading"`
	DataLoading         string `json:"dataLoading" yaml:"dataLoading"`
}

// JSONLatency is a JSON/YAML-friendly version of Latencies with string durations.
type JSONLatency struct {
	Min  string `json:"min" yaml:"min"`
	Max  string `json:"max" yaml:"max"`
	Mean string `json:"mean" yaml:"mean"`
	P50  string `json:"p50" yaml:"p50"`
	P95  string `json:"p95" yaml:"p95"`
	P99  string `json:"p99" yaml:"p99"`
}

// JSONMemoryStats is a JSON/YAML-friendly version of MemoryStats.
type JSONMemoryStats struct {
	AllocsPerReview uint64 `json:"allocsPerReview" yaml:"allocsPerReview"`
	BytesPerReview  string `json:"bytesPerReview" yaml:"bytesPerReview"`
	TotalAllocs     uint64 `json:"totalAllocs" yaml:"totalAllocs"`
	TotalBytes      string `json:"totalBytes" yaml:"totalBytes"`
}

// JSONStatsEntry is a JSON/YAML-friendly version of StatsEntry.
type JSONStatsEntry struct {
	Scope    string          `json:"scope" yaml:"scope"`
	StatsFor string          `json:"statsFor,omitempty" yaml:"statsFor,omitempty"`
	Stats    []JSONStat      `json:"stats" yaml:"stats"`
	Labels   []JSONStatLabel `json:"labels,omitempty" yaml:"labels,omitempty"`
}

// JSONStat is a JSON/YAML-friendly version of instrumentation.Stat.
type JSONStat struct {
	Name   string      `json:"name" yaml:"name"`
	Value  interface{} `json:"value" yaml:"value"`
	Source string      `json:"source" yaml:"source"`
}

// JSONStatLabel is a JSON/YAML-friendly version of instrumentation.Label.
type JSONStatLabel struct {
	Name  string      `json:"name" yaml:"name"`
	Value interface{} `json:"value" yaml:"value"`
}

func toJSONResults(results []Results) []JSONResults { _ = "STUB: not implemented"; return nil }

// Add memory stats if available

// Add stats entries if available

// Convert stats

// Convert labels
