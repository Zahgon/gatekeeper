package types

// Anything is a struct wrapper around a field of type `interface{}`
// that plays nicely with controller-gen
// +kubebuilder:object:generate=false
// +kubebuilder:validation:Type=""
type Anything struct {
	Value interface{} `json:"-"`
}

func (in *Anything) GetValue() interface{} { _ = "STUB: not implemented"; return nil }

func (in *Anything) UnmarshalJSON(val []byte) error { _ = "STUB: not implemented"; return nil }

// MarshalJSON should be implemented against a value
// per http://stackoverflow.com/questions/21390979/custom-marshaljson-never-gets-called-in-go
// credit to K8s api machinery's RawExtension for finding this.
func (in Anything) MarshalJSON() ([]byte, error) { _ = "STUB: not implemented"; return nil, nil }

func (in *Anything) DeepCopy() *Anything { _ = "STUB: not implemented"; return nil }

func (in *Anything) DeepCopyInto(out *Anything) { _ = "STUB: not implemented"; return }
