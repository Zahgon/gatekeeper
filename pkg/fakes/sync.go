package fakes

import (
	configv1alpha1 "github.com/open-policy-agent/gatekeeper/v3/apis/config/v1alpha1"
	syncsetv1alpha1 "github.com/open-policy-agent/gatekeeper/v3/apis/syncset/v1alpha1"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// SyncSetFor returns a syncset resource with the given name for the requested set of resources.
func SyncSetFor(name string, kinds []schema.GroupVersionKind) *syncsetv1alpha1.SyncSet {
	_ = "STUB: not implemented"
	return nil
}

// ConfigFor returns a config resource with a SyncOnly containing the requested set of resources.
func ConfigFor(kinds []schema.GroupVersionKind) *configv1alpha1.Config {
	_ = "STUB: not implemented"
	return nil
}
