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

package configstatus

import (
	"context"

	"github.com/go-logr/logr"
	"github.com/open-policy-agent/gatekeeper/v3/apis/status/v1beta1"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/logging"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/readiness"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/watch"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var log = logf.Log.WithName("controller").WithValues(logging.Process, "config_status_controller")

type Adder struct {
	WatchManager *watch.Manager
}

func (a *Adder) InjectTracker(_ *readiness.Tracker) {
	_ = "STUB: not implemented"

	// Add creates a new config Status Controller and adds it to the Manager. The Manager will set fields on the Controller
	// and Start it when the Manager is Started.
	return
}

func (a *Adder) Add(mgr manager.Manager) error { _ = "STUB: not implemented"; return nil }

// newReconciler returns a new reconcile.Reconciler.
func newReconciler(mgr manager.Manager) reconcile.Reconciler {
	_ = "STUB: not implemented"
	return *new(reconcile.Reconciler)
}

// Separate reader and writer because manager's default client bypasses the cache for unstructured resources.

// PodStatusToConfigMapper correlates a ConfigPodStatus with its corresponding Config.
// `selfOnly` tells the mapper to only map statuses corresponding to the current pod.
func PodStatusToConfigMapper(selfOnly bool) handler.TypedMapFunc[*v1beta1.ConfigPodStatus, reconcile.Request] {
	_ = "STUB: not implemented"
	return nil
}

// Do not attempt to reconcile the resource when other pods have changed their status

// Add creates a new config status Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func add(mgr manager.Manager, r reconcile.Reconciler) error { _ = "STUB: not implemented"; return nil }

var _ reconcile.Reconciler = &ReconcileConfigStatus{}

// ReconcileConfigStatus provides the dependencies required to reconcile
// the status of a Config resource.
type ReconcileConfigStatus struct {
	reader       client.Reader
	writer       client.Writer
	statusClient client.StatusClient

	scheme *runtime.Scheme
	log    logr.Logger
}

// +kubebuilder:rbac:groups=config.gatekeeper.sh,resources=*,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=status.gatekeeper.sh,resources=*,verbs=get;list;watch;create;update;patch;delete

// Reconcile reads that state of the cluster for a config object and makes changes based on the state read
// and what is in the constraint.Spec.
func (r *ReconcileConfigStatus) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// If the Config does not exist then we are done

// Don't report status if it's not for the correct object. This can happen
// if a watch gets interrupted, causing the constraint status to be deleted
// out from underneath it

type sortableStatuses []v1beta1.ConfigPodStatus

func (s sortableStatuses) Len() int { _ = "STUB: not implemented"; return 0 }

func (s sortableStatuses) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (s sortableStatuses) Swap(i, j int) { _ = "STUB: not implemented"; return }
