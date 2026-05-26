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
	"k8s.io/apimachinery/pkg/runtime/schema"
)

type gvkErr struct {
	err      error
	gvk      schema.GroupVersionKind
	isRemove bool
}

func (w gvkErr) String() string { _ = "STUB: not implemented"; return "" }

func (w gvkErr) Error() string { _ = "STUB: not implemented"; return "" }

// ErrorList is an error that aggregates multiple errors.
type ErrorList struct {
	errs          []error
	hasGeneralErr bool
}

func NewErrorList() *ErrorList { _ = "STUB: not implemented"; return nil }

func (e *ErrorList) String() string { _ = "STUB: not implemented"; return "" }

func (e *ErrorList) Error() string { _ = "STUB: not implemented"; return "" }

// Return gvks for which there were errors adding watches.
func (e *ErrorList) AddGVKFailures() []schema.GroupVersionKind {
	_ = "STUB: not implemented"
	return nil
}

// Return gvks for which there were errors removing watches.
func (e *ErrorList) RemoveGVKFailures() []schema.GroupVersionKind {
	_ = "STUB: not implemented"
	return nil
}

func (e *ErrorList) HasGeneralErr() bool { _ = "STUB: not implemented"; return false }

// adds a non gvk specific error to the list.
func (e *ErrorList) Err(err error) { _ = "STUB: not implemented"; return }

// adds a gvk specific error for failing to add a gvk watch to the list.
func (e *ErrorList) AddGVKErr(gvk schema.GroupVersionKind, err error) {
	_ = "STUB: not implemented"
	return
}

// adds a gvk specific error for failing to remove a gvk watch to the list.
func (e *ErrorList) RemoveGVKErr(gvk schema.GroupVersionKind, err error) {
	_ = "STUB: not implemented"
	return
}

func (e *ErrorList) Size() int { _ = "STUB: not implemented"; return 0 }
