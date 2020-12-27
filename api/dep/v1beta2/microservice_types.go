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

package v1beta2

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.
type Workload struct {
	Name    string `json:"name"`
	Version string `json:"version,omitempty"`
}

type ServiceItem struct {
	Name             string `json:"name"`
	IsCreateBySystem bool   `json:"iscreatebysystem"`
}

// MicroServiceSpec defines the desired state of MicroService
type MicroServiceSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Deployments  []Workload    `json:"deployments,omitempty"`
	Statefulsets []Workload    `json:"statefulsets,omitempty"`
	Daemonsets   []Workload    `json:"daemonsets,omitempty"`
	Services     []ServiceItem `json:"services,omitempty"`
}

// MicroServiceStatus defines the observed state of MicroService
type MicroServiceStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true
// +kubebuilder:storageversion

// MicroService is the Schema for the microservices API
type MicroService struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   MicroServiceSpec   `json:"spec,omitempty"`
	Status MicroServiceStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// MicroServiceList contains a list of MicroService
type MicroServiceList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []MicroService `json:"items"`
}

func init() {
	SchemeBuilder.Register(&MicroService{}, &MicroServiceList{})
}
