package watch

import (
	"context"
	"sync"

	"go.opentelemetry.io/otel/metric"
)

const (
	gvkCountMetricName       = "watch_manager_watched_gvk"
	gvkIntentCountMetricName = "watch_manager_intended_watch_gvk"
)

var r *reporter

func (r *reporter) reportGvkCount(count int64) error { _ = "STUB: not implemented"; return nil }

func (r *reporter) reportGvkIntentCount(count int64) error { _ = "STUB: not implemented"; return nil }

// newStatsReporter creates a reporter for watch metrics.
func newStatsReporter() (*reporter, error) { _ = "STUB: not implemented"; return nil, nil }

type reporter struct {
	mu          sync.RWMutex
	gvkCount    int64
	intentCount int64
}

func (r *reporter) observeGvkCount(_ context.Context, observer metric.Int64Observer) error {
	_ = "STUB: not implemented"
	return nil
}

// count returns total gvk count across all registrars.
func (r *reporter) observeGvkIntentCount(_ context.Context, observer metric.Int64Observer) error {
	_ = "STUB: not implemented"
	return nil
}
