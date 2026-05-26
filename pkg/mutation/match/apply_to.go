package match

import "k8s.io/apimachinery/pkg/runtime/schema"

// AppliesTo checks if any item the given slice of ApplyTo applies to the given object.
func AppliesTo(applyTo []ApplyTo, gvk schema.GroupVersionKind) bool {
	_ = "STUB: not implemented"
	return false
}

// ApplyTo determines what GVKs items the mutation should apply to.
// Globs are not allowed.
// +kubebuilder:object:generate=true
type ApplyTo struct {
	Groups   []string `json:"groups,omitempty"`
	Kinds    []string `json:"kinds,omitempty"`
	Versions []string `json:"versions,omitempty"`
}

// Flatten returns the set of GroupVersionKinds this ApplyTo matches.
// The GVKs are not guaranteed to be sorted or unique.
func (a ApplyTo) Flatten() []schema.GroupVersionKind { _ = "STUB: not implemented"; return nil }

// Matches returns true if the Object's Group, Version, and Kind are contained
// in the ApplyTo's match lists.
func (a ApplyTo) Matches(gvk schema.GroupVersionKind) bool { _ = "STUB: not implemented"; return false }
