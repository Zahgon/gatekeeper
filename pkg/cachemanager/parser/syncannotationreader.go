package parser

import (
	"github.com/open-policy-agent/frameworks/constraint/pkg/core/templates"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// syncAnnotationName is the name of the annotation that stores
// GVKS that are required to be synced.
const SyncAnnotationName = "metadata.gatekeeper.sh/requires-sync-data"

// SyncRequirements contains a list of ANDed requirements, each of which
// contains a GVK equivalence set.
type SyncRequirements []GVKEquivalenceSet

// GVKEquivalenceSet is a set of GVKs that a template can use
// interchangeably in its referential policy implementation.
type GVKEquivalenceSet map[schema.GroupVersionKind]struct{}

// CompactSyncRequirements contains a list of ANDed requirements, each of
// which contains a list of GVK clauses.
type CompactSyncRequirements [][]GVKClause

// GVKClause contains a set of equivalent GVKs, expressed
// in the form [groups, versions, kinds] where any combination of
// items from these three fields can be considered a valid option.
// Used for unmarshalling as this is the form used in requiressync annotations.
type GVKClause struct {
	Groups   []string `json:"groups"`
	Versions []string `json:"versions"`
	Kinds    []string `json:"kinds"`
}

// ReadSyncRequirements parses the sync requirements from a
// constraint template.
func ReadSyncRequirements(t *templates.ConstraintTemplate) (SyncRequirements, error) {
	_ = "STUB: not implemented"
	return *new(SyncRequirements), nil
}

// Takes a GVK Clause and expands it into a GVKEquivalenceSet (to be unioned
// with the GVKEquivalenceSet expansions of the other clauses).
func ExpandGVKClause(clause GVKClause) GVKEquivalenceSet {
	_ = "STUB: not implemented"
	return *new(GVKEquivalenceSet)
}

// Takes a CompactSyncRequirements (the json form provided in the template
// annotation) and expands it into a SyncRequirements.
func ExpandCompactRequirements(compactSyncRequirements CompactSyncRequirements) (SyncRequirements, error) {
	_ = "STUB: not implemented"
	return *new(SyncRequirements), nil
}

func (s GVKEquivalenceSet) String() string { _ = "STUB: not implemented"; return "" }

func (s SyncRequirements) String() string { _ = "STUB: not implemented"; return "" }
