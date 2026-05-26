package k8scel

import (
	"context"
	"sync"

	"github.com/open-policy-agent/frameworks/constraint/pkg/client/drivers"
	"github.com/open-policy-agent/frameworks/constraint/pkg/client/reviews"
	"github.com/open-policy-agent/frameworks/constraint/pkg/core/templates"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/logging"
	"github.com/open-policy-agent/opa/storage"
	admissionv1 "k8s.io/api/admission/v1"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apiserver/pkg/admission/plugin/policy/validating"
	logf "sigs.k8s.io/controller-runtime/pkg/log"
)

// NOTE: This is a PROTOTYPE driver. Do not use this for any critical work
// and be aware that its behavior may change at any time.

// Friction log:
//   there is no way to re-use the matcher interface here, as it requires an informer... not sure we need to use
//   the matchers, as match Criteria should take care of things.

//   "Expression" is a bit confusing, since it doesn't tell me whether "true" implies violation or not: "requirement", "mustSatisfy"?
//
//
//   From the Validation help text:
//      Equality on arrays with list type of 'set' or 'map' ignores element order, i.e. [1, 2] == [2, 1].
//      Concatenation on arrays with x-kubernetes-list-type use the semantics of the list type:
//   Is this type metadata available shift-left? Likely not. Can the expectation be built into the operators?
//
//   Other friction points are commented with the keyword FRICTION.

const (
	runTimeNS            = "runTimeNS"
	runTimeNSDescription = "the number of nanoseconds it took to evaluate the constraint"
)

var (
	_   drivers.Driver = &Driver{}
	log                = logf.Log.WithName("k8scel-driver").WithValues(logging.Process, "k8scel")
)

type Driver struct {
	mux         sync.RWMutex
	validators  map[string]*validatorWrapper
	gatherStats bool
}

type validatorWrapper struct {
	validator validating.Validator
}

func (d *Driver) Name() string { _ = "STUB: not implemented"; return "" }

func (d *Driver) AddTemplate(_ context.Context, ct *templates.ConstraintTemplate) error {
	_ = "STUB: not implemented"
	return nil
}

// FRICTION: Note that compilation errors are possible, but we cannot introspect to see whether any
// occurred

// We don't want to have access to parameters for anything other than driver-defined logic, so we
// can keep the user from accessing the full constraint schema.

// StrictCost is now enabled by default in MustBaseEnvSet (via StrictCostOpt in baseOpts).

func (d *Driver) RemoveTemplate(_ context.Context, ct *templates.ConstraintTemplate) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *Driver) AddConstraint(_ context.Context, _ *unstructured.Unstructured) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *Driver) RemoveConstraint(_ context.Context, _ *unstructured.Unstructured) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *Driver) AddData(_ context.Context, _ string, _ storage.Path, _ interface{}) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *Driver) RemoveData(_ context.Context, _ string, _ storage.Path) error {
	_ = "STUB: not implemented"
	return nil
}

func (d *Driver) Query(ctx context.Context, target string, constraints []*unstructured.Unstructured, review interface{}, opts ...reviews.ReviewOpt) (*drivers.QueryResponse, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Gatekeeper sets Object to oldObject on DELETE requests
// however, Kubernetes does not do this for ValidatingAdmissionPolicy.
// in order for evaluation in both environments to behave identically,
// we must be sure that Object is unset on DELETE. Users who need
// the "if DELETE Object == OldObject" behavior should use the
// `variables.anyObject` variable instead.

// template name is the lowercase of its kind

// this should never happen, but best not to panic if the pointer is ever nil.

// Convert namespace from map[string]interface{} to *corev1.Namespace if provided.
// This enables CEL expressions to access namespaceObject for namespace-based policies.

func (d *Driver) Dump(_ context.Context) (string, error) { _ = "STUB: not implemented"; return "", nil }

func (d *Driver) GetDescriptionForStat(statName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

type ARGetter interface {
	GetAdmissionRequest() *admissionv1.AdmissionRequest
}

type IsAdmissionGetter interface {
	IsAdmissionRequest() bool
}

// mapToNamespace converts a map[string]interface{} representation of a namespace
// to a *corev1.Namespace. This is used to pass the namespace object to the CEL
// validator for namespaceObject support.
func mapToNamespace(ns map[string]interface{}) (*corev1.Namespace, error) {
	_ = "STUB: not implemented"
	return nil, nil
}
