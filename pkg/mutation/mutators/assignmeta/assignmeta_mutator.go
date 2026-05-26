package assignmeta

import (
	mutationsunversioned "github.com/open-policy-agent/gatekeeper/v3/apis/mutations/unversioned"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/logging"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/mutation/path/parser"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/mutation/path/tester"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/mutation/types"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

var log = logf.Log.WithName("mutation").WithValues(logging.Process, "mutation", logging.Mutator, "assignmeta")

var (
	labelsValidSubPath = []parser.Node{
		&parser.Object{
			Reference: "metadata",
		},
		&parser.Object{
			Reference: "labels",
		},
	}

	annotationValidSubPath = []parser.Node{
		&parser.Object{
			Reference: "metadata",
		},
		&parser.Object{
			Reference: "annotations",
		},
	}
)

// Mutator is a mutator built out of an
// AssignMeta instance.
type Mutator struct {
	id             types.ID
	assignMetadata *mutationsunversioned.AssignMetadata

	path parser.Path

	tester *tester.Tester
}

// Mutator implements mutator.
var _ types.Mutator = &Mutator{}

func (m *Mutator) Matches(mutable *types.Mutable) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (m *Mutator) Mutate(mutable *types.Mutable) (bool, error) {
	_ = "STUB: not implemented"
	// Note: Performance here can be improved by ~3x by writing a specialized
	// function instead of using a generic function. AssignMetadata only ever
	// mutates metadata.annotations or metadata.labels, and we spend ~70% of
	// compute covering cases that aren't valid for this Mutator.
	return false, nil
}

func (m *Mutator) MustTerminate() bool { _ = "STUB: not implemented"; return false }

func (m *Mutator) ID() types.ID { _ = "STUB: not implemented"; return *new(types.ID) }

func (m *Mutator) Path() parser.Path { _ = "STUB: not implemented"; return *new(parser.Path) }

func (m *Mutator) HasDiff(mutator types.Mutator) bool { _ = "STUB: not implemented"; return false }

// different types, different

// any difference in spec may be enough

func (m *Mutator) DeepCopy() types.Mutator { _ = "STUB: not implemented"; return *new(types.Mutator) }

func (m *Mutator) String() string { _ = "STUB: not implemented"; return "" }

// MutatorForAssignMetadata builds a Mutator from the given AssignMetadata object.
func MutatorForAssignMetadata(assignMeta *mutationsunversioned.AssignMetadata) (*Mutator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// This is not always set by the kubernetes API server

// Verifies that the given path is valid for metadata.
func isValidMetadataPath(path parser.Path) bool {
	_ = "STUB: not implemented"
	// Path must be metadata.annotations.something or metadata.labels.something
	return false
}

// IsValidAssignMetadata returns an error if the given assignmetadata object is not
// semantically valid.
func IsValidAssignMetadata(assignMeta *mutationsunversioned.AssignMetadata) error {
	_ = "STUB: not implemented"
	return nil
}
