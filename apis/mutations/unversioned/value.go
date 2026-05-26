package unversioned

import (
	"errors"

	"github.com/open-policy-agent/gatekeeper/v3/pkg/mutation/types"
)

var (
	ErrInvalidAssignField  = errors.New("invalid assign field")
	ErrInvalidFromMetadata = errors.New("invalid fromMetadata field")
)

type Field string

const (
	// ObjNamespace => metadata.namespace.
	ObjNamespace = Field("namespace")

	// ObjName => metadata.name.
	ObjName = Field("name")
)

var validFields = map[Field]bool{
	ObjNamespace: true,
	ObjName:      true,
}

type AssignField struct {
	// Value is a constant value that will be assigned to `location`
	// +kubebuilder:validation:Schemaless
	// +kubebuilder:pruning:PreserveUnknownFields
	Value *types.Anything `json:"value,omitempty"`

	// FromMetadata assigns a value from the specified metadata field.
	FromMetadata *FromMetadata `json:"fromMetadata,omitempty"`

	// ExternalData describes the external data provider to be used for mutation.
	ExternalData *ExternalData `json:"externalData,omitempty"`
}

func (a *AssignField) GetValue(mutable *types.Mutable) (interface{}, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

func (a *AssignField) Validate() error { _ = "STUB: not implemented"; return nil }

type FromMetadata struct {
	// Field specifies which metadata field provides the assigned value. Valid fields are `namespace` and `name`.
	Field Field `json:"field,omitempty"`
}

func (fm *FromMetadata) GetValue(obj types.MetadataGetter) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}

func (fm *FromMetadata) Validate() error { _ = "STUB: not implemented"; return nil }
