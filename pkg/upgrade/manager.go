package upgrade

// TODO consider whether this needs to exist after https://github.com/kubernetes/kubernetes/pull/79495
// is merged, or we make the minimum supported version of k8s v1.14

import (
	"context"

	"github.com/open-policy-agent/gatekeeper/v3/pkg/util"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

var log = logf.Log.WithName("controller").WithValues("metaKind", "upgrade")

const (
	crdName = "constrainttemplates.templates.gatekeeper.sh"
)

// Manager allows us to upgrade resources on startup.
type Manager struct {
	client client.Client
	mgr    manager.Manager
}

// New creates a new manager for audit.
func New(mgr manager.Manager) *Manager { _ = "STUB: not implemented"; return nil }

// Start implements the Runnable interface.
func (um *Manager) Start(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// We must block indefinitely or manager will exit

func (um *Manager) ensureCRDExists(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

func (um *Manager) getAllKinds(groupVersion string) (*metav1.APIResourceList, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (um *Manager) upgrade(ctx context.Context) error { _ = "STUB: not implemented"; return nil }

// upgradeGroupVersion touches each resource in a given groupVersion, incrementing its storage version.
func (um *Manager) upgradeGroupVersion(ctx context.Context, groupVersion string) error {
	_ = "STUB: not implemented"
	// new client to get updated restmapper
	return nil
}

// get all resource kinds

// If the resource doesn't exist, it doesn't need upgrading

// For some reason we have seen duplicate kinds, suppress that

// get resource for each Kind

// get each resource

type updateResourceLoop struct {
	ur      map[util.KindVersionName]unstructured.Unstructured
	client  client.Client
	stop    chan struct{}
	stopped chan struct{}
}

func (urloop *updateResourceLoop) update(ctx context.Context) { _ = "STUB: not implemented"; return }

// get the latest constraint
