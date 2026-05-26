package testutils

import (
	"context"
	"testing"

	"github.com/open-policy-agent/gatekeeper/v3/pkg/watch"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/manager"
)

// StartManager starts mgr. Registers a cleanup function to stop the manager at the completion of the test.
func StartManager(ctx context.Context, t *testing.T, mgr manager.Manager) {
	_ = "STUB: not implemented"
	return
}

// SetupManager sets up a controller-runtime manager with registered watch manager.
func SetupManager(t *testing.T, cfg *rest.Config) (manager.Manager, *watch.Manager) {
	_ = "STUB: not implemented"
	return *new(manager.Manager), nil
}
