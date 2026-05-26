package expand

import (
	"github.com/open-policy-agent/gatekeeper/v3/apis/expansion/unversioned"
	mutationsunversioned "github.com/open-policy-agent/gatekeeper/v3/apis/mutations/unversioned"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/expansion"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/mutation"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/mutation/types"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

var mutatorKinds = map[string]bool{
	"Assign":         true,
	"AssignMetadata": true,
	"ModifySet":      true,
	"AssignImage":    true,
}

type Expander struct {
	mutators           []types.Mutator
	templateExpansions []*unversioned.ExpansionTemplate
	objects            []*unstructured.Unstructured
	namespaces         map[string]*corev1.Namespace
	expSystem          *expansion.System
	mutSystem          *mutation.System
}

func Expand(resources []*unstructured.Unstructured) ([]*unstructured.Unstructured, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func NewExpander(resources []*unstructured.Unstructured) (*Expander, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (er *Expander) Expand(resource *unstructured.Unstructured) ([]*expansion.Resultant, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Mutate the base resource before expanding it

func (er *Expander) NamespaceForResource(r *unstructured.Unstructured) (*corev1.Namespace, bool) {
	_ = "STUB: not implemented"
	return nil, false
}

func (er *Expander) addResources(resources []*unstructured.Unstructured) error {
	_ = "STUB: not implemented"
	return nil
}

func (er *Expander) addMutator(mut *unstructured.Unstructured) error {
	_ = "STUB: not implemented"
	return nil
}

func (er *Expander) add(u *unstructured.Unstructured) error { _ = "STUB: not implemented"; return nil }

// Any resource can technically be a generator

func (er *Expander) addExpansionTemplate(u *unstructured.Unstructured) error {
	_ = "STUB: not implemented"
	return nil
}

func (er *Expander) addNamespace(u *unstructured.Unstructured) error {
	_ = "STUB: not implemented"
	return nil
}

func isExpansion(u *unstructured.Unstructured) bool { _ = "STUB: not implemented"; return false }

func isMutator(obj *unstructured.Unstructured) bool { _ = "STUB: not implemented"; return false }

func isNamespace(obj *unstructured.Unstructured) bool { _ = "STUB: not implemented"; return false }

func convertUnstructuredToTyped(u *unstructured.Unstructured, obj interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func convertExpansionTemplate(u *unstructured.Unstructured) (*unversioned.ExpansionTemplate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func convertAssign(u *unstructured.Unstructured) (*mutationsunversioned.Assign, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func convertAssignMetadata(u *unstructured.Unstructured) (*mutationsunversioned.AssignMetadata, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func convertModifySet(u *unstructured.Unstructured) (*mutationsunversioned.ModifySet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func convertAssignImage(u *unstructured.Unstructured) (*mutationsunversioned.AssignImage, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func convertNamespace(u *unstructured.Unstructured) (*corev1.Namespace, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
