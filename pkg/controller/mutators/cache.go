package mutators

import (
	"sync"

	"github.com/open-policy-agent/gatekeeper/v3/pkg/mutation/types"
)

type mutatorStatus struct {
	ingestion MutatorIngestionStatus
	conflict  bool
}

type Cache struct {
	cache map[types.ID]mutatorStatus
	mux   sync.RWMutex
}

func NewMutationCache() *Cache { _ = "STUB: not implemented"; return nil }

func (c *Cache) Upsert(mID types.ID, ingestionStatus MutatorIngestionStatus, conflict bool) {
	_ = "STUB: not implemented"
	return
}

func (c *Cache) Remove(mID types.ID) { _ = "STUB: not implemented"; return }

// TallyStatus calculates the number of mutators in each of the available
// MutatorIngestionStatus states and returns a map of those states and the
// count for each.
func (c *Cache) TallyStatus() map[MutatorIngestionStatus]int { _ = "STUB: not implemented"; return nil }

// TallyConflict calculates and returns the number of mutators that are
// currently in a conflict state as maintained in the Cache.
func (c *Cache) TallyConflict() int { _ = "STUB: not implemented"; return 0 }
