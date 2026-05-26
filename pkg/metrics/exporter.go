package metrics

import (
	"context"

	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

var log = logf.Log.WithName("metrics")

var _ manager.Runnable = &runner{}

type runner struct {
	mgr manager.Manager
}

func AddToManager(m manager.Manager) error { _ = "STUB: not implemented"; return nil }

func create(mgr manager.Manager) *runner { _ = "STUB: not implemented"; return nil }

// Start implements the Runnable interface.
func (r *runner) Start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }
