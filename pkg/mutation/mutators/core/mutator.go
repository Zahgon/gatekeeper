package core

import (
	"github.com/open-policy-agent/gatekeeper/v3/apis/mutations/unversioned"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/mutation/match"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/mutation/path/parser"
	patht "github.com/open-policy-agent/gatekeeper/v3/pkg/mutation/path/tester"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/mutation/types"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// NewTester returns a path.Tester for the given object name, kind path and
// pathtests.
func NewTester(name string, kind string, path parser.Path, ptests []unversioned.PathTest) (*patht.Tester, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// NewValidatedBindings returns a slice of gvks from the given applies, or an
// error if the applies are invalid.
func NewValidatedBindings(name string, kind string, applies []match.ApplyTo) ([]schema.GroupVersionKind, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func gatherPathTests(mutName string, mutKind string, pts []unversioned.PathTest) ([]patht.Test, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func getSortedGVKs(bindings []match.ApplyTo) []schema.GroupVersionKind {
	_ = "STUB: not implemented"
	// deduplicate GVKs
	return nil
}

// we iterate over the map in a stable order so that
// unit tests won't be flaky.

// HasMetadataRoot returns true if the root node at given path references the
// metadata field.
func HasMetadataRoot(path parser.Path) bool { _ = "STUB: not implemented"; return false }

// CheckKeyNotChanged does not allow to change the key field of
// a list element. A path like foo[name: bar].name is rejected.
func CheckKeyNotChanged(p parser.Path) error { _ = "STUB: not implemented"; return nil }

func MatchWithApplyTo(mut *types.Mutable, applies []match.ApplyTo, mat *match.Match) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func ValidateName(name string) error { _ = "STUB: not implemented"; return nil }
