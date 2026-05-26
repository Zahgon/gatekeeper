package core

import (
	"github.com/open-policy-agent/gatekeeper/v3/pkg/mutation/path/parser"
	path "github.com/open-policy-agent/gatekeeper/v3/pkg/mutation/path/tester"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/mutation/types"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var _ types.MetadataGetter = &metadata{}

type metadata client.ObjectKey

func (m *metadata) GetName() string { _ = "STUB: not implemented"; return "" }

func (m *metadata) GetNamespace() string { _ = "STUB: not implemented"; return "" }

func Mutate(
	path parser.Path,
	tester *path.Tester,
	setter Setter,
	obj *unstructured.Unstructured,
) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

type mutatorState struct {
	path     parser.Path
	tester   *path.Tester
	setter   Setter
	metadata *metadata
}

// mutateInternal mutates the resource recursively. It returns false if there has been no change
// to any downstream objects in the tree, indicating that the mutation should not be persisted.
func (s *mutatorState) mutateInternal(current interface{}, depth int) (bool, interface{}, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

// Path entry type does not match current object

// we have hit the end of our path, this is the base case

// Next element is missing and needs to be added

// Path entry type does not match current object

// base case

// if someone says "MustNotExist" for a glob, that condition can never be satisfied

// If no matching element in the array was found in non Globbed list, create a new element

func (s *mutatorState) setListElementToValue(currentAsList []interface{}, listPathEntry *parser.List, depth int) (bool, []interface{}, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

func (s *mutatorState) createMissingElement(depth int) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Create new element of type

// Set new keyfield

// Path entry type does not match current object

func nestedFieldNoCopy(current interface{}, key string) (interface{}, bool, error) {
	_ = "STUB: not implemented"
	return nil, false, nil
}
