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

package controller

import (
	"context"
	"flag"
	"sync"

	constraintclient "github.com/open-policy-agent/frameworks/constraint/pkg/client"
	"github.com/open-policy-agent/frameworks/constraint/pkg/externaldata"
	cm "github.com/open-policy-agent/gatekeeper/v3/pkg/cachemanager"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/controller/config/process"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/controller/webhookconfig/webhookconfigcache"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/expansion"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/export"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/mutation"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/readiness"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/watch"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/event"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

var (
	debugUseFakePod = flag.Bool("debug-use-fake-pod", false, "Use a fake pod name so the Gatekeeper executable can be run outside of Kubernetes")
	remoteCluster   = flag.Bool("enable-remote-cluster", false, "(alpha) Enable remote cluster mode where Gatekeeper operates against a target cluster specified via --kubeconfig while running in the local cluster. Mutually exclusive with --debug-use-fake-pod.")
)

type Injector interface {
	InjectTracker(tracker *readiness.Tracker)

	Add(mgr manager.Manager) error
}

type GetPodInjector interface {
	InjectGetPod(func(context.Context) (*corev1.Pod, error))
}

type ExportInjector interface {
	InjectExportSystem(exportSystem export.Exporter)
}

type DataClientInjector interface {
	InjectCFClient(*constraintclient.Client)
}

type WatchManagerInjector interface {
	InjectWatchManager(*watch.Manager)
}

type MutationSystemInjector interface {
	InjectMutationSystem(mutationSystem *mutation.System)
}

type ExpansionSystemInjector interface {
	InjectExpansionSystem(expansionSystem *expansion.System)
}

type ProviderCacheInjector interface {
	InjectProviderCache(providerCache *externaldata.ProviderCache)
}

type CacheManagerInjector interface {
	InjectCacheManager(cm *cm.CacheManager)
}

type GetProcessExcluderInjector interface {
	InjectProcessExcluder(processExcluder *process.Excluder)
}

type ConstraintTemplateEventInjector interface {
	InjectConstraintTemplateEvent(constraintTemplateEvents chan event.GenericEvent)
}

type WebhookConfigCacheInjector interface {
	InjectWebhookConfigCache(webhookConfigCache *webhookconfigcache.WebhookConfigCache)
}

// Injectors is a list of adder structs that need injection. We can convert this
// to an interface once we create controllers for things like data sync.
var Injectors []Injector

// AddToManagerFuncs is a list of functions to add all Controllers to the Manager.
var AddToManagerFuncs []func(manager.Manager) error

// Dependencies are dependencies that can be injected into controllers.
type Dependencies struct {
	CFClient           *constraintclient.Client
	WatchManger        *watch.Manager
	Tracker            *readiness.Tracker
	GetPod             func(context.Context) (*corev1.Pod, error)
	ProcessExcluder    *process.Excluder
	MutationSystem     *mutation.System
	ExpansionSystem    *expansion.System
	ProviderCache      *externaldata.ProviderCache
	ExportSystem       *export.System
	SyncEventsCh       chan event.GenericEvent
	CacheMgr           *cm.CacheManager
	CtEvents           chan event.GenericEvent
	WebhookConfigCache *webhookconfigcache.WebhookConfigCache
}

type defaultPodGetter struct {
	client client.Client
	scheme *runtime.Scheme
	pod    *corev1.Pod
	mux    sync.RWMutex
}

func (g *defaultPodGetter) GetPod(ctx context.Context) (*corev1.Pod, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// guard against the race condition where the pod has been retrieved
// between releasing the read lock and acquiring the write lock

// use unstructured to avoid inadvertently creating a watch on pods

// AddToManager adds all Controllers to the Manager.
func AddToManager(m manager.Manager, deps *Dependencies) error {
	_ = "STUB: not implemented"
	return nil
}

// In remote cluster mode, the pod doesn't exist in the target cluster. Skip setting OwnerReferences.

// In remote cluster mode, use InClusterConfig to connect to local cluster.

// Adding the CacheManager as a runnable;
// manager will start CacheManager.

// Create subordinate controller - we will feed it events dynamically via watch

// this is used by the config controller to sync
