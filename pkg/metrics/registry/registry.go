// Package registry provides a dynamic registry of available exporters.
// this makes it easy for forks to inject new metrics exporters as-needed.
package registry

import (
	"context"
	"flag"

	// register exporters with the registry.
	"github.com/open-policy-agent/gatekeeper/v3/pkg/metrics/exporters/opentelemetry"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/metrics/exporters/prometheus"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/metrics/exporters/stackdriver"
)

func init() {
	flag.Var(exporters, "metrics-backend", "Backend used for metrics. e.g. `prometheus`, `stackdriver`. This flag can be declared more than once. Omitting will default to supporting `prometheus`.")
}

type Exporter interface {
	// Start starts the exporter behavior. If Start()
	// returns an error, that is bubbled up to the
	// controller manager
	Start(context.Context) error
}

var exporters = newExporterSet(
	map[string]StartExporter{
		opentelemetry.Name: opentelemetry.Start,
		prometheus.Name:    prometheus.Start,
		stackdriver.Name:   stackdriver.Start,
	},
)

type StartExporter func(context.Context) error

type exporterSet struct {
	validExporters      []string
	registeredExporters map[string]StartExporter
	assignedExporters   map[string]StartExporter
}

var _ flag.Value = &exporterSet{}

func newExporterSet(exporters map[string]StartExporter) *exporterSet {
	_ = "STUB: not implemented"
	return nil
}

func (es *exporterSet) String() string { _ = "STUB: not implemented"; return "" }

func (es *exporterSet) Set(s string) error { _ = "STUB: not implemented"; return nil }

func (es *exporterSet) MustRegister(name string, new StartExporter) {
	_ = "STUB: not implemented" // nolint:revive
	return
}

func Exporters() []StartExporter { _ = "STUB: not implemented"; return nil }
