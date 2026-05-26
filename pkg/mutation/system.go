package mutation

import (
	"context"
	"sync"

	"github.com/google/uuid"
	"github.com/open-policy-agent/frameworks/constraint/pkg/externaldata"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/mutation/schema"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/mutation/types"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/certwatcher"
)

// ErrNotConverging reports that applying all Mutators isn't converging.
var ErrNotConverging = errors.New("mutation not converging")

// ErrNotRemoved reports that we were unable to remove a Mutator properly as
// System was in an inconsistent state.
var ErrNotRemoved = errors.New("failed to find mutator on sorted list")

// System keeps the list of mutators and provides an interface to apply mutations.
type System struct {
	schemaDB                          schema.DB
	orderedMutators                   orderedIDs
	mutatorsMap                       map[types.ID]types.Mutator
	mux                               sync.RWMutex
	reporter                          StatsReporter
	newUUID                           func() uuid.UUID
	providerCache                     *externaldata.ProviderCache
	sendRequestToExternalDataProvider externaldata.SendRequestToProvider
	clientCertWatcher                 *certwatcher.CertWatcher
}

// SystemOpts allows for optional dependencies to be passed into the mutation System.
type SystemOpts struct {
	Reporter                          StatsReporter
	NewUUID                           func() uuid.UUID
	ProviderCache                     *externaldata.ProviderCache
	SendRequestToExternalDataProvider externaldata.SendRequestToProvider
	ClientCertWatcher                 *certwatcher.CertWatcher
}

// NewSystem initializes an empty mutation system.
func NewSystem(options SystemOpts) *System { _ = "STUB: not implemented"; return nil }

// Get mutator for given id.
func (s *System) Get(id types.ID) types.Mutator {
	_ = "STUB: not implemented"
	return *new(types.Mutator)
}

// Upsert updates or inserts the given object. Returns an error in case of
// schema conflicts.
func (s *System) Upsert(m types.Mutator) error { _ = "STUB: not implemented"; return nil }

// Handle the case where a previous reconcile successfully updated System,
// but the update to PodStatus failed.

// Check schema consistency only if the mutator has schema.

// This means the error is not due to a schema conflict, and is most likely
// a bug.

// Remove removes the mutator from the mutation system.
func (s *System) Remove(id types.ID) error { _ = "STUB: not implemented"; return nil }

func (s *System) GetConflicts(id types.ID) map[types.ID]bool { _ = "STUB: not implemented"; return nil }

// Mutate applies the mutation in place to the given object. Returns
// true if applying Mutators caused any changes to the object.
func (s *System) Mutate(ctx context.Context, mutable *types.Mutable) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

// mutate runs all Mutators on obj. Returns the number of iterations required
// to converge, and any error encountered attempting to run Mutators.
func (s *System) mutate(ctx context.Context, mutable *types.Mutable) (int, error) {
	_ = "STUB: not implemented"
	return 0, nil
}

// Don't try to apply Mutators which have conflicts.

// If no mutations were applied, we can safely assume the object is
// identical to before.

func mutateErr(err error, uid uuid.UUID, mID types.ID, obj *unstructured.Unstructured) error {
	_ = "STUB: not implemented"
	return nil
}

func matchesErr(err error, mID types.ID, obj *unstructured.Unstructured) error {
	_ = "STUB: not implemented"
	return nil
}
