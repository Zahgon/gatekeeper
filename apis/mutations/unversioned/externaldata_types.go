package unversioned

import (
	"errors"

	"github.com/open-policy-agent/gatekeeper/v3/pkg/mutation/types"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
)

var (
	ErrExternalDataFeatureFlag = errors.New("external data feature flag is not enabled")
	ErrExternalDataNoDefault   = errors.New("`default` must not be empty when `failurePolicy` is set to `UseDefault`")
)

// ExternalData describes the external data source to use for the mutation.
type ExternalData struct {
	// Provider is the name of the external data provider.
	// +kubebuilder:validation:Required
	Provider string `json:"provider,omitempty"`

	// DataSource specifies where to extract the data that will be sent
	// to the external data provider as parameters.
	// +kubebuilder:default="ValueAtLocation"
	DataSource types.ExternalDataSource `json:"dataSource,omitempty"`

	// FailurePolicy specifies the policy to apply when the external data
	// provider returns an error.
	// +kubebuilder:default="Fail"
	FailurePolicy types.ExternalDataFailurePolicy `json:"failurePolicy,omitempty"`

	// Default specifies the default value to use when the external data
	// provider returns an error and the failure policy is set to "UseDefault".
	Default string `json:"default,omitempty"`
}

// Validate validates the external data configuration.
func (e *ExternalData) Validate() error { _ = "STUB: not implemented"; return nil }

// GetPlaceholder returns the placeholder for the given external data provider.
func (e *ExternalData) GetPlaceholder() *ExternalDataPlaceholder {
	_ = "STUB: not implemented"
	return nil
}

// ExternalDataPlaceholder contains a placeholder value for a field
// that will eventually be resolved by an external data provider.
type ExternalDataPlaceholder struct {
	// Ref is the reference to the external data object.
	Ref *ExternalData

	// ValueAtLocation is the current value of the field that will
	// eventually be resolved by an external data provider.
	ValueAtLocation string
}

// DeepCopyWithPlaceholders returns a deep copy of the object.
func DeepCopyWithPlaceholders(u *unstructured.Unstructured) *unstructured.Unstructured {
	_ = "STUB: not implemented"
	return nil
}

// deepCopy is a copy of the runtime.DeepCopyJSONValue function that is aware
// of the ExternalDataPlaceholder type in addition to all the valid JSON types
// ref: https://github.com/kubernetes/apimachinery/blob/a58f9b57c0c7f9c017891e44431fe3a032f12f8c/pkg/runtime/converter.go#L611-L641

// Typed nil - an interface{} that contains a type map[string]interface{} with a value of nil

// Typed nil - an interface{} that contains a type []interface{} with a value of nil

// nolint:forcetypeassert
