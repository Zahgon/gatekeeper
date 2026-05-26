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

package config

import (
	"context"
	"sync"

	"github.com/open-policy-agent/frameworks/constraint/pkg/apis/templates/v1beta1"
	configv1alpha1 "github.com/open-policy-agent/gatekeeper/v3/apis/config/v1alpha1"
	statusv1beta1 "github.com/open-policy-agent/gatekeeper/v3/apis/status/v1beta1"
	cm "github.com/open-policy-agent/gatekeeper/v3/pkg/cachemanager"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/readiness"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

const (
	ctrlName = "config-controller"
)

var (
	log       = logf.Log.WithName("controller").WithValues("kind", "Config")
	configGVK = configv1alpha1.GroupVersion.WithKind("Config")
)

func (r *ReconcileConfig) markDirtyTemplate(template *v1beta1.ConstraintTemplate) {
	_ = "STUB: not implemented"
	return
}

func (r *ReconcileConfig) getDirtyTemplatesAndClear() []*v1beta1.ConstraintTemplate {
	_ = "STUB: not implemented"
	return nil
}

func (r *ReconcileConfig) triggerConstraintTemplateReconciliation(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *ReconcileConfig) triggerDirtyTemplateReconciliation(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *ReconcileConfig) sendEventWithRetry(ctx context.Context, template *v1beta1.ConstraintTemplate) error {
	_ = "STUB: not implemented"
	return nil
}

type ChannelFullError struct{}

func (e *ChannelFullError) Error() string { _ = "STUB: not implemented"; return "" }

func (e *ChannelFullError) Temporary() bool { _ = "STUB: not implemented"; return false }

type Adder struct {
	Tracker      *readiness.Tracker
	CacheManager *cm.CacheManager
	CtEvents     chan<- event.GenericEvent
	// GetPod returns an instance of the currently running Gatekeeper pod
	GetPod func(context.Context) (*corev1.Pod, error)
}

// Add creates a new ConfigController and adds it to the Manager with default RBAC. The Manager will set fields on the Controller
// and Start it when the Manager is Started.
func (a *Adder) Add(mgr manager.Manager) error { _ = "STUB: not implemented"; return nil }

func (a *Adder) InjectTracker(t *readiness.Tracker) { _ = "STUB: not implemented"; return }

func (a *Adder) InjectCacheManager(cm *cm.CacheManager) { _ = "STUB: not implemented"; return }

func (a *Adder) InjectGetPod(getPod func(ctx context.Context) (*corev1.Pod, error)) {
	_ = "STUB: not implemented"
	return
}

func (a *Adder) InjectConstraintTemplateEvent(ctEvents chan event.GenericEvent) {
	_ = "STUB: not implemented"
	return

	// newReconciler returns a new reconcile.Reconciler.
}

func newReconciler(mgr manager.Manager, cm *cm.CacheManager, tracker *readiness.Tracker, getPod func(context.Context) (*corev1.Pod, error), ctEvents chan<- event.GenericEvent) (*ReconcileConfig, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// add adds a new Controller to mgr with r as the reconcile.Reconciler.
func add(mgr manager.Manager, r reconcile.Reconciler) error {
	_ = "STUB: not implemented"
	// Create a new controller
	return nil
}

// Watch for changes to Config

var _ reconcile.Reconciler = &ReconcileConfig{}

// ReconcileConfig reconciles a Config object.
type ReconcileConfig struct {
	reader       client.Reader
	writer       client.Writer
	statusClient client.StatusClient

	scheme       *runtime.Scheme
	cacheManager *cm.CacheManager

	tracker *readiness.Tracker

	getPod func(context.Context) (*corev1.Pod, error)

	ctEvents chan<- event.GenericEvent

	dirtyMu        sync.Mutex
	dirtyTemplates map[string]*v1beta1.ConstraintTemplate
}

// +kubebuilder:rbac:groups=*,resources=*,verbs=get;list;watch
// +kubebuilder:rbac:groups=config.gatekeeper.sh,resources=configs,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=config.gatekeeper.sh,resources=configs/status,verbs=get;update;patch
// +kubebuilder:rbac:groups="",resources=events,verbs=create;patch;

// Reconcile reads that state of the cluster for a Config object and makes changes based on the state read
// and what is in the Config.Spec
// Automatically generate RBAC rules to allow the Controller to read all things (for sync).
func (r *ReconcileConfig) Reconcile(ctx context.Context, request reconcile.Request) (reconcile.Result, error) {
	_ = "STUB: not implemented"
	// Fetch the Config instance
	return *new(reconcile.Result), nil
}

// if config is not found, we should remove cached data

// Error reading the object - requeue the request.

// If the config is being deleted the user is saying they don't want to
// sync anything

// K8s API conventions consider an object to be deleted when either the object no longer exists or when a deletion timestamp has been set.

// Enable verbose readiness stats if requested.

// Directly accessing the NamespaceName.String(), as NamespaceName is embedded within reconcile.Request.

func (r *ReconcileConfig) deleteStatus(ctx context.Context, cfgNamespace string, cfgName string) error {
	_ = "STUB: not implemented"
	return nil
}

func (r *ReconcileConfig) updateOrCreatePodStatus(ctx context.Context, cfg *configv1alpha1.Config, upsertErr error) error {
	_ = "STUB: not implemented"
	return nil
}

// Check if it exists already

func (r *ReconcileConfig) newConfigStatus(pod *corev1.Pod, cfg *configv1alpha1.Config) (*statusv1beta1.ConfigPodStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func setStatusError(status *statusv1beta1.ConfigPodStatus, etErr error) {
	_ = "STUB: not implemented"
	return
}
