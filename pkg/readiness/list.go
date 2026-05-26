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
	"context"

	"sigs.k8s.io/controller-runtime/pkg/client"
)

type listerFunc func(ctx context.Context, out client.ObjectList, opts ...client.ListOption) error

func (f listerFunc) List(ctx context.Context, out client.ObjectList, opts ...client.ListOption) error {
	_ = "STUB: not implemented"
	return nil
}

// retryLister returns a delegating lister that retries until it succeeds or
// its context is canceled.  Optionally, a predicate can be provided to
// determine if errors are transient and the operation should be retried.  If
// the predicate returns false, the error is terminal and the operation will be
// abandoned.  If predicate is nil, all errors are considered recoverable.
func retryLister(r Lister, predicate retryPredicate) Lister {
	_ = "STUB: not implemented"
	return *new(Lister)
}

// Give up when our parent context is canceled

// Log and retry w/ backoff

// Success

// retryPredicate is a function that determines whether an error is recoverable
// in the context of a retryable operation.  If the predicate returns true, the
// operation can be retried. Otherwise, the error is considered terminal.
type retryPredicate func(err error) bool

// retryAll is a retryPredicate that will retry any error.
func retryAll(_ error) bool {
	_ = "STUB: not implemented"

	// retryNone is a retryPredicate that will never retry an error.
	return false
}

func retryNone(_ error) bool {
	_ = "STUB: not implemented"

	// retryUnlessUnregistered is a retryPredicate that retries all errors except
	// *NoResourceMatchError, *NoKindMatchError, e.g. a resource was not registered to
	// the RESTMapper.
	return false
}

func retryUnlessUnregistered(err error) bool {
	_ = "STUB: not implemented"
	// NoKindMatchError is non-recoverable, otherwise we'll retry.
	return false
}
