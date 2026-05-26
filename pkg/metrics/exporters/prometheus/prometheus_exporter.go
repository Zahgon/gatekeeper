package prometheus

import (
	"context"
	"flag"
	"net/http"
	"time"

	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

const (
	Name              = "prometheus"
	namespace         = "gatekeeper"
	readHeaderTimeout = 60 * time.Second
)

var (
	log            = logf.Log.WithName("prometheus-exporter")
	prometheusPort = flag.Int("prometheus-port", 8888, "Prometheus port for metrics backend")
)

func Start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

func newPromSrv(port int) *http.Server { _ = "STUB: not implemented"; return nil }
