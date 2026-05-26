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

package sync

import (
	"context"

	"github.com/go-logr/logr"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/cachemanager"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/syncutil"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var log = logf.Log.WithName("controller").WithValues("metaKind", "Sync")

type Adder struct {
	CacheManager *cachemanager.CacheManager
	Events       <-chan event.GenericEvent
}

// Add creates a new Sync Controller and adds it to the Manager with default RBAC. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func (a *Adder) Add(mgr manager.Manager) error { _ = "STUB: not implemented"; return nil }

// newReconciler returns a new reconcile.Reconciler.
func newReconciler(
	mgr manager.Manager,
	reporter *syncutil.Reporter,
	cm *cachemanager.CacheManager,
) reconcile.Reconciler {
	_ = "STUB: not implemented"
	return *new(reconcile.Reconciler)
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler.
func add(mgr manager.Manager, r reconcile.Reconciler, events <-chan event.GenericEvent) error {
	_ = "STUB: not implemented"
	// Create a new controller
	return nil
}

// Watch for changes to the provided resource

var _ reconcile.Reconciler = &ReconcileSync{}

// ReconcileSync reconciles an arbitrary object described by Kind.
type ReconcileSync struct {
	reader client.Reader

	scheme   *runtime.Scheme
	log      logr.Logger
	reporter *syncutil.Reporter
	cm       *cachemanager.CacheManager
}

// +kubebuilder:rbac:groups=constraints.gatekeeper.sh,resources=*,verbs=get;list;watch;create;update;patch;delete

// Reconcile reads that state of the cluster for an object and makes changes based on the state read
// and what is in the constraint.Spec.
func (r *ReconcileSync) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// Unrecoverable, do not retry.
// TODO(OREN) add metric

// This is a deletion; remove the data

// Error reading the object - requeue the request.
