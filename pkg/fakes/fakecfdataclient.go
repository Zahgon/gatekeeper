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
package fakes

import (
	"context"
	gosync "sync"

	constraintTypes "github.com/open-policy-agent/frameworks/constraint/pkg/types"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type CfDataKey struct {
	Gvk schema.GroupVersionKind
	Key string
}

// FakeCfClient is an CfDataClient for testing.
type FakeCfClient struct {
	mu           gosync.Mutex
	data         map[CfDataKey]interface{}
	needsToError bool
}

// KeyFor returns a CfDataKey for the provided resource.
// Returns error if the resource is not a runtime.Object w/ metadata.
func KeyFor(obj interface{}) (CfDataKey, error) {
	_ = "STUB: not implemented"
	return *new(CfDataKey), nil
}

func (f *FakeCfClient) AddData(_ context.Context, data interface{}) (*constraintTypes.Responses, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (f *FakeCfClient) RemoveData(_ context.Context, data interface{}) (*constraintTypes.Responses, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// GetData returns data for a CfDataKey. It assumes that the
// key is present in the FakeCfClient. Also the data returned is not copied
// and it's meant only for assertions not modifications.
func (f *FakeCfClient) GetData(key CfDataKey) interface{} { _ = "STUB: not implemented"; return nil }

// Contains returns true if all expected resources are in the cache.
func (f *FakeCfClient) Contains(expected map[CfDataKey]interface{}) bool {
	_ = "STUB: not implemented"
	return false
}

// HasGVK returns true if the cache has any data of the requested kind.
func (f *FakeCfClient) HasGVK(gvk schema.GroupVersionKind) bool {
	_ = "STUB: not implemented"
	return false
}

// ContainsGVKs returns true if the cache has data for the gvks given and those gvks only.
func (f *FakeCfClient) ContainsGVKs(gvks []schema.GroupVersionKind) bool {
	_ = "STUB: not implemented"
	return false
}

// Len returns the number of items in the cache.
func (f *FakeCfClient) Len() int { _ = "STUB: not implemented"; return 0 }

// SetErroring will error out on AddObject or RemoveObject.
func (f *FakeCfClient) SetErroring(enabled bool) { _ = "STUB: not implemented"; return }
