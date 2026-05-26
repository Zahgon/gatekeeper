package reader

import (
	"io"
	"io/fs"

	"github.com/open-policy-agent/frameworks/constraint/pkg/core/templates"
	configv1alpha1 "github.com/open-policy-agent/gatekeeper/v3/apis/config/v1alpha1"
	gvkmanifestv1alpha1 "github.com/open-policy-agent/gatekeeper/v3/apis/gvkmanifest/v1alpha1"
	syncsetv1alpha1 "github.com/open-policy-agent/gatekeeper/v3/apis/syncset/v1alpha1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

type versionless interface {
	ToVersionless() (*templates.ConstraintTemplate, error)
}

// jsonLookaheadBytes is the number of bytes the JSON and YAML decoder will
// look into the data it's reading to determine if the document is JSON or
// YAML.  1024 was a guess that's worked so far.
const jsonLookaheadBytes int = 1024

// clean removes the following from yaml:
// 1) Empty lines
// 2) Lines with only space characters
// 3) Lines which are only comments
//
// This prevents us from attempting to parse an empty yaml document and failing.
func clean(yaml string) string { _ = "STUB: not implemented"; return "" }

func ReadUnstructureds(bytes []byte) ([]*unstructured.Unstructured, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ReadUnstructured(bytes []byte) (*unstructured.Unstructured, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReadTemplate reads the contents of the path and returns the
// ConstraintTemplate it defines. Returns an error if the file does not define
// a ConstraintTemplate.
func ReadTemplate(scheme *runtime.Scheme, f fs.FS, path string) (*templates.ConstraintTemplate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ToStructured converts an unstructured object into an object with the schema defined
// by u's group, version, and kind.
func ToStructured(scheme *runtime.Scheme, u *unstructured.Unstructured) (runtime.Object, error) {
	_ = "STUB: not implemented"
	return *new(runtime.Object), nil
}

// The type isn't registered in the scheme.

// YAML parsing doesn't properly handle ObjectMeta, so we must
// marshal/unmashal through JSON.

// Indicates a bug in unstructured.MarshalJSON(). Any Unstructured
// unmarshalled from YAML should be marshallable to JSON.

// ToTemplate converts an unstructured template into a versionless ConstraintTemplate struct.
func ToTemplate(scheme *runtime.Scheme, u *unstructured.Unstructured) (*templates.ConstraintTemplate, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// This shouldn't happen unless there's a bug in the conversion functions.
// Most likely it means the conversion functions weren't generated.

// ToSyncSet converts an unstructured SyncSet into a SyncSet struct.
func ToSyncSet(scheme *runtime.Scheme, u *unstructured.Unstructured) (*syncsetv1alpha1.SyncSet, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ToConfig converts an unstructured Config into a Config struct.
func ToConfig(scheme *runtime.Scheme, u *unstructured.Unstructured) (*configv1alpha1.Config, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ToGVKManifest converts an unstructured GVKManifest into a GVKManifest struct.
func ToGVKManifest(scheme *runtime.Scheme, u *unstructured.Unstructured) (*gvkmanifestv1alpha1.GVKManifest, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReadObject reads a file from the filesystem abstraction at the specified
// path, and returns an unstructured.Unstructured object if the file can be
// successfully unmarshalled.
func ReadObject(f fs.FS, path string) (*unstructured.Unstructured, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ReadConstraint(f fs.FS, path string) (*unstructured.Unstructured, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func ReadExpansions(f fs.FS, path string) ([]*unstructured.Unstructured, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ReadK8sResources reads JSON or YAML k8s resources from an io.Reader,
// decoding them into Unstructured objects and returning those objects as a
// slice.
func ReadK8sResources(r io.Reader) ([]*unstructured.Unstructured, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// skip empty resources

func IsTemplate(u *unstructured.Unstructured) bool { _ = "STUB: not implemented"; return false }

func IsConfig(u *unstructured.Unstructured) bool { _ = "STUB: not implemented"; return false }

func IsSyncSet(u *unstructured.Unstructured) bool { _ = "STUB: not implemented"; return false }

func IsGVKManifest(u *unstructured.Unstructured) bool { _ = "STUB: not implemented"; return false }

func IsConstraint(u *unstructured.Unstructured) bool { _ = "STUB: not implemented"; return false }
