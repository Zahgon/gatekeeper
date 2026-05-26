package constrainttemplate

import (
	"context"
	"sync"
	"time"

	"github.com/open-policy-agent/gatekeeper/v3/pkg/metrics"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/metrics/exporters/view"
	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"k8s.io/apimachinery/pkg/types"
)

const (
	ctMetricName = "constraint_templates"
	// celCTMetricName is a separate metric for backward compatibility with existing metrics. Together with ctMetricName, it allows to derive the number of non-CEL constraint templates.
	// Example: rego_only_count = constraint_templates - constraint_templates_with_cel.
	celCTMetricName = "constraint_templates_with_cel"
	vapMetricName   = "validating_admission_policies"
	ingestCount     = "constraint_template_ingestion_count"
	ingestDuration  = "constraint_template_ingestion_duration_seconds"
	statusKey       = "status"

	ctDesc    = "Number of observed constraint templates"
	celCTDesc = "Number of constraint templates with CEL engine"
)

var (
	ingestCountM    metric.Int64Counter
	ingestDurationM metric.Float64Histogram
)

func init() {
	view.Register(sdkmetric.NewView(
		sdkmetric.Instrument{Name: ingestDuration},
		sdkmetric.Stream{
			Aggregation: sdkmetric.AggregationExplicitBucketHistogram{
				Boundaries: []float64{0.01, 0.02, 0.03, 0.04, 0.05, 0.06, 0.07, 0.08, 0.09, 0.1, 0.2, 0.3, 0.4, 0.5, 1, 2, 3, 4, 5},
			},
		},
	))
}

func (r *reporter) reportIngestDuration(ctx context.Context, status metrics.Status, d time.Duration) error {
	_ = "STUB: not implemented"
	return nil
}

// newStatsReporter creates a reporter for watch metrics.
func newStatsReporter() *reporter { _ = "STUB: not implemented"; return nil }

type reporter struct {
	mu          sync.RWMutex
	ctReport    map[metrics.Status]int64
	registry    *ctRegistry
	vapRegistry *metrics.VAPStatusRegistry
	celRegistry *celRegistry
}

type celRegistry struct {
	mu    sync.RWMutex
	cache map[types.NamespacedName]bool
}

func newCelRegistry() *celRegistry { _ = "STUB: not implemented"; return nil }

func (r *celRegistry) add(key types.NamespacedName) { _ = "STUB: not implemented"; return }

func (r *celRegistry) remove(key types.NamespacedName) { _ = "STUB: not implemented"; return }

func (r *celRegistry) count() int64 { _ = "STUB: not implemented"; return 0 }

func (r *reporter) observeCelCTM(_ context.Context, o metric.Int64Observer) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *reporter) ReportCelCT(templateName types.NamespacedName) {
	_ = "STUB: not implemented"
	return
}

func (r *reporter) DeleteCelCT(templateName types.NamespacedName) {
	_ = "STUB: not implemented"
	return
}

type ctRegistry struct {
	cache map[types.NamespacedName]metrics.Status
	dirty bool
}

func (r *ctRegistry) add(key types.NamespacedName, status metrics.Status) {
	_ = "STUB: not implemented"
	return
}

func (r *reporter) reportCtMetric(status metrics.Status, count int64) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *ctRegistry) remove(key types.NamespacedName) { _ = "STUB: not implemented"; return }

func (r *ctRegistry) report(_ context.Context, mReporter *reporter) {
	_ = "STUB: not implemented"
	return
}

func (r *reporter) observeCTM(_ context.Context, o metric.Int64Observer) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *reporter) observeVAP(_ context.Context, observer metric.Int64Observer) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *reporter) ReportVAPStatus(templateName types.NamespacedName, status metrics.VAPStatus) {
	_ = "STUB: not implemented"
	return
}

func (r *reporter) DeleteVAPStatus(templateName types.NamespacedName) {
	_ = "STUB: not implemented"
	return
}
