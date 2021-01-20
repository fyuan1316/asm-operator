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

package v1alpha1

import (
	"github.com/fyuan1316/operatorlib/api"
	"github.com/fyuan1316/operatorlib/manage/model"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// AsmSpec defines the desired state of Asm
type AsmSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	api.OperatorSpec `json:",inline"`
}

// AsmStatus defines the observed state of Asm
type AsmStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	api.OperatorStatus `json:",inline"`
}

// +kubebuilder:object:root=true
// +kubebuilder:subresource:status
// +kubebuilder:resource:scope=Cluster

// Asm is the Schema for the asms API
type Asm struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   AsmSpec   `json:"spec,omitempty"`
	Status AsmStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// AsmList contains a list of Asm
type AsmList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []Asm `json:"items"`
}

func init() {
	SchemeBuilder.Register(&Asm{}, &AsmList{})
}

var _ model.CommonOperator = &Asm{}

func (in Asm) GetOperatorParams() (map[string]interface{}, error) {
	return in.Spec.OperatorSpec.GetOperatorParams()
}
