package externaldata

import (
	"context"
	"sync"

	"github.com/open-policy-agent/gatekeeper/v3/pkg/metrics"
	"go.opentelemetry.io/otel/metric"
	"k8s.io/apimachinery/pkg/types"
)

const (
	providerMetricName     = "providers"
	providerErrorCountName = "provider_error_count"
	statusKey              = "status"

	providerDesc      = "Number of external data providers by status"
	providerErrorDesc = "Incremental counter for all provider errors occurring over time"
)

var providerErrorCountM metric.Int64Counter

func (r *reporter) observeProviderMetric(_ context.Context, o metric.Int64Observer) error {
	_ = "STUB: not implemented"
	return nil
}

// newStatsReporter creates a reporter for external data provider metrics.
func newStatsReporter() *reporter { _ = "STUB: not implemented"; return nil }

// Register the gatekeeper_providers gauge metric

// Register the gatekeeper_provider_error_count counter metric

// reportProviderError increments the provider error counter with the specific error type.
func (r *reporter) reportProviderError(ctx context.Context) { _ = "STUB: not implemented"; return }

type reporter struct {
	mu           sync.RWMutex
	cache        map[types.NamespacedName]metrics.Status
	dirty        bool
	statusReport map[metrics.Status]int64
}

func (r *reporter) add(key types.NamespacedName, status metrics.Status) {
	_ = "STUB: not implemented"
	return
}

func (r *reporter) remove(key types.NamespacedName) { _ = "STUB: not implemented"; return }

func (r *reporter) report(_ context.Context) { _ = "STUB: not implemented"; return }
