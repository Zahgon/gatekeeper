package v1beta1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

// ExpansionTemplatePodStatusStatus defines the observed state of ExpansionTemplatePodStatus.
type ExpansionTemplatePodStatusStatus struct {
	// Important: Run "make" to regenerate code after modifying this file
	ID                 string                    `json:"id,omitempty"`
	TemplateUID        types.UID                 `json:"templateUID,omitempty"`
	Operations         []string                  `json:"operations,omitempty"`
	ObservedGeneration int64                     `json:"observedGeneration,omitempty"`
	Errors             []*ExpansionTemplateError `json:"errors,omitempty"`
}

// +kubebuilder:object:generate=true

type ExpansionTemplateError struct {
	Type    string `json:"type,omitempty"`
	Message string `json:"message"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced

// ExpansionTemplatePodStatus is the Schema for the expansiontemplatepodstatuses API.
type ExpansionTemplatePodStatus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Status ExpansionTemplatePodStatusStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ExpansionTemplatePodStatusList contains a list of ExpansionTemplatePodStatus.
type ExpansionTemplatePodStatusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ExpansionTemplatePodStatus `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ExpansionTemplatePodStatus{}, &ExpansionTemplatePodStatusList{})
}

// NewExpansionTemplateStatusForPod returns an expansion template status object
// that has been initialized with the bare minimum of fields to make it functional
// with the expansion template status controller.
func NewExpansionTemplateStatusForPod(pod *corev1.Pod, templateName string, scheme *runtime.Scheme) (*ExpansionTemplatePodStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// Skip OwnerReference in remote cluster mode

// KeyForExpansionTemplate returns a unique status object name given the Pod ID and
// a template object.
func KeyForExpansionTemplate(id string, templateName string) (string, error) {
	_ = "STUB: not implemented"
	return "", nil
}
