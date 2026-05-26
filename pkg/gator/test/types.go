package test

import (
	"github.com/open-policy-agent/frameworks/constraint/pkg/instrumentation"
	"github.com/open-policy-agent/frameworks/constraint/pkg/types"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

type GatorResult struct {
	types.Result

	ViolatingObject *unstructured.Unstructured `json:"violatingObject"`

	// Trace is an explanation of the underlying constraint evaluation.
	// For instance, for OPA based evaluations, the trace is an explanation of the rego query:
	// https://www.openpolicyagent.org/docs/v0.44.0/policy-reference/#tracing
	// NOTE: This is a string pointer to differentiate between an empty ("") trace and an unset one (nil);
	// also for efficiency reasons as traces could be arbitrarily large theoretically.
	Trace *string `json:"trace"`
}

func fromFrameworkResult(frameResult *types.Result, violatingObject *unstructured.Unstructured) *GatorResult {
	_ = "STUB: not implemented"
	return nil
}

// do a deep copy to detach us from the Constraint Framework's references

// set the violating object, which is no longer part of framework results

// Response is a collection of Constraint violations for a particular Target.
// Each Result is for a distinct Constraint.
type GatorResponse struct {
	Trace   *string
	Target  string
	Results []*GatorResult
}

type GatorResponses struct {
	ByTarget     map[string]*GatorResponse
	Handled      map[string]bool
	StatsEntries []*instrumentation.StatsEntry
}

func (r *GatorResponses) Results() []*GatorResult { _ = "STUB: not implemented"; return nil }

// Make results more (but not completely) deterministic.
// After we shard Rego compilation environments, we will be able to tie
// responses to individual constraints. This is a stopgap to make tests easier
// to write until then.

// YamlGatorResult is a GatorResult minues a level of indirection on
// the ViolatingObject and with struct tags for yaml marshaling.
type YamlGatorResult struct {
	types.Result
	ViolatingObject map[string]interface{} `yaml:"violatingObject"`
	Trace           *string                `yaml:"trace,flow"`
}

// GetYamlFriendlyResults is a convenience func to remove a level of indirection between
// unstructured.Unstructured and unstructured.Unstructured.Object when calling MarshalYaml.
func GetYamlFriendlyResults(results []*GatorResult) []*YamlGatorResult {
	_ = "STUB: not implemented"
	return nil
}
