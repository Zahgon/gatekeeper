package expansion

import (
	"context"

	"github.com/open-policy-agent/gatekeeper/v3/apis/expansion/unversioned"
	statusv1beta1 "github.com/open-policy-agent/gatekeeper/v3/apis/status/v1beta1"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/expansion"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/logging"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/readiness"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/watch"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sigs.k8s.io/controller-runtime/pkg/source"
)

var log = logf.Log.WithName("controller").WithValues("kind", "ExpansionTemplate", logging.Process, "template_expansion_controller")

// eventQueueSize is how many events to queue before blocking.
const eventQueueSize = 1024

type Adder struct {
	WatchManager    *watch.Manager
	ExpansionSystem *expansion.System
	Tracker         *readiness.Tracker
	// GetPod returns an instance of the currently running Gatekeeper pod
	GetPod func(context.Context) (*corev1.Pod, error)
}

func (a *Adder) Add(mgr manager.Manager) error { _ = "STUB: not implemented"; return nil }

func (a *Adder) InjectTracker(tracker *readiness.Tracker) { _ = "STUB: not implemented"; return }

func (a *Adder) InjectExpansionSystem(expansionSystem *expansion.System) {
	_ = "STUB: not implemented"
	return
}

func (a *Adder) InjectGetPod(getPod func(ctx context.Context) (*corev1.Pod, error)) {
	_ = "STUB: not implemented"
	return
}

type Reconciler struct {
	client.Client
	system       *expansion.System
	scheme       *runtime.Scheme
	registry     *etRegistry
	statusClient client.StatusClient
	tracker      *readiness.Tracker
	events       chan event.GenericEvent
	eventSource  source.Source

	getPod func(context.Context) (*corev1.Pod, error)
}

func newReconciler(mgr manager.Manager,
	system *expansion.System,
	getPod func(ctx context.Context) (*corev1.Pod, error),
	tracker *readiness.Tracker,
) *Reconciler {
	_ = "STUB: not implemented"
	return nil
}

func add(mgr manager.Manager, r *Reconciler) error { _ = "STUB: not implemented"; return nil }

// Watch for enqueued events

// Watch for changes to ExpansionTemplates

// Watch for changes to ExpansionTemplateStatuses

func (r *Reconciler) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// et will be an empty struct. We set the metadata name, which is
// used as a key to delete it from the expansion system

func (r *Reconciler) queueConflicts(old expansion.IDSet) { _ = "STUB: not implemented"; return }

// ExpansionTemplate is cluster-scoped, so we do not set namespace

func symmetricDiff(x, y expansion.IDSet) expansion.IDSet {
	_ = "STUB: not implemented"
	return *new(expansion.IDSet)
}

func (r *Reconciler) deleteStatus(ctx context.Context, etName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *Reconciler) updateOrCreatePodStatus(ctx context.Context, et *unversioned.ExpansionTemplate, etErr error) error {
	_ = "STUB: not implemented"
	return nil
}

// Check if it exists already

func (r *Reconciler) newETStatus(pod *corev1.Pod, et *unversioned.ExpansionTemplate) (*statusv1beta1.ExpansionTemplatePodStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (r *Reconciler) getTracker() readiness.Expectations {
	_ = "STUB: not implemented"
	return *new(readiness.Expectations)
}

func setStatusError(status *statusv1beta1.ExpansionTemplatePodStatus, etErr error) {
	_ = "STUB: not implemented"
	return
}
