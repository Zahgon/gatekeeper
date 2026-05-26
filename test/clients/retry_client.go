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

package clients

import (
	"context"

	"golang.org/x/time/rate"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var (
	// defaultRefillRate is the default rate at which potential calls are
	// added back to the "bucket" of allowed calls.
	defaultRefillRate = 5
	// defaultLimitSize is the default starting/max number of potential calls
	// per second.  Once a call is used, it's added back to the bucket at a rate
	// of defaultRefillRate per second.
	defaultLimitSize = 5
)

// RetryClient wraps a client to provide rate-limiter respecting retry behavior.
type RetryClient struct {
	Limiter *rate.Limiter
	client.Client
}

func NewRetryClient(c client.Client) *RetryClient { _ = "STUB: not implemented"; return nil }

// retry will run the provided function, retrying if it fails due to rate limiting.
// If context is canceled, it will return early.
func retry(ctx context.Context, limiter *rate.Limiter, f func() error) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *RetryClient) Get(ctx context.Context, key client.ObjectKey, obj client.Object, _ ...client.GetOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *RetryClient) List(ctx context.Context, list client.ObjectList, opts ...client.ListOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *RetryClient) Create(ctx context.Context, obj client.Object, opts ...client.CreateOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *RetryClient) Delete(ctx context.Context, obj client.Object, opts ...client.DeleteOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *RetryClient) Update(ctx context.Context, obj client.Object, opts ...client.UpdateOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *RetryClient) Patch(ctx context.Context, obj client.Object, patch client.Patch, opts ...client.PatchOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *RetryClient) DeleteAllOf(ctx context.Context, obj client.Object, opts ...client.DeleteAllOfOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *RetryClient) Status() client.StatusWriter {
	_ = "STUB: not implemented"
	return *new(client.StatusWriter)
}

type RetryStatusWriter struct {
	client.StatusWriter
	Limiter *rate.Limiter
}

func (c *RetryStatusWriter) Update(ctx context.Context, obj client.Object, opts ...client.SubResourceUpdateOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (c *RetryStatusWriter) Patch(ctx context.Context, obj client.Object, patch client.Patch, opts ...client.SubResourcePatchOption) error {
	_ = "STUB: not implemented"
	return nil
}
