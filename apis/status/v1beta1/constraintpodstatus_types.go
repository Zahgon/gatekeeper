/*

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1beta1

import (
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/apis/meta/v1/unstructured"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

// ConstraintsGroup is the API Group for Gatekeeper Constraints.
const ConstraintsGroup = "constraints.gatekeeper.sh"

// ConstraintPodStatusStatus defines the observed state of ConstraintPodStatus.
type ConstraintPodStatusStatus struct {
	// Important: Run "make" to regenerate code after modifying this file

	ID string `json:"id,omitempty"`
	// Storing the constraint UID allows us to detect drift, such as
	// when a constraint has been recreated after its CRD was deleted
	// out from under it, interrupting the watch
	ConstraintUID           types.UID                `json:"constraintUID,omitempty"`
	Operations              []string                 `json:"operations,omitempty"`
	Enforced                bool                     `json:"enforced,omitempty"`
	Errors                  []Error                  `json:"errors,omitempty"`
	ObservedGeneration      int64                    `json:"observedGeneration,omitempty"`
	EnforcementPointsStatus []EnforcementPointStatus `json:"enforcementPointsStatus,omitempty"`
}

// Error represents a single error caught while adding a constraint to engine.
type Error struct {
	Code     string `json:"code"`
	Message  string `json:"message"`
	Location string `json:"location,omitempty"`
}

// EnforcementPointStatus represents the status of a single enforcement point.
type EnforcementPointStatus struct {
	EnforcementPoint   string `json:"enforcementPoint"`
	State              string `json:"state"`
	Message            string `json:"message,omitempty"`
	ObservedGeneration int64  `json:"observedGeneration,omitempty"`
}

// +kubebuilder:object:root=true
// +kubebuilder:resource:scope=Namespaced

// ConstraintPodStatus is the Schema for the constraintpodstatuses API.
type ConstraintPodStatus struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Status ConstraintPodStatusStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConstraintPodStatusList contains a list of ConstraintPodStatus.
type ConstraintPodStatusList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConstraintPodStatus `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ConstraintPodStatus{}, &ConstraintPodStatusList{})
}

// NewConstraintStatusForPod returns a constraint status object
// that has been initialized with the bare minimum of fields to make it functional
// with the constraint status controller.
func NewConstraintStatusForPod(pod *corev1.Pod, constraint *unstructured.Unstructured, scheme *runtime.Scheme) (*ConstraintPodStatus, error) {
	_ = "STUB: not implemented"
	return nil, nil
}

// the template name is the lower-case of the constraint kind

// Skip OwnerReference in remote cluster mode

// KeyForConstraint returns a unique status object name given the Pod ID and
// a constraint object.
func KeyForConstraint(id string, constraint *unstructured.Unstructured) (string, error) {
	_ = "STUB: not implemented"
	// We don't need to worry that lower-casing the kind will cause a collision because
	// the constraint framework requires resource == lower-case kind. We must do this
	// because K8s requires all lowercase letters for resource names
	return "", nil
}
