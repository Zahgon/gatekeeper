package mutators

import (
	"context"
	"sync"
	"time"

	"github.com/open-policy-agent/gatekeeper/v3/pkg/metrics/exporters/view"
	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
)

const (
	mutatorIngestionCountMetricName     = "mutator_ingestion_count"
	mutatorIngestionDurationMetricName  = "mutator_ingestion_duration_seconds"
	mutatorsMetricName                  = "mutators"
	mutatorsConflictingCountMetricsName = "mutator_conflicting_count"
	statusKey                           = "status"
)

var (
	mutatorIngestionCountM metric.Int64Counter
	responseTimeInSecM     metric.Float64Histogram
)

// MutatorIngestionStatus defines the outcomes of an attempt to add a Mutator to the mutation System.
type MutatorIngestionStatus string

const (
	// MutatorStatusActive denotes a successfully ingested mutator, ready to mutate objects.
	MutatorStatusActive MutatorIngestionStatus = "active"
	// MutatorStatusError denotes a mutator that failed to ingest.
	MutatorStatusError MutatorIngestionStatus = "error"
)

// StatsReporter reports mutator-related controller metrics.
type StatsReporter interface {
	ReportMutatorIngestionRequest(ms MutatorIngestionStatus, d time.Duration) error
	RegisterTally(statusFn func() map[MutatorIngestionStatus]int, conflictFn func() int)
}

// reporter implements StatsReporter interface.
type reporter struct {
	mu          sync.RWMutex
	statusFns   []func() map[MutatorIngestionStatus]int
	conflictFns []func() int
}

func init() {
	view.Register(sdkmetric.NewView(
		sdkmetric.Instrument{Name: mutatorIngestionDurationMetricName},
		sdkmetric.Stream{
			Aggregation: sdkmetric.AggregationExplicitBucketHistogram{
				Boundaries: []float64{0.001, 0.002, 0.003, 0.004, 0.005, 0.006, 0.007, 0.008, 0.009, 0.01, 0.02, 0.03, 0.04, 0.05},
			},
		},
	))
}

// NewStatsReporter creates a reporter for webhook metrics.
func NewStatsReporter() StatsReporter { _ = "STUB: not implemented"; return *new(StatsReporter) }

// ReportMutatorIngestionRequest reports both the action of a mutator ingestion and the time
// required for this request to complete.  The outcome of the ingestion attempt is recorded via the
// status argument.
func (r *reporter) ReportMutatorIngestionRequest(ms MutatorIngestionStatus, d time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *reporter) RegisterTally(statusFn func() map[MutatorIngestionStatus]int, conflictFn func() int) {
	_ = "STUB: not implemented"
	return
}

// observeMutatorsStatus reports the current number of mutators by status.
// Note: statusFns are called while r.mu is held.
// Registered functions must not call back into this reporter.
func (r *reporter) observeMutatorsStatus(_ context.Context, observer metric.Int64Observer) error {
	_ = "STUB: not implemented"
	return nil
}

// observeMutatorsInConflict reports the current number of conflicting mutators.
// Note: conflictFns are called while r.mu is held.
// Registered functions must not call back into this reporter.
func (r *reporter) observeMutatorsInConflict(_ context.Context, observer metric.Int64Observer) error {
	_ = "STUB: not implemented"
	return nil
}
