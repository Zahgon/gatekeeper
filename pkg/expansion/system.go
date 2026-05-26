package expansion

import (
	"flag"
	"sync"

	expansionunversioned "github.com/open-policy-agent/gatekeeper/v3/apis/expansion/unversioned"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/logging"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/mutation"
	mutationtypes "github.com/open-policy-agent/gatekeeper/v3/pkg/mutation/types"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

var (
	ExpansionEnabled *bool
	log              = logf.Log.WithName("expansion").WithValues(logging.Process, "expansion")
)

// maxRecursionDepth specifies the maximum call depth for recursive expansion.
// Theoretically, it should be impossible for a cycle to be created but this
// measure is put in place as a safeguard.
const maxRecursionDepth = 30

func init() {
	ExpansionEnabled = flag.Bool("enable-generator-resource-expansion", true, "(beta) Enable the expansion of generator resources")
}

type System struct {
	lock           sync.RWMutex
	mutationSystem *mutation.System
	db             templateDB
}

type Resultant struct {
	Obj               *unstructured.Unstructured
	TemplateName      string
	EnforcementAction string
}

type TemplateID string

type IDSet map[TemplateID]bool

func keyForTemplate(template *expansionunversioned.ExpansionTemplate) TemplateID {
	_ = "STUB: not implemented"
	return *new(TemplateID)
}

func (s *System) UpsertTemplate(template *expansionunversioned.ExpansionTemplate) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *System) RemoveTemplate(template *expansionunversioned.ExpansionTemplate) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *System) GetConflicts() IDSet { _ = "STUB: not implemented"; return *new(IDSet) }

func ValidateTemplate(template *expansionunversioned.ExpansionTemplate) error {
	_ = "STUB: not implemented"
	return nil
}

// Make sure template does not form a self-edge (i.e. a template configured
// to expand its own output)

func sourcePath(source string) []string { _ = "STUB: not implemented"; return nil }

func prettyResource(v interface{}) string { _ = "STUB: not implemented"; return "" }

func genGVKToSchemaGVK(gvk expansionunversioned.GeneratedGVK) schema.GroupVersionKind {
	_ = "STUB: not implemented"
	return *new(schema.GroupVersionKind)
}

// Expand expands `base` into resultant resources, and applies any applicable
// mutators. If no ExpansionTemplates match `base`, an empty slice
// will be returned. If `s.mutationSystem` is nil, no mutations will be applied.
func (s *System) Expand(base *mutationtypes.Mutable) ([]*Resultant, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (s *System) expandRecursive(base *mutationtypes.Mutable, resultants *[]*Resultant, depth int) error {
	_ = "STUB: not implemented"
	return nil
}

func (s *System) expand(base *mutationtypes.Mutable) ([]*Resultant, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func expandResource(obj *unstructured.Unstructured, ns *corev1.Namespace, template *expansionunversioned.ExpansionTemplate) (*unstructured.Unstructured, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// if not found, then the resulting resource may be cluster scoped.

// ensureOwnerReference appends an OwnerReference describing parent to the resultant
// resource if one is not already present.
func ensureOwnerReference(resultant, parent *unstructured.Unstructured) {
	_ = "STUB: not implemented"
	return
}

// mockNameForResource returns a mock name for a resultant resource created
// from expanding `gen`. The name will be of the form:
// "<generator name>-<resultant kind>". For example, a deployment named
// `nginx-deployment` will produce a resultant named `nginx-deployment-pod`.
func mockNameForResource(gen *unstructured.Unstructured, gvk schema.GroupVersionKind) string {
	_ = "STUB: not implemented"
	return ""
}

func NewSystem(mutationSystem *mutation.System) *System { _ = "STUB: not implemented"; return nil }
