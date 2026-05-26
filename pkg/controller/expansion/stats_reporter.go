package expansion

import (
	"context"
	"sync"

	"github.com/open-policy-agent/gatekeeper/v3/pkg/metrics"
	"go.opentelemetry.io/otel/metric"
	"k8s.io/apimachinery/pkg/types"
)

const (
	etMetricName = "expansion_templates"
	etDesc       = "Number of observed expansion templates"
	statusKey    = "status"
)

func newRegistry() *etRegistry { _ = "STUB: not implemented"; return nil }

type etRegistry struct {
	mu           sync.RWMutex
	cache        map[types.NamespacedName]metrics.Status
	dirty        bool
	statusReport map[metrics.Status]int64
}

func (r *etRegistry) add(key types.NamespacedName, status metrics.Status) {
	_ = "STUB: not implemented"
	return
}

func (r *etRegistry) remove(key types.NamespacedName) { _ = "STUB: not implemented"; return }

func (r *etRegistry) report(_ context.Context) { _ = "STUB: not implemented"; return }

func (r *etRegistry) observeETM(_ context.Context, o metric.Int64Observer) error {
	_ = "STUB: not implemented"
	return nil
}
