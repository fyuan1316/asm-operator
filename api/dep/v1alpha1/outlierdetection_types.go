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
	"github.com/gogo/protobuf/proto"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

func (m *OutlierDetectionSpec) Reset()         { *m = OutlierDetectionSpec{} }
func (m *OutlierDetectionSpec) String() string { return proto.CompactTextString(m) }
func (*OutlierDetectionSpec) ProtoMessage()    {}

// OutlierDetectionSpec defines the desired state of OutlierDetection
type OutlierDetectionSpec struct {
	Host               string `protobuf:"bytes,4,opt,name=host,proto3" json:"host,omitempty"`
	ConsecutiveErrors  int32  `protobuf:"varint,4,opt,name=consecutiveErrors,json=consecutiveErrors,proto3" json:"consecutiveErrors,omitempty"`
	Interval           string `protobuf:"bytes,4,opt,name=interval,proto3" json:"interval,omitempty"`
	BaseEjectionTime   string `protobuf:"bytes,4,opt,name=baseEjectionTime,json=baseEjectionTime,proto3" json:"baseEjectionTime,omitempty"`
	MaxEjectionPercent int32  `protobuf:"varint,4,opt,name=maxEjectionPercent,json=maxEjectionPercent,proto3" json:"maxEjectionPercent,omitempty"`
}

// OutlierDetectionStatus defines the observed state of OutlierDetection
type OutlierDetectionStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// OutlierDetection is the Schema for the outlierdetections API
type OutlierDetection struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   OutlierDetectionSpec   `json:"spec,omitempty"`
	Status OutlierDetectionStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// OutlierDetectionList contains a list of OutlierDetection
type OutlierDetectionList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []OutlierDetection `json:"items"`
}

func init() {
	SchemeBuilder.Register(&OutlierDetection{}, &OutlierDetectionList{})
}
