package syncset

import (
	"context"

	syncsetv1alpha1 "github.com/open-policy-agent/gatekeeper/v3/apis/syncset/v1alpha1"
	cm "github.com/open-policy-agent/gatekeeper/v3/pkg/cachemanager"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/logging"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/readiness"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

const (
	ctrlName = "syncset-controller"
)

var (
	log = logf.Log.WithName("controller").WithValues("kind", "SyncSet", logging.Process, "syncset_controller")

	syncsetGVK = syncsetv1alpha1.GroupVersion.WithKind("SyncSet")
)

type Adder struct {
	CacheManager *cm.CacheManager
	Tracker      *readiness.Tracker
}

// Add creates a new controller for SyncSets and adds it to the Manager.
func (a *Adder) Add(mgr manager.Manager) error { _ = "STUB: not implemented"; return nil }

func (a *Adder) InjectCacheManager(o *cm.CacheManager) { _ = "STUB: not implemented"; return }

func (a *Adder) InjectTracker(t *readiness.Tracker) { _ = "STUB: not implemented"; return }

func newReconciler(mgr manager.Manager, cm *cm.CacheManager, tracker *readiness.Tracker) (*ReconcileSyncSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func add(mgr manager.Manager, r reconcile.Reconciler) error { _ = "STUB: not implemented"; return nil }

var _ reconcile.Reconciler = &ReconcileSyncSet{}

// ReconcileSyncSet reconciles a SyncSet object.
type ReconcileSyncSet struct {
	reader client.Reader

	scheme       *runtime.Scheme
	cacheManager *cm.CacheManager
	tracker      *readiness.Tracker
}

func (r *ReconcileSyncSet) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// Error reading the object - requeue the request.

// Directly accessing the NamespaceName.String(), as NamespaceName is embedded within reconcile.Request.
