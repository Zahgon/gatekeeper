/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package watch

import (
	"context"
	"sync"

	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/cache"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

var log = logf.Log.WithName("watch-manager")

// Manager allows us to dynamically configure what kinds are watched.
type Manager struct {
	cache      RemovableCache
	startedMux sync.Mutex
	stopped    chan struct{}
	// started is a bool
	started bool
	// managedKinds stores the kinds that should be managed, mapping CRD Kind to CRD Name
	managedKinds *recordKeeper
	watchedMux   sync.RWMutex
	// watchedKinds are the kinds that have a currently running constraint controller
	watchedKinds vitalsByGVK
	metrics      *reporter

	// replayTracker allows us to block until replays are complete (or fully canceled)
	replayTracker *replayTracker

	// Events are passed internally from informer event handlers to handleEvents for distribution.
	events chan interface{}
	// replayRequests is used to request or cancel replay for a registrar joining an existing watch.
	replayRequests chan replayRequest
}

type AddFunction func(manager.Manager) error

// RemovableCache is a subset variant of the cache.Cache interface.
// It supports non-blocking calls to get informers, as well as the
// ability to remove an informer dynamically.
type RemovableCache interface {
	GetInformer(_ context.Context, obj client.Object, opts ...cache.InformerGetOption) (cache.Informer, error)
	List(ctx context.Context, list client.ObjectList, opts ...client.ListOption) error
	RemoveInformer(_ context.Context, obj client.Object) error
}

func New(c RemovableCache) (*Manager, error) { _ = "STUB: not implemented"; return nil, nil }

func (wm *Manager) NewRegistrar(parent string, events chan<- event.GenericEvent) (*Registrar, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// RemoveRegistrar removes a registrar and all its watches.
func (wm *Manager) RemoveRegistrar(parentName string) error { _ = "STUB: not implemented"; return nil }

// Start runs the watch manager, processing events received from dynamic informers and distributing them
// to registrars.
func (wm *Manager) Start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// Unblock any informer event handlers

// Routine for distributing events to listeners.

// Routine for asynchronous replay of past events to joining listeners.

func (wm *Manager) checkStarted() error { _ = "STUB: not implemented"; return nil }

func (wm *Manager) GetManagedGVK() []schema.GroupVersionKind { _ = "STUB: not implemented"; return nil }

func (wm *Manager) addWatch(ctx context.Context, r *Registrar, gvk schema.GroupVersionKind) error {
	_ = "STUB: not implemented"
	return nil
}

func (wm *Manager) doAddWatch(ctx context.Context, r *Registrar, gvk schema.GroupVersionKind) error {
	_ = "STUB: not implemented"
	// lock acquired by caller
	return nil
}

// watchers is everyone who is *already* watching.

// m is everyone who *wants* to watch.
// Not a deadlock but beware if assumptions change...

// Sanity

// Already watching.

// Someone else was watching, replay events in the cache to the new watcher.

// This is expected to fail if a CRD is unregistered.

// First watcher gets a fresh informer, register for events.

// Mark it as watched.

func (wm *Manager) removeWatch(ctx context.Context, r *Registrar, gvk schema.GroupVersionKind) error {
	_ = "STUB: not implemented"
	return nil
}

func (wm *Manager) doRemoveWatch(ctx context.Context, r *Registrar, gvk schema.GroupVersionKind) error {
	_ = "STUB: not implemented"
	// lock acquired by caller
	return nil
}

// Not watching.

// Cancel any replays that may be pending

// Remove this registrar from the watch list

// Skip if there are additional watchers that would prevent us from removing it

// Wait until all replays have exited before canceling the watch,
// otherwise the list may unintentionally restart a watch

// replaceWatches ensures all and only desired watches are running.
func (wm *Manager) replaceWatches(ctx context.Context, r *Registrar) error {
	_ = "STUB: not implemented"
	return nil
}

// This registrar still desires this gvk, skip.

// Add desired watches. This is idempotent for existing watches.

// OnAdd implements cache.ResourceEventHandler. Called by informers.
func (wm *Manager) OnAdd(obj interface{}, _ bool) {
	_ = "STUB: not implemented"
	// Send event to eventLoop() for processing
	return
}

// OnUpdate implements cache.ResourceEventHandler. Called by informers.
func (wm *Manager) OnUpdate(oldObj, newObj interface{}) {
	_ = "STUB: not implemented"
	// Send event to eventLoop() for processing
	return
}

// OnDelete implements cache.ResourceEventHandler. Called by informers.
func (wm *Manager) OnDelete(obj interface{}) {
	_ = "STUB: not implemented"
	// Send event to eventLoop() for processing
	return
}

// eventLoop receives events from informer callbacks and distributes them to registrars.
func (wm *Manager) eventLoop(stop <-chan struct{}) { _ = "STUB: not implemented"; return }

// distributeEvent distributes a single event to all registrars listening for that resource kind.
func (wm *Manager) distributeEvent(stop <-chan struct{}, obj interface{}) {
	_ = "STUB: not implemented"
	return
}

// Invalid object, drop it

// Critical lock section

// Nobody is watching, drop it

// TODO(OREN) reduce allocations here

// Distribute the event

// TODO(OREN) add timeout
