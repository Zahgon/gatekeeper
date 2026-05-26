package target

import (
	"github.com/open-policy-agent/frameworks/constraint/pkg/core/constraints"
	"github.com/open-policy-agent/frameworks/constraint/pkg/handler"
	"github.com/open-policy-agent/gatekeeper/v3/pkg/mutation/match"
	"github.com/pkg/errors"
	"k8s.io/apiextensions-apiserver/pkg/apis/apiextensions"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime/schema"
)

// nolint: revive // Moved error out of pkg/webhook/admission; needs capitalization for backwards compat.
var ErrOldObjectIsNil = errors.New("oldObject cannot be nil for DELETE operations")

// Name is the name of Gatekeeper's Kubernetes validation target.
const Name = "admission.k8s.gatekeeper.sh"

type K8sValidationTarget struct {
	cache nsCache
}

var (
	_ handler.TargetHandler = &K8sValidationTarget{}
	_ handler.Cacher        = &K8sValidationTarget{}
)

func (h *K8sValidationTarget) GetName() string { _ = "STUB: not implemented"; return "" }

func (h *K8sValidationTarget) processUnstructured(o *unstructured.Unstructured) (bool, []string, interface{}, error) {
	_ = "STUB: not implemented"
	// Namespace will be "" for cluster objects
	return false, nil, nil, nil
}

func clusterScopedKey(gvk schema.GroupVersionKind, name string) []string {
	_ = "STUB: not implemented"
	return nil
}

func namespaceScopedKey(namespace string, gvk schema.GroupVersionKind, name string) []string {
	_ = "STUB: not implemented"
	return nil
}

func (h *K8sValidationTarget) ProcessData(obj interface{}) (bool, []string, interface{}, error) {
	_ = "STUB: not implemented"
	return false, nil, nil, nil
}

func (h *K8sValidationTarget) HandleReview(obj interface{}) (bool, interface{}, error) {
	_ = "STUB: not implemented"
	return false,

		// handleReview returns a complete *gkReview to pass to the Client.
		nil, nil
}

func (h *K8sValidationTarget) handleReview(obj interface{}) (bool, *gkReview, error) {
	_ = "STUB: not implemented"
	return false, nil, nil
}

func augmentedUnstructuredToAdmissionRequest(obj AugmentedUnstructured) (*gkReview, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func unstructuredToAdmissionRequest(obj *unstructured.Unstructured) (*gkReview, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (h *K8sValidationTarget) MatchSchema() apiextensions.JSONSchemaProps {
	_ = "STUB: not implemented"
	return *new(apiextensions.JSONSchemaProps)
}

func (h *K8sValidationTarget) ValidateConstraint(u *unstructured.Unstructured) error {
	_ = "STUB: not implemented"
	return nil
}

func convertToLabelSelector(object map[string]interface{}) (*metav1.LabelSelector, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func convertToMatch(object map[string]interface{}) (*match.Match, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// ToMatcher converts .spec.match in mutators to Matcher.
func (h *K8sValidationTarget) ToMatcher(u *unstructured.Unstructured) (constraints.Matcher, error) {
	_ = "STUB: not implemented"
	return *new(constraints.Matcher), nil
}

func (h *K8sValidationTarget) GetCache() handler.Cache {
	_ = "STUB: not implemented"

	// setObjectOnDelete enforces that we use at least K8s API v1.15.0+ on DELETE operations
	// and copies over the oldObject into the Object field for the given AdmissionRequest.
	return *new(handler.Cache)
}

func setObjectOnDelete(review *gkReview) error {
	_ = "STUB: not implemented"
	// Directly accessing the Operation field from AdmissionRequest, as it is embedded within gkReview.
	return nil
}

// oldObject is the existing object.
// It is null for DELETE operations in API servers prior to v1.15.0.
// https://github.com/kubernetes/website/pull/14671

// For admission webhooks registered for DELETE operations on k8s built APIs or CRDs,
// the apiserver now sends the existing object as admissionRequest.Request.OldObject to the webhook
// object is the new object being admitted.
// It is null for DELETE operations.
// https://github.com/kubernetes/kubernetes/pull/76346
