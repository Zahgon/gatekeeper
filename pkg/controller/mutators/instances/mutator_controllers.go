package instances

import (
	"context"

	"github.com/open-policy-agent/gatekeeper/v3/pkg/mutation"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/readiness"
	corev1 "k8s.io/api/core/v1"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

// eventQueueSize is how many events to queue before blocking.
const eventQueueSize = 1024

type Adder struct {
	MutationSystem *mutation.System
	Tracker        *readiness.Tracker
	GetPod         func(context.Context) (*corev1.Pod, error)
}

func routeConflictEvents(ctx context.Context, events <-chan event.GenericEvent, assignCh, modifySetCh, assignImageCh chan<- event.GenericEvent) error {
	_ = "STUB: not implemented"
	return nil
}

// Add creates all mutation controllers and adds them to the manager.
func (a *Adder) Add(mgr manager.Manager) error {
	_ = "STUB: not implemented"
	// events is shared across all mutators that can affect the implied schema
	// of kinds to be mutated, since these mutators can set each other into conflict
	return nil
}

// Per-controller channels for fan-out of conflict events.

// The type is provided by the `NewObj` function above. If we
// are fed the wrong type, this is a non-recoverable error and we
// may as well crash for visibility
// nolint:forcetypeassert

// The type is provided by the `NewObj` function above. If we
// are fed the wrong type, this is a non-recoverable error and we
// may as well crash for visibility
// nolint:forcetypeassert

// The type is provided by the `NewObj` function above. If we
// are fed the wrong type, this is a non-recoverable error and we
// may as well crash for visibility
// nolint:forcetypeassert

// The type is provided by the `NewObj` function above. If we
// are fed the wrong type, this is a non-recoverable error and we
// may as well crash for visibility
// nolint:forcetypeassert

func (a *Adder) InjectTracker(t *readiness.Tracker) { _ = "STUB: not implemented"; return }

func (a *Adder) InjectGetPod(getPod func(ctx context.Context) (*corev1.Pod, error)) {
	_ = "STUB: not implemented"
	return
}

func (a *Adder) InjectMutationSystem(mutationSystem *mutation.System) {
	_ = "STUB: not implemented"
	return
}
