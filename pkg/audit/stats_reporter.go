package audit

import (
	"context"
	"sync"
	"time"

	"github.com/open-policy-agent/gatekeeper/v3/pkg/metrics/exporters/view"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/util"
	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
)

const (
	violationsMetricName       = "violations"
	auditDurationMetricName    = "audit_duration_seconds"
	lastRunStartTimeMetricName = "audit_last_run_time"
	lastRunEndTimeMetricName   = "audit_last_run_end_time"
	enforcementActionKey       = "enforcement_action"
)

var auditDurationM metric.Float64Histogram

func init() {
	view.Register(sdkmetric.NewView(
		sdkmetric.Instrument{Name: auditDurationMetricName},
		sdkmetric.Stream{
			Aggregation: sdkmetric.AggregationExplicitBucketHistogram{
				Boundaries: []float64{1 * 60, 3 * 60, 5 * 60, 10 * 60, 15 * 60, 20 * 60, 40 * 60, 80 * 60, 160 * 60, 320 * 60},
			},
		},
	))
}

func (r *reporter) observeTotalViolations(_ context.Context, o metric.Int64Observer) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *reporter) reportTotalViolations(enforcementAction util.EnforcementAction, v int64) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *reporter) reportRunStart(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (r *reporter) reportLatency(d time.Duration) error { _ = "STUB: not implemented"; return nil }

func (r *reporter) reportRunEnd(t time.Time) error { _ = "STUB: not implemented"; return nil }

func (r *reporter) observeRunStart(_ context.Context, o metric.Float64Observer) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *reporter) observeRunEnd(_ context.Context, o metric.Float64Observer) error {
	_ = "STUB: not implemented"
	return nil
}

// newStatsReporter creates a reporter for audit metrics.
func newStatsReporter() (*reporter, error) { _ = "STUB: not implemented"; return nil, nil }

type reporter struct {
	mu                                  sync.RWMutex
	endTime                             time.Time
	startTime                           time.Time
	totalViolationsPerEnforcementAction map[util.EnforcementAction]int64
}
