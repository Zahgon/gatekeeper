package labels

import (
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

const (
	// LabelManagedBy indicates the resource is managed by gator.
	LabelManagedBy = "gatekeeper.sh/managed-by"
	// LabelBundle indicates which bundle the resource belongs to (if any).
	LabelBundle = "gatekeeper.sh/bundle"

	// AnnotationVersion stores the policy version.
	AnnotationVersion = "gatekeeper.sh/policy-version"
	// AnnotationSource stores the policy source repository.
	AnnotationSource = "gatekeeper.sh/policy-source"
	// AnnotationInstalledAt stores the installation timestamp.
	AnnotationInstalledAt = "gatekeeper.sh/installed-at"

	// ManagedByValue is the value for the managed-by label.
	ManagedByValue = "gator"
)

// AddManagedLabels adds gator management labels and annotations to a resource.
// The source parameter identifies the catalog repository (e.g., catalog.DefaultRepository).
func AddManagedLabels(obj *unstructured.Unstructured, version, bundle, source string) {
	_ = "STUB: not implemented"
	return
}

// IsManagedByGator checks if a resource is managed by gator.
// A resource is considered managed if it has BOTH the managed-by label
// AND the policy-source annotation.
func IsManagedByGator(obj *unstructured.Unstructured) bool { _ = "STUB: not implemented"; return false }

// Require both label and annotation for a resource to be considered managed

// GetPolicyVersion returns the policy version from annotations.
func GetPolicyVersion(obj *unstructured.Unstructured) string { _ = "STUB: not implemented"; return "" }

// GetBundle returns the bundle name from labels.
func GetBundle(obj *unstructured.Unstructured) string { _ = "STUB: not implemented"; return "" }

// GetInstalledAt returns the installation timestamp from annotations.
func GetInstalledAt(obj *unstructured.Unstructured) string { _ = "STUB: not implemented"; return "" }
