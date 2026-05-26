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

package mutatorstatus

import (
	"context"

	"github.com/go-logr/logr"
	mutationsv1 "github.com/open-policy-agent/gatekeeper/v3/apis/mutations/v1"
	mutationsv1alpha1 "github.com/open-policy-agent/gatekeeper/v3/apis/mutations/v1alpha1"
	"github.com/open-policy-agent/gatekeeper/v3/apis/status/v1beta1"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/logging"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/readiness"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/watch"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var log = logf.Log.WithName("controller").WithValues(logging.Process, "mutator_status_controller")

type Adder struct {
	WatchManager *watch.Manager
}

func (a *Adder) InjectTracker(_ *readiness.Tracker) {
	_ = "STUB: not implemented"

	// Add creates a new Mutator Status Controller and adds it to the Manager. The Manager will set fields on the Controller
	// and Start it when the Manager is Started.
	return
}

func (a *Adder) Add(mgr manager.Manager) error { _ = "STUB: not implemented"; return nil }

// newReconciler returns a new reconcile.Reconciler.
func newReconciler(
	mgr manager.Manager,
) reconcile.Reconciler {
	_ = "STUB: not implemented"
	return *new(reconcile.Reconciler)
}

// Separate reader and writer because manager's default client bypasses the cache for unstructured resources.

type PackerMap func(obj client.Object) []reconcile.Request

// PodStatusToMutatorMapper correlates a MutatorPodStatus with its corresponding mutator.
func PodStatusToMutatorMapper(selfOnly bool, kindMatch string, packerMap handler.MapFunc) handler.TypedMapFunc[*v1beta1.MutatorPodStatus, reconcile.Request] {
	_ = "STUB: not implemented"
	return nil
}

// Do not attempt to reconcile the resource when other pods have changed their status

// AssignImage is the only mutator in v1alpha1 still

func eventPackerMapFuncHardcodeGVKForAssign(gvk schema.GroupVersionKind) handler.TypedMapFunc[*mutationsv1.Assign, reconcile.Request] {
	_ = "STUB: not implemented"
	return nil
}

func eventPackerMapFuncHardcodeGVKForAssignMetadata(gvk schema.GroupVersionKind) handler.TypedMapFunc[*mutationsv1.AssignMetadata, reconcile.Request] {
	_ = "STUB: not implemented"
	return nil
}

func eventPackerMapFuncHardcodeGVKForAssignImage(gvk schema.GroupVersionKind) handler.TypedMapFunc[*mutationsv1alpha1.AssignImage, reconcile.Request] {
	_ = "STUB: not implemented"
	return nil
}

func eventPackerMapFuncHardcodeGVKForModifySet(gvk schema.GroupVersionKind) handler.TypedMapFunc[*mutationsv1.ModifySet, reconcile.Request] {
	_ = "STUB: not implemented"
	return nil
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler.
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	_ = "STUB: not implemented"
	// Create a new controller
	return nil
}

// Watch for changes to MutatorStatus

// Watch for changes to mutators

var _ reconcile.Reconciler = &ReconcileMutatorStatus{}

// ReconcileMutatorStatus reconciles an arbitrary mutator object described by Kind.
type ReconcileMutatorStatus struct {
	reader       client.Reader
	writer       client.Writer
	statusClient client.StatusClient
	scheme       *runtime.Scheme
	log          logr.Logger
}

// +kubebuilder:rbac:groups=mutations.gatekeeper.sh,resources=*,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=status.gatekeeper.sh,resources=*,verbs=get;list;watch;create;update;patch;delete

// Reconcile reads that state of the cluster for a mutator object and makes changes based on the state read
// and what is in the mutator.Spec.
func (r *ReconcileMutatorStatus) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// Unrecoverable, do not retry.

// Sanity - make sure it is a mutator resource.

// Unrecoverable, do not retry.

// If the mutator does not exist, we are done

// Don't report status if it's not for the correct object. This can happen
// if a watch gets interrupted, causing the mutator status to be deleted
// out from underneath it

type sortableStatuses []v1beta1.MutatorPodStatus

func (s sortableStatuses) Len() int { _ = "STUB: not implemented"; return 0 }

func (s sortableStatuses) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (s sortableStatuses) Swap(i, j int) { _ = "STUB: not implemented"; return }
