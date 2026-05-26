package pruner

import (
	"context"
	"time"

	"github.com/open-policy-agent/gatekeeper/v3/pkg/cachemanager"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/readiness"
)

const tickDuration = 3 * time.Second

// ExpectationsPruner polls the ReadyTracker and other data sources in Gatekeeper to remove
// un-satisfiable expectations in the RT that would incorrectly block startup.
type ExpectationsPruner struct {
	cacheMgr *cachemanager.CacheManager
	tracker  *readiness.Tracker
}

func NewExpectationsPruner(cm *cachemanager.CacheManager, rt *readiness.Tracker) *ExpectationsPruner {
	_ = "STUB: not implemented"
	return nil
}

func (e *ExpectationsPruner) Start(ctx context.Context) error {
	_ = "STUB: not implemented"
	return nil
}

// we're done, there's no need to
// further manage the data sync expectations.

// pruneUnwatchedGVKs prunes data expectations that are no longer correct based on the up-to-date
// information in the CacheManager.
func (e *ExpectationsPruner) pruneUnwatchedGVKs() { _ = "STUB: not implemented"; return }
