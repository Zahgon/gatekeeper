package schema

import (
	"github.com/open-policy-agent/frameworks/constraint/pkg/core/templates"
	admissionv1 "k8s.io/api/admissionregistration/v1"
	admissionv1beta1 "k8s.io/api/admissionregistration/v1beta1"
	"k8s.io/apiserver/pkg/admission/plugin/cel"
)

const (
	// Name is the name of the driver.
	Name = "K8sNativeValidation"
	// ReservedPrefix signifies a prefix that no user-defined value (variable, matcher, etc.) is allowed to have.
	// This gives us the ability to add new variables in the future without worrying about breaking pre-existing templates.
	ReservedPrefix = "gatekeeper_internal_"
	// ParamsName is the VAP variable constraint parameters will be bound to.
	ParamsName = "params"
	// ObjectName is the VAP variable that describes either an object or (on DELETE requests) oldObject.
	ObjectName = "anyObject"
)

type Validation struct {
	// A CEL expression. Maps to ValidationAdmissionPolicy's `spec.validations`.
	Expression        string `json:"expression,omitempty"`
	Message           string `json:"message,omitempty"`
	MessageExpression string `json:"messageExpression,omitempty"`
}

type MatchCondition struct {
	Name       string `json:"name"`
	Expression string `json:"expression"`
}

type Variable struct {
	// A CEL variable definition. Maps to ValidationAdmissionPolicy's `spec.variables`.
	Name       string `json:"name,omitempty"`
	Expression string `json:"expression,omitempty"`
}

type Source struct {
	// Validations maps to ValidatingAdmissionPolicy's `spec.validations`.
	Validations []Validation `json:"validations,omitempty"`

	// FailurePolicy maps to ValidatingAdmissionPolicy's `spec.failurePolicy`.
	FailurePolicy *string `json:"failurePolicy,omitempty"`

	// MatchConditions maps to ValidatingAdmissionPolicy's `spec.matchConditions`.
	MatchConditions []MatchCondition `json:"matchCondition,omitempty"`

	// Variables maps to ValidatingAdmissionPolicy's `spec.variables`.
	Variables []Variable `json:"variables,omitempty"`

	// GenerateVAP enables/disables VAP generation and enforcement for policy.
	GenerateVAP *bool `json:"generateVAP,omitempty"`
}

func (in *Source) Validate() error { _ = "STUB: not implemented"; return nil }

func (in *Source) validateMatchConditions() error { _ = "STUB: not implemented"; return nil }

func (in *Source) GetMatchConditions() ([]cel.ExpressionAccessor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (in *Source) GetV1Beta1MatchConditions() ([]admissionv1beta1.MatchCondition, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (in *Source) validateVariables() error { _ = "STUB: not implemented"; return nil }

func (in *Source) GetVariables() ([]cel.NamedExpressionAccessor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (in *Source) GetV1Beta1Variables() ([]admissionv1beta1.Variable, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (in *Source) GetValidations() ([]cel.ExpressionAccessor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (in *Source) GetV1Beta1Validatons() ([]admissionv1beta1.Validation, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (in *Source) GetMessageExpressions() ([]cel.ExpressionAccessor, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (in *Source) GetFailurePolicy() (*admissionv1.FailurePolicyType, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (in *Source) GetV1Beta1FailurePolicy() (*admissionv1beta1.FailurePolicyType, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// MustToUnstructured() is a convenience method for converting to unstructured.
// Intended for testing. It will panic on error.
func (in *Source) MustToUnstructured() map[string]interface{} {
	_ = "STUB: not implemented"
	return nil
}

func GetSource(code templates.Code) (*Source, error) { _ = "STUB: not implemented"; return nil, nil }

func GetSourceFromTemplate(ct *templates.ConstraintTemplate) (*Source, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// HasCELEngine checks if the ConstraintTemplate has a CEL engine code block,
// without validating the source content.
func HasCELEngine(ct *templates.ConstraintTemplate) bool { _ = "STUB: not implemented"; return false }
