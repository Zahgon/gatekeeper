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

	"k8s.io/apimachinery/pkg/api/meta"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type NoopClient struct{}

func (f *NoopClient) Get(_ context.Context, _ client.ObjectKey, _ runtime.Object) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *NoopClient) List(_ context.Context, _ client.ObjectList, _ ...client.ListOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *NoopClient) Create(_ context.Context, _ client.Object, _ ...client.CreateOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *NoopClient) Apply(_ context.Context, _ runtime.ApplyConfiguration, _ ...client.ApplyOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *NoopClient) Delete(_ context.Context, _ client.Object, _ ...client.DeleteOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *NoopClient) Update(_ context.Context, _ client.Object, _ ...client.UpdateOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *NoopClient) Patch(_ context.Context, _ client.Object, _ client.Patch, _ ...client.PatchOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *NoopClient) DeleteAllOf(_ context.Context, _ client.Object, _ ...client.DeleteAllOfOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *NoopClient) Status() client.StatusWriter {
	_ = "STUB: not implemented"
	return *new(client.StatusWriter)
}

func (f *NoopClient) RESTMapper() meta.RESTMapper {
	_ = "STUB: not implemented"
	return *new(meta.RESTMapper)
}

func (f *NoopClient) Scheme() *runtime.Scheme { _ = "STUB: not implemented"; return nil }

type SubResourceNoopClient struct{}

func (f *SubResourceNoopClient) Create(_ context.Context, _, _ client.Object, _ ...client.SubResourceCreateOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *SubResourceNoopClient) Update(_ context.Context, _ client.Object, _ ...client.SubResourceUpdateOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *SubResourceNoopClient) Patch(_ context.Context, _ client.Object, _ client.Patch, _ ...client.SubResourcePatchOption) error {
	_ = "STUB: not implemented"
	return nil
}

func (f *SubResourceNoopClient) Apply(_ context.Context, _ runtime.ApplyConfiguration, _ ...client.SubResourceApplyOption) error {
	_ = "STUB: not implemented"
	return nil
}
