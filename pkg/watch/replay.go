package watch

import (
	"context"
	"sync"

	"k8s.io/apimachinery/pkg/runtime/schema"
)

type replayRequest struct {
	r        *Registrar
	gvk      schema.GroupVersionKind
	isCancel bool // If true, this is a request to cancel a pending replay.
}

type cancelMap map[*Registrar]context.CancelFunc

func (m *cancelMap) Set(r *Registrar, c context.CancelFunc) { _ = "STUB: not implemented"; return }

// replayCounter lists the number of replays by GVK and lets users
// wait until all replays for a GVK have been canceled. A WaitGroup
// would be risky to use here, as it only allows for one call to Wait().
type replayTracker struct {
	mux      sync.Mutex
	counts   map[schema.GroupVersionKind]int
	channels map[schema.GroupVersionKind]chan struct{}
	intent   map[*Registrar]map[schema.GroupVersionKind]bool
}

func newReplayTracker() *replayTracker { _ = "STUB: not implemented"; return nil }

// SetIntent sets whether a registrar wants a replay or not.
// Setting this before sending a replay request avoids a
// race condition where a replay request is canceled before
// the original request is sent.
func (r *replayTracker) setIntent(reg *Registrar, gvk schema.GroupVersionKind, wantReplay bool) {
	_ = "STUB: not implemented"
	return
}

// ReplayIntended returns whether a given registrar still wants a replay.
func (r *replayTracker) replayIntended(reg *Registrar, gvk schema.GroupVersionKind) bool {
	_ = "STUB: not implemented"
	return false
}

// Add a GVK to the replay counter.
func (r *replayTracker) addReplay(gvk schema.GroupVersionKind) { _ = "STUB: not implemented"; return }

// Done decrements the replay counter for a GVK by 1.
func (r *replayTracker) replayDone(gvk schema.GroupVersionKind) { _ = "STUB: not implemented"; return }

func (r *replayTracker) replayWaitCh(gvk schema.GroupVersionKind) chan struct{} {
	_ = "STUB: not implemented"
	return nil
}

// null channels hang, return a closed channel instead

// replayEventsLoop processes requests to start and stop replaying events for
// registrars that join an existing informer and need historical data.
// Events for each registrar will be listed and replayed in an independent goroutine.
func (wm *Manager) replayEventsLoop(ctx context.Context) func() error {
	_ = "STUB: not implemented"
	return nil
}

// Entries remain until a watch is removed.

// Handle cancellation requests

// Cancel the pending replay.
// Note we do not wait for the individual goroutine to complete,
// but we will sync on all of them when stopping the watch manager.

// Replay in progress or cancel request, either way, do not proceed to replay again.

// A requested replay has since been canceled by the registrar, do not start it

// Handle replay requests

// Log and retry w/ backoff

// Give up

// Success

// requestReplay sends a request to replayEventsLoop to start replaying for the specified registrar.
// If a replay is in progress, this is a no-op.
// NOTE: blocks if the manager is not running.
func (wm *Manager) requestReplay(r *Registrar, gvk schema.GroupVersionKind) {
	_ = "STUB: not implemented"
	return
}

// cancelReplay sends a request to replayEventsLoop to cancel replaying for the specified registrar.
// If no replay is in progress, this is a no-op.
// NOTE: blocks if the manager is not running.
func (wm *Manager) cancelReplay(r *Registrar, gvk schema.GroupVersionKind) {
	_ = "STUB: not implemented"
	return
}

// replayEvents replays all resources of type gvk currently in the cache to the requested registrar.
// This is called when a registrar begins watching an existing informer.
func (wm *Manager) replayEvents(ctx context.Context, r *Registrar, gvk schema.GroupVersionKind) error {
	_ = "STUB: not implemented"
	return nil
}

// Skip replay if there's no channel to deliver to
