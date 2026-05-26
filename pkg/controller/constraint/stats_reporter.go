package constraint

import (
	"context"
	"sync"

	"github.com/open-policy-agent/gatekeeper/v3/pkg/metrics"
	"go.opentelemetry.io/otel/metric"
	"k8s.io/apimachinery/pkg/types"
)

const (
	constraintsMetricName = "constraints"
	vapbMetricName        = "validating_admission_policy_bindings"
	enforcementActionKey  = "enforcement_action"
	statusKey             = "status"
)

func (r *reporter) observeConstraints(_ context.Context, observer metric.Int64Observer) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *reporter) observeVAPB(_ context.Context, observer metric.Int64Observer) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *reporter) reportConstraints(_ context.Context, t tags, v int64) error {
	_ = "STUB: not implemented"
	return nil
}

// StatsReporter reports audit metrics.
type StatsReporter interface {
	reportConstraints(ctx context.Context, t tags, v int64) error
	ReportVAPBStatus(name types.NamespacedName, status metrics.VAPStatus)
	DeleteVAPBStatus(name types.NamespacedName)
}

// newStatsReporter creates a reporter for audit metrics.
func newStatsReporter() (*reporter, error) { _ = "STUB: not implemented"; return nil, nil }

type reporter struct {
	mux               sync.RWMutex
	constraintsReport map[tags]int64
	vapbRegistry      *metrics.VAPStatusRegistry
}

func (r *reporter) ReportVAPBStatus(name types.NamespacedName, status metrics.VAPStatus) {
	_ = "STUB: not implemented"
	return
}

func (r *reporter) DeleteVAPBStatus(name types.NamespacedName) { _ = "STUB: not implemented"; return }
