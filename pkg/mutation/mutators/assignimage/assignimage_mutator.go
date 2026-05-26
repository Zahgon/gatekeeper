package assignimage

import (
	mutationsunversioned "github.com/open-policy-agent/gatekeeper/v3/apis/mutations/unversioned"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/logging"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/mutation/mutators/core"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/mutation/path/parser"
	patht "github.com/open-policy-agent/gatekeeper/v3/pkg/mutation/path/tester"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/mutation/schema"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/mutation/types"
	runtimeschema "k8s.io/apimachinery/pkg/runtime/schema"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

var log = logf.Log.WithName("mutation").WithValues(logging.Process, "mutation", logging.Mutator, "assignimage")

// Mutator is a mutator object built out of an AssignImage instance.
type Mutator struct {
	id          types.ID
	assignImage *mutationsunversioned.AssignImage

	path parser.Path

	// bindings are the set of GVKs this Mutator applies to.
	bindings []runtimeschema.GroupVersionKind
	tester   *patht.Tester
}

// Mutator implements mutatorWithSchema.
var _ schema.MutatorWithSchema = &Mutator{}

func (m *Mutator) Matches(mutable *types.Mutable) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (m *Mutator) TerminalType() parser.NodeType {
	_ = "STUB: not implemented"
	return *new(parser.NodeType)
}

func (m *Mutator) Mutate(mutable *types.Mutable) (bool, error) {
	_ = "STUB: not implemented"
	return false, nil
}

func (m *Mutator) MustTerminate() bool { _ = "STUB: not implemented"; return false }

func (m *Mutator) ID() types.ID { _ = "STUB: not implemented"; return *new(types.ID) }

func (m *Mutator) SchemaBindings() []runtimeschema.GroupVersionKind {
	_ = "STUB: not implemented"
	return nil
}

func (m *Mutator) HasDiff(mutator types.Mutator) bool { _ = "STUB: not implemented"; return false }

// different types, different

// any difference in spec may be enough

func (m *Mutator) Path() parser.Path { _ = "STUB: not implemented"; return *new(parser.Path) }

func (m *Mutator) DeepCopy() types.Mutator { _ = "STUB: not implemented"; return *new(types.Mutator) }

func (m *Mutator) String() string { _ = "STUB: not implemented"; return "" }

// MutatorForAssignImage returns a mutator built from
// the given assignImage instance.
func MutatorForAssignImage(assignImage *mutationsunversioned.AssignImage) (*Mutator, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// This is not always set by the kubernetes API server

func hasListTerminal(path parser.Path) bool { _ = "STUB: not implemented"; return false }

var _ core.Setter = setter{}

type setter struct {
	tag    string
	domain string
	path   string
}

func (s setter) KeyedListOkay() bool { _ = "STUB: not implemented"; return false }

func (s setter) KeyedListValue() (map[string]interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s setter) SetValue(obj map[string]interface{}, key string) error {
	_ = "STUB: not implemented"
	return nil
}

// IsValidAssignImage returns an error if the given assignImage object is not
// semantically valid.
func IsValidAssignImage(assignImage *mutationsunversioned.AssignImage) error {
	_ = "STUB: not implemented"
	return nil
}
