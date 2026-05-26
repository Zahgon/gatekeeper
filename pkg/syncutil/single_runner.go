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

package syncutil

import (
	"context"
	"sync"
)

// SingleRunner wraps an errgroup to run keyed goroutines as singletons.
// Keys are single-use and subsequent usage to schedule will be silently ignored.
// Goroutines can be individually canceled provided they respect the context passed to them.
type SingleRunner struct {
	m  map[string]context.CancelFunc
	mu sync.RWMutex
	wg *sync.WaitGroup
	ec chan<- error
}

// NewSingleRunner returns an initialized SingleRunner.
func NewSingleRunner(errChan chan<- error) *SingleRunner { _ = "STUB: not implemented"; return nil }

// Wait waits for all goroutines managed by the SingleRunner to complete.
// Returns the first error returned from a managed goroutine, or nil.
func (s *SingleRunner) Wait() { _ = "STUB: not implemented"; return }

// Go schedules the provided function on a new goroutine if the provided key has
// not been used for scheduling before.
func (s *SingleRunner) Go(ctx context.Context, key string, f func(context.Context, chan<- error)) {
	_ = "STUB: not implemented"
	return
}

// Reject if already running

// Cancel cancels a keyed goroutine if it exists.
func (s *SingleRunner) Cancel(key string) { _ = "STUB: not implemented"; return }

// Leave the key in the map to prevent its re-use.
