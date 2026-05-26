package stackdriver

import (
	"context"
	"flag"
	"time"

	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

const (
	Name                          = "stackdriver"
	metricPrefix                  = "custom.googleapis.com/opencensus/gatekeeper"
	defaultMetricsCollectInterval = 10 * time.Second
)

var (
	ignoreMissingCreds = flag.Bool("stackdriver-only-when-available", false, "Only attempt to start the stackdriver exporter if credentials are available")
	metricInterval     = flag.Duration("stackdriver-metric-interval", defaultMetricsCollectInterval, "interval to read metrics for stackdriver exporter. defaulted to 10 secs if unspecified")
	log                = logf.Log.WithName("stackdriver-exporter")
)

func Start(ctx context.Context) error {
	_ = "STUB: not implemented"
	// Verify that default stackdriver credentials are available
	return nil
}
