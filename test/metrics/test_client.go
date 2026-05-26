package metrics

import (
	"context"

	sdkmetric "go.opentelemetry.io/otel/sdk/metric"
	"go.opentelemetry.io/otel/sdk/metric/metricdata"
)

type FnExporter struct {
	temporalityFunc sdkmetric.TemporalitySelector
	aggregationFunc sdkmetric.AggregationSelector
	exportFunc      func(context.Context, *metricdata.ResourceMetrics) error
	flushFunc       func(context.Context) error
	shutdownFunc    func(context.Context) error
}

func (e *FnExporter) Temporality(k sdkmetric.InstrumentKind) metricdata.Temporality {
	_ = "STUB: not implemented"
	return *new(metricdata.Temporality)
}

func (e *FnExporter) Aggregation(k sdkmetric.InstrumentKind) sdkmetric.Aggregation {
	_ = "STUB: not implemented"
	return *new(sdkmetric.Aggregation)
}

func (e *FnExporter) Export(ctx context.Context, m *metricdata.ResourceMetrics) error {
	_ = "STUB: not implemented"
	return nil
}

func (e *FnExporter) ForceFlush(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func (e *FnExporter) Shutdown(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
