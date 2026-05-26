package fakes

import (
	gvkmanifestv1alpha1 "github.com/open-policy-agent/gatekeeper/v3/apis/gvkmanifest/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// GVKManifestFor returns a GVKManifest resource with the given name for the requested set of resources.
func GVKManifestFor(name string, gvks []schema.GroupVersionKind) *gvkmanifestv1alpha1.GVKManifest {
	_ = "STUB: not implemented"
	return nil
}
