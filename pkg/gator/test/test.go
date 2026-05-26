package test

import (
	"github.com/open-policy-agent/frameworks/constraint/pkg/apis"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/gator"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
)

var scheme *runtime.Scheme

func init() {
	scheme = runtime.NewScheme()
	err := apis.AddToScheme(scheme)
	if err != nil {
		panic(err)
	}
}

func Test(objs []*unstructured.Unstructured, opts ...gator.Opt) (*GatorResponses, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// search for templates, add them if they exist

// add all constraints.  A constraint must be added after its associated
// template or OPA will return an error

// finally, add all the data.

// create the expander

// now audit all objects

// Try to attach the namespace if it was supplied (ns will be nil otherwise)

// Attempt to expand the obj and review resultant resources (if any)

// convert framework results to gator results, which contain a
// reference to the violating resource

func WithGatherStats() gator.Opt { _ = "STUB: not implemented"; return *new(gator.Opt) }

func WithTrace() gator.Opt { _ = "STUB: not implemented"; return *new(gator.Opt) }

func WithK8sCEL(gatherStats bool) gator.Opt { _ = "STUB: not implemented"; return *new(gator.Opt) }
