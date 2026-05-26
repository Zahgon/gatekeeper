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

package constrainttemplatestatus

import (
	"context"

	"github.com/go-logr/logr"
	constraintclient "github.com/open-policy-agent/frameworks/constraint/pkg/client"
	"github.com/open-policy-agent/gatekeeper/v3/apis/status/v1beta1"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/logging"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/watch"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/handler"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var log = logf.Log.WithName("controller").WithValues(logging.Process, "constraint_template_status_controller")

type Adder struct {
	CfClient     *constraintclient.Client
	WatchManager *watch.Manager
}

// Add creates a new Constraint Status Controller and adds it to the Manager. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func (a *Adder) Add(mgr manager.Manager) error { _ = "STUB: not implemented"; return nil }

// newReconciler returns a new reconcile.Reconciler.
func newReconciler(
	mgr manager.Manager,
) reconcile.Reconciler {
	_ = "STUB: not implemented"
	return *new(reconcile.Reconciler)
}

// Separate reader and writer because manager's default client bypasses the cache for unstructured resources.

// PodStatusToConstraintTemplateMapper correlates a ConstraintTemplatePodStatus with its corresponding constraint template
// `selfOnly` tells the mapper to only map statuses corresponding to the current pod.
func PodStatusToConstraintTemplateMapper(selfOnly bool) handler.TypedMapFunc[*v1beta1.ConstraintTemplatePodStatus, reconcile.Request] {
	_ = "STUB: not implemented"
	return nil
}

// Do not attempt to reconcile the resource when other pods have changed their status

// add adds a new Controller to mgr with r as the reconcile.Reconciler.
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	_ = "STUB: not implemented"
	// Create a new controller
	return nil
}

// Watch for changes to ConstraintTemplateStatus

// Watch for changes to the provided constraint
// Watch for changes to ConstraintTemplate

var _ reconcile.Reconciler = &ReconcileConstraintStatus{}

// ReconcileConstraintStatus reconciles an arbitrary constraint object described by Kind.
type ReconcileConstraintStatus struct {
	reader       client.Reader
	writer       client.Writer
	statusClient client.StatusClient
	scheme       *runtime.Scheme
	log          logr.Logger
}

// +kubebuilder:rbac:groups=constraints.gatekeeper.sh,resources=*,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=status.gatekeeper.sh,resources=*,verbs=get;list;watch;create;update;patch;delete

// Reconcile reads that state of the cluster for a constraint object and makes changes based on the state read
// and what is in the constraint.Spec.
func (r *ReconcileConstraintStatus) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// If the template does not exist, we are done

// created is true if at least one Pod hasn't reported any errors

// Don't report status if it's not for the correct object. This can happen
// if a watch gets interrupted, causing the constraint status to be deleted
// out from underneath it

type sortableStatuses []v1beta1.ConstraintTemplatePodStatus

func (s sortableStatuses) Len() int { _ = "STUB: not implemented"; return 0 }

func (s sortableStatuses) Less(i, j int) bool { _ = "STUB: not implemented"; return false }

func (s sortableStatuses) Swap(i, j int) { _ = "STUB: not implemented"; return }
