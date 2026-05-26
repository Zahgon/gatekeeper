package opentelemetry

import (
	"context"
	"flag"
	"time"
)

const (
	Name                          = "opentelemetry"
	defaultMetricsCollectInterval = 10 * time.Second
	defaultMetricsTimeout         = 30 * time.Second
)

var (
	otlpEndPoint   = flag.String("otlp-endpoint", "", "Opentelemetry exporter endpoint")
	metricInterval = flag.Duration("otlp-metric-interval", defaultMetricsCollectInterval, "interval to read metrics for opentelemetry exporter. defaulted to 10 secs if unspecified")
)

func Start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
