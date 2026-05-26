package export

import (
	"context"

	connectionv1alpha1 "github.com/open-policy-agent/gatekeeper/v3/apis/connection/v1alpha1"
	statusv1alpha1 "github.com/open-policy-agent/gatekeeper/v3/apis/status/v1alpha1"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/export"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/logging"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/readiness"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var log = logf.Log.WithName("controller").WithValues(logging.Process, "export_controller")

type Adder struct {
	ExportSystem export.Exporter
	// GetPod returns an instance of the currently running Gatekeeper pod
	GetPod func(context.Context) (*corev1.Pod, error)
}

func (a *Adder) Add(mgr manager.Manager) error { _ = "STUB: not implemented"; return nil }

func (a *Adder) InjectTracker(_ *readiness.Tracker) { _ = "STUB: not implemented"; return }

func (a *Adder) InjectExportSystem(exportSystem export.Exporter) { _ = "STUB: not implemented"; return }

func (a *Adder) InjectGetPod(getPod func(ctx context.Context) (*corev1.Pod, error)) {
	_ = "STUB: not implemented"
	return
}

type Reconciler struct {
	reader client.Reader
	writer client.Writer
	scheme *runtime.Scheme
	system export.Exporter
	// TODO: Refactor this once multiple connections are supported, for now this helps with injecting dependency for tests
	auditConnectionName string
	getPod              func(context.Context) (*corev1.Pod, error)
}

func newReconciler(mgr manager.Manager, system export.Exporter, auditConnectionName string, getPod func(context.Context) (*corev1.Pod, error)) *Reconciler {
	_ = "STUB: not implemented"
	return nil
}

func add(mgr manager.Manager, r reconcile.Reconciler) error { _ = "STUB: not implemented"; return nil }

// +kubebuilder:rbac:groups=connection.gatekeeper.sh,resources=*,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=status.gatekeeper.sh,resources=*,verbs=get;list;watch;create;update;patch;delete
func (r *Reconciler) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	return *new(reconcile.Result), nil
}

// Reset the active connection status to false if UpsertConnection fails

func UpdateOrCreateConnectionPodStatus(
	ctx context.Context,
	reader client.Reader,
	writer client.Writer,
	scheme *runtime.Scheme,
	connObjName string,
	exportErrors []*statusv1alpha1.ConnectionError,
	activeConnection *bool,
	getPod func(context.Context) (*corev1.Pod, error),
) error {
	_ = "STUB: not implemented"
	// Since the caller from Audit won't have an incoming request
	// use the connection name from the audit connection flag as the predetermined connection name
	return nil
}

func updateOrCreateConnectionPodStatus(ctx context.Context,
	reader client.Reader,
	writer client.Writer,
	scheme *runtime.Scheme,
	connObj *connectionv1alpha1.Connection,
	exportErrors []*statusv1alpha1.ConnectionError,
	activeConnection *bool,
	getPod func(context.Context) (*corev1.Pod, error),
) error {
	_ = "STUB: not implemented"
	return nil
}

// Check if it exists already

// ConnectionPodStatus object exists so get the existing active state

// nil indicates expected active Connection state is unknown by caller during Upsert

// Reset the active connection state when there are updates to the Connection object to ensure the active state is only true when the Publish succeeds for the current Connection

// Trust the existing object when the Connection hasn't change - since active can only be true when Publish succeeds, we don't want to potentially reset active state between every Audit causing thrashing

// ObservedGeneration is used to track the generation of the Connection object

func deleteStatus(ctx context.Context,
	writer client.Writer,
	connectionNamespace string,
	connectionName string,
	getPod func(context.Context) (*corev1.Pod, error),
) error {
	_ = "STUB: not implemented"
	return nil
}

func newConnectionPodStatus(scheme *runtime.Scheme,
	pod *corev1.Pod,
	connObj *connectionv1alpha1.Connection,
) (*statusv1alpha1.ConnectionPodStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func setStatusErrors(
	connPodStatusObj *statusv1alpha1.ConnectionPodStatus,
	exportErrors []*statusv1alpha1.ConnectionError,
) {
	_ = "STUB: not implemented"
	return
}
