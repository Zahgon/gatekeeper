package view

import (
	"go.opentelemetry.io/otel/sdk/metric"
)

var views []metric.View

func init() {
	views = []metric.View{}
}

func Register(v ...metric.View) { _ = "STUB: not implemented"; return }

func Views() []metric.View { _ = "STUB: not implemented"; return nil }
