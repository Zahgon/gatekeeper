package test

import (
	cfapis "github.com/open-policy-agent/frameworks/constraint/pkg/apis"
	gkapis "github.com/open-policy-agent/gatekeeper/v3/apis"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/cachemanager/parser"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

var scheme *runtime.Scheme

func init() {
	scheme = runtime.NewScheme()
	err := cfapis.AddToScheme(scheme)
	if err != nil {
		panic(err)
	}
	err = gkapis.AddToScheme(scheme)
	if err != nil {
		panic(err)
	}
}

// Reads a list of unstructured objects and a string containing supported GVKs and
// outputs a set of missing sync requirements per template and ingestion problems per template.
func Test(unstrucs []*unstructured.Unstructured, omitGVKManifest bool) (map[string]parser.SyncRequirements, map[string]error, error) {
	_ = "STUB: not implemented"
	return nil, nil, nil
}

// Don't assess requirement fulfillment if there was an error parsing any of the templates.

// Crosscheck synced gvks with supported gvks.

// Fetch syncrequirements from template
