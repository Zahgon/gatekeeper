package testutils

import (
	"context"
	"sync"
	"testing"
	"time"

	constraintclient "github.com/open-policy-agent/frameworks/constraint/pkg/client"
	"k8s.io/apimachinery/pkg/util/wait"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
)

var gkCRDPath = []string{"config", "crd", "bases"}

// ConstantRetry makes 3,000 attempts at a rate of 100 per second. Since this
// is a test instance and not a "real" cluster, this is fine and there's no need
// to increase the wait time each iteration.
var ConstantRetry = wait.Backoff{
	Steps:    3000,
	Duration: 10 * time.Millisecond,
}

// CreateGatekeeperNamespace bootstraps the gatekeeper-system namespace for use in tests.
func CreateGatekeeperNamespace(cfg *rest.Config) error { _ = "STUB: not implemented"; return nil }

// Create gatekeeper namespace

// DeleteObjectAndConfirm returns a callback which deletes obj from the passed
// Client. Does result in mutations to obj. The callback includes a cached copy
// of all information required to delete obj in the callback, so it is safe to
// mutate obj afterwards. Similarly - client.Delete mutates its input, but
// the callback does not call client.Delete on obj. Instead, it creates a
// single-purpose Unstructured for this purpose. Thus, obj is not mutated after
// the callback is run.
func DeleteObjectAndConfirm(ctx context.Context, t *testing.T, c client.Client, obj client.Object) func() {
	_ = "STUB: not implemented"

	// Cache the identifying information from obj. We refer to this cached
	// information in the callback, and not obj itself.
	return nil
}

// We can't send a proper delete request with an Unstructured without
// filling in GVK. The alternative would be to require tests to construct
// a valid Scheme or provide a factory method for the type to delete - this
// is easier.

// Construct a single-use Unstructured to send the Delete request.

// Construct a single-use Unstructured to send the Get request. It isn't
// safe to reuse Unstructureds for each retry as Get modifies its input.

// Marshal the currently-gotten object, so it can be printed in test
// failure output.

func StartControlPlane(m *testing.M, cfg **rest.Config, testerDepth int) {
	_ = "STUB: not implemented"
	return
}

///TODO(ritazh): remove when vap is GAed in k/k

// CreateThenCleanup creates obj in Client, and then registers obj to be deleted
// at the end of the test. The passed obj is safely deepcopied before being
// passed to client.Create, so it is not mutated by this call.
func CreateThenCleanup(ctx context.Context, t *testing.T, c client.Client, obj client.Object) {
	_ = "STUB: not implemented"
	return
}

// It is unnecessary to deepcopy obj as deleteObjectAndConfirm does not pass
// obj to any Client calls.

func SetupDataClient(t *testing.T) *constraintclient.Client { _ = "STUB: not implemented"; return nil }

// SetupTestReconcile returns a reconcile.Reconcile implementation that delegates to inner and
// writes the request to requests after Reconcile is finished.
func SetupTestReconcile(inner reconcile.Reconciler) (reconcile.Reconciler, *sync.Map) {
	_ = "STUB: not implemented"
	return *new(reconcile.Reconciler), nil
}
