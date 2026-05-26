package testhelpers

import (
	"github.com/open-policy-agent/gatekeeper/v3/pkg/mutation/match"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/mutation/path/parser"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/mutation/types"
)

var _ types.Mutator = &DummyMutator{}

// DummyMutator is a blank mutator that makes it easier to test the core mutation function.
type DummyMutator struct {
	name  string
	value interface{}
	path  parser.Path
	match match.Match
}

func (d *DummyMutator) DeepCopy() types.Mutator {
	_ = "STUB: not implemented"
	return *new(types.Mutator)
}

func (d *DummyMutator) HasDiff(m types.Mutator) bool { _ = "STUB: not implemented"; return false }

func (d *DummyMutator) ID() types.ID { _ = "STUB: not implemented"; return *new(types.ID) }

func (d *DummyMutator) Path() parser.Path { _ = "STUB: not implemented"; return *new(parser.Path) }

func (d *DummyMutator) Matches(mutable *types.Mutable) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (d *DummyMutator) Mutate(mutable *types.Mutable) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (d *DummyMutator) MustTerminate() bool { _ = "STUB: not implemented"; return false }

func (d *DummyMutator) String() string { _ = "STUB: not implemented"; return "" }

func NewDummyMutator(name, path string, value interface{}) *DummyMutator {
	_ = "STUB: not implemented"
	return nil
}

// BigName returns a 64-length string.
func BigName() string { _ = "STUB: not implemented"; return "" }

// 8 X 8 = 64
