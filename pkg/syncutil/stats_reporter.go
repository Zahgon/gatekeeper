package syncutil

import (
	"context"
	"sync"
	"time"

	"github.com/open-policy-agent/gatekeeper/v3/pkg/metrics"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/metrics/exporters/view"
	"go.opentelemetry.io/otel/metric"
	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

var log = logf.Log.WithName("reporter").WithValues("metaKind", "Sync")

const (
	syncMetricName         = "sync"
	syncDurationMetricName = "sync_duration_seconds"
	lastRunTimeMetricName  = "sync_last_run_time"
	kindKey                = "kind"
	statusKey              = "status"
)

var (
	syncDurationM metric.Float64Histogram
	r             *Reporter
)

type MetricsCache struct {
	mux        sync.RWMutex
	KnownKinds map[string]bool
	Cache      map[string]Tags
}

type Tags struct {
	Kind   string
	Status metrics.Status
}

func NewMetricsCache() *MetricsCache { _ = "STUB: not implemented"; return nil }

func GetKeyForSyncMetrics(namespace string, name string) string {
	_ = "STUB: not implemented"
	return ""
}

// need to know encountered kinds to reset metrics for that kind
// this is a known memory leak
// footprint should naturally reset on Pod upgrade b/c the container restarts.
func (c *MetricsCache) AddKind(key string) { _ = "STUB: not implemented"; return }

func (c *MetricsCache) ResetCache() { _ = "STUB: not implemented"; return }

func (c *MetricsCache) AddObject(key string, t Tags) { _ = "STUB: not implemented"; return }

func (c *MetricsCache) DeleteObject(key string) { _ = "STUB: not implemented"; return }

func (c *MetricsCache) GetTags(key string) *Tags { _ = "STUB: not implemented"; return nil }

func (c *MetricsCache) HasObject(key string) bool { _ = "STUB: not implemented"; return false }

func (c *MetricsCache) ReportSync() { _ = "STUB: not implemented"; return }

type Reporter struct {
	mu         sync.RWMutex
	lastRun    float64
	syncReport map[Tags]int64
	now        func() float64
}

func init() {
	view.Register(
		sdkmetric.NewView(
			sdkmetric.Instrument{Name: syncDurationMetricName},
			sdkmetric.Stream{
				Aggregation: sdkmetric.AggregationExplicitBucketHistogram{
					Boundaries: []float64{0.0001, 0.0002, 0.0003, 0.0004, 0.0005, 0.0006, 0.0007, 0.0008, 0.0009, 0.001, 0.002, 0.003, 0.004, 0.005, 0.01, 0.02, 0.03, 0.04, 0.05},
				},
			},
		))
}

// NewStatsReporter creates a reporter for sync metrics.
func NewStatsReporter() (*Reporter, error) { _ = "STUB: not implemented"; return nil, nil }

func (r *Reporter) ReportSyncDuration(d time.Duration) error { _ = "STUB: not implemented"; return nil }

func (r *Reporter) ReportLastSync() error { _ = "STUB: not implemented"; return nil }

func (r *Reporter) ReportSync(t Tags, v int64) error { _ = "STUB: not implemented"; return nil }

func (r *Reporter) observeLastSync(_ context.Context, observer metric.Float64Observer) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reporter) observeSync(_ context.Context, observer metric.Int64Observer) error {
	_ = "STUB: not implemented"
	return nil
}

// now returns the timestamp as a second-denominated float.
func now() float64 { _ = "STUB: not implemented"; return 0 }
