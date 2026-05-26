package mutation

import (
	"github.com/open-policy-agent/gatekeeper/v3/pkg/mutation/types"
)

type orderedIDs struct {
	ids []types.ID
}

func (o *orderedIDs) insert(id types.ID) { _ = "STUB: not implemented"; return }

// Add to the end of the list.

// remove removes the Mutator with id. Returns true if the Mutator was removed,
// or false if the Mutator was not found.
func (o *orderedIDs) remove(id types.ID) bool { _ = "STUB: not implemented"; return false }

// The map is expected to be in sync with the list, so if we don't find it
// we return an error.

func (o *orderedIDs) find(id types.ID) (int, bool) { _ = "STUB: not implemented"; return 0, false }

func greaterOrEqual(id1, id2 types.ID) bool { _ = "STUB: not implemented"; return false }
