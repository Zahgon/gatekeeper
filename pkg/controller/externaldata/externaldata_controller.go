package externaldata

import (
	"context"

	externaldatav1beta1 "github.com/open-policy-agent/frameworks/constraint/pkg/apis/externaldata/v1beta1"
	constraintclient "github.com/open-policy-agent/frameworks/constraint/pkg/client"
	frameworksexternaldata "github.com/open-policy-agent/frameworks/constraint/pkg/externaldata"
	statusv1beta1 "github.com/open-policy-agent/gatekeeper/v3/apis/status/v1beta1"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/logging"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/readiness"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/runtime/schema"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var (
	log = logf.Log.WithName("controller").WithValues(logging.Process, "externaldata_controller")

	gvkExternalData = schema.GroupVersionKind{
		Group:   "externaldata.gatekeeper.sh",
		Version: "v1beta1",
		Kind:    "Provider",
	}
)

type Adder struct {
	CFClient      *constraintclient.Client
	ProviderCache *frameworksexternaldata.ProviderCache
	Tracker       *readiness.Tracker
	// GetPod returns an instance of the currently running Gatekeeper pod
	GetPod func(context.Context) (*corev1.Pod, error)
}

func (a *Adder) InjectCFClient(c *constraintclient.Client) { _ = "STUB: not implemented"; return }

func (a *Adder) InjectTracker(t *readiness.Tracker) { _ = "STUB: not implemented"; return }

func (a *Adder) InjectProviderCache(providerCache *frameworksexternaldata.ProviderCache) {
	_ = "STUB: not implemented"
	return
}

func (a *Adder) InjectGetPod(getPod func(ctx context.Context) (*corev1.Pod, error)) {
	_ = "STUB: not implemented"

	// Add creates a new ExternalData Controller and adds it to the Manager. The Manager will set fields on the Controller
	// and Start it when the Manager is Started.
	return
}

func (a *Adder) Add(mgr manager.Manager) error { _ = "STUB: not implemented"; return nil }

// Reconciler reconciles a ExternalData object.
type Reconciler struct {
	client.Client
	cfClient      *constraintclient.Client
	providerCache *frameworksexternaldata.ProviderCache
	tracker       *readiness.Tracker
	scheme        *runtime.Scheme
	metrics       *reporter

	getPod func(context.Context) (*corev1.Pod, error)
}

// newReconciler returns a new reconcile.Reconciler.
func newReconciler(mgr manager.Manager, client *constraintclient.Client, providerCache *frameworksexternaldata.ProviderCache, tracker *readiness.Tracker, getPod func(ctx context.Context) (*corev1.Pod, error)) *Reconciler {
	_ = "STUB: not implemented"
	return nil
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler.
func add(mgr manager.Manager, r reconcile.Reconciler) error { _ = "STUB: not implemented"; return nil }

// Create a new controller

// Watch for changes to Provider

func (r *Reconciler) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

func (r *Reconciler) updateOrCreatePodStatus(ctx context.Context, provider *externaldatav1beta1.Provider, providerErrors []*statusv1beta1.ProviderError) error {
	_ = "STUB: not implemented"
	return nil
}

// Check if it exists already

func (r *Reconciler) newProviderStatus(pod *corev1.Pod, provider *externaldatav1beta1.Provider) (*statusv1beta1.ProviderPodStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Reconciler) deleteStatus(ctx context.Context, providerName string) error {
	_ = "STUB: not implemented"
	return nil
}

func setStatus(status *statusv1beta1.ProviderPodStatus, providerErrors []*statusv1beta1.ProviderError, updateLastCacheTime bool) {
	_ = "STUB: not implemented"
	return
}

func errorChanged(oldErrors, newErrors []*statusv1beta1.ProviderError) bool {
	_ = "STUB: not implemented"
	return false
}

// Check errors without considering order

// If any old errors remain, they weren't found in new errors
