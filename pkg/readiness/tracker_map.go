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

package readiness

import (
	"sync"

	"k8s.io/apimachinery/pkg/runtime/schema"
)

type trackerMap struct {
	mu          sync.RWMutex
	m           map[schema.GroupVersionKind]*objectTracker
	removed     map[schema.GroupVersionKind]struct{}
	tryCanceled map[schema.GroupVersionKind]objData
	fn          objDataFactory
}

func newTrackerMap(fn objDataFactory) *trackerMap { _ = "STUB: not implemented"; return nil }

// Has returns true if the map is tracking the requested resource kind.
func (t *trackerMap) Has(gvk schema.GroupVersionKind) bool { _ = "STUB: not implemented"; return false }

// Get returns an objectTracker for the requested resource kind.
// A new one is created if the resource was not previously tracked.
func (t *trackerMap) Get(gvk schema.GroupVersionKind) Expectations {
	_ = "STUB: not implemented"
	return *new(Expectations)
}

// Return a throwaway tracker if it was previously removed.

// avoids https://golang.org/doc/faq#nil_error

// re-retrieve map entry in case it was added after releasing the read lock.

// Keys returns the resource kinds currently being tracked.
func (t *trackerMap) Keys() []schema.GroupVersionKind { _ = "STUB: not implemented"; return nil }

// Remove stops tracking a resource kind. It cannot be tracked again by the same map.
func (t *trackerMap) Remove(gvk schema.GroupVersionKind) { _ = "STUB: not implemented"; return }

func (t *trackerMap) removeNoLock(gvk schema.GroupVersionKind) { _ = "STUB: not implemented"; return }

// Satisfied returns true if all tracked expectations have been satisfied.
func (t *trackerMap) Satisfied() bool { _ = "STUB: not implemented"; return false }

// Populated returns true if all objectTrackers are populated.
func (t *trackerMap) Populated() bool { _ = "STUB: not implemented"; return false }

// TryCancel will check the readinessRetries left on this GVK, and remove
// the expectation for its objectTracker if no retries remain.
// Returns True if it stopped tracking a resource kind.
func (t *trackerMap) TryCancel(g schema.GroupVersionKind) bool {
	_ = "STUB: not implemented"
	return false
}

// need to create a record of this TryCancel call

// set the changed obj back to the map, as the value is not a pointer
