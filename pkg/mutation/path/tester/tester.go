package tester

import (
	"errors"

	"github.com/open-policy-agent/gatekeeper/v3/pkg/mutation/path/parser"
)

// Condition describes whether the path either MustExist or MustNotExist in the original object
// +kubebuilder:validation:Enum=MustExist;MustNotExist
type Condition string

const (
	// MustExist means that an object must exist at the given path entry.
	MustExist = Condition("MustExist")
	// MustNotExist means that an object must not exist at the given path entry.
	MustNotExist = Condition("MustNotExist")
)

var conditions = map[string]Condition{
	"MustExist":    MustExist,
	"MustNotExist": MustNotExist,
}

// Base errors for validating path tests.
var (
	ErrPrefix   = errors.New("all subpaths must be a prefix of the `location` value of the mutation")
	ErrConflict = errors.New("conflicting path test conditions")
)

// StringToCondition translates a user-provided string into a Test Condition.
func StringToCondition(s string) (Condition, error) {
	_ = "STUB: not implemented"
	return *new(Condition), nil
}

// Test describes a condition that the object must satisfy.
type Test struct {
	SubPath   parser.Path
	Condition Condition
}

func isPrefix(short, long parser.Path) bool { _ = "STUB: not implemented"; return false }

// validatePathTests returns whether a set of path tests are valid against the provided location.
func validatePathTests(location parser.Path, pathTests []Test) error {
	_ = "STUB: not implemented"
	return nil
}

// New creates a new Tester object.
func New(location parser.Path, tests []Test) (*Tester, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Read in all tests before checking for conflicts.

// Check for conflicts.

// Check that a single node must not be marked both MustExist and MustNotExist.

// Check that child nodes are not required if they have a forbidden parent.

// Tester knows whether it's okay that an object exists at a given path depth.
type Tester struct {
	tests map[int]Condition
}

// ExistsOkay returns true if it's okay that an object exists.
func (pt *Tester) ExistsOkay(depth int) bool { _ = "STUB: not implemented"; return false }

// MissingOkay returns true if it's okay that an object is missing.
func (pt *Tester) MissingOkay(depth int) bool { _ = "STUB: not implemented"; return false }

// DeepCopy returns a deep copy of the tester.
func (pt *Tester) DeepCopy() *Tester { _ = "STUB: not implemented"; return nil }
