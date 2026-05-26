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
	"sync"

	"k8s.io/apimachinery/pkg/runtime/schema"
)

// Set tracks a set of watched resource GVKs.
type Set struct {
	mux sync.RWMutex
	set map[schema.GroupVersionKind]bool
}

// RLock acquires a read lock on Set.
func (w *Set) RLock() {
	_ = "STUB: not implemented"

	// RUnlock releases a read lock on Set.
	return
}

func (w *Set) RUnlock() {
	_ = "STUB: not implemented"

	// DoForEach locks Set to prevent mutations and executes f on every element
	// currently in the set.
	// Exits early if f returns an error.
	return
}

func (w *Set) DoForEach(f func(gvk schema.GroupVersionKind) error) error {
	_ = "STUB: not implemented"
	return nil
}

// NewSet constructs a new watchSet.
func NewSet() *Set { _ = "STUB: not implemented"; return nil }

// SetFrom constructs a new watchSet from the given gvks.
func SetFrom(items []schema.GroupVersionKind) *Set { _ = "STUB: not implemented"; return nil }

func (w *Set) Size() int { _ = "STUB: not implemented"; return 0 }

func (w *Set) Items() []schema.GroupVersionKind { _ = "STUB: not implemented"; return nil }

func (w *Set) String() string { _ = "STUB: not implemented"; return "" }

func (w *Set) Add(gvks ...schema.GroupVersionKind) { _ = "STUB: not implemented"; return }

func (w *Set) Remove(gvks ...schema.GroupVersionKind) { _ = "STUB: not implemented"; return }

func (w *Set) Dump() map[schema.GroupVersionKind]bool { _ = "STUB: not implemented"; return nil }

func (w *Set) AddSet(other *Set) { _ = "STUB: not implemented"; return }

func (w *Set) RemoveSet(other *Set) { _ = "STUB: not implemented"; return }

func (w *Set) Equals(other *Set) bool { _ = "STUB: not implemented"; return false }

// Replace locks Set for mutation, replaces Set with other, and then executes
// any passed callbacks before releasing the lock.
func (w *Set) Replace(other *Set, fns ...func()) { _ = "STUB: not implemented"; return }

func (w *Set) Contains(gvk schema.GroupVersionKind) bool { _ = "STUB: not implemented"; return false }

// Difference returns items in the set that are not in the other (provided) set.
func (w *Set) Difference(other *Set) *Set { _ = "STUB: not implemented"; return nil }

// Intersection returns a set composed of all items that are both in set w and other.
func (w *Set) Intersection(other *Set) *Set { _ = "STUB: not implemented"; return nil }
