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
	"github.com/gogo/protobuf/proto"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

//type LoadBalancerSettings struct {
//	// Upstream load balancing policy.
//	LbPolicy LbPolicy `json:",inline"`
//}
type LbPolicy struct {
	Simple         *LoadBalancerSettings_Simple         `json:"simple,omitempty" protobuf:"varint,1,opt,name=simple"`
	ConsistentHash *LoadBalancerSettings_ConsistentHash `json:"consistentHash,omitempty" protobuf:"varint,1,opt,name=consistentHash"`
}

// Simple:ROUND_ROBIN/LEAST_CONN/RANDOM
type LoadBalancerSettings_Simple string
type LoadBalancerSettings_ConsistentHash struct {
	HttpHeaderName  *string     `json:"httpHeaderName,omitempty" protobuf:"bytes,1,opt,name=httpHeaderName,json=httpHeaderName,proto3,oneof"`
	UseSourceIp     *bool       `json:"useSourceIp,omitempty" protobuf:"varint,3,opt,name=useSourceIp,json=useSourceIp,proto3,oneof"`
	HttpCookie      *HTTPCookie `json:"httpCookie,omitempty" protobuf:"bytes,2,opt,name=httpCookie"`
	MinimumRingSize uint64      `json:"minimumRingSize,omitempty" protobuf:"varint,4,opt,name=minimumRingSize,json=minimumRingSize,proto3"`
}
type HTTPCookie struct {
	Name string `protobuf:"bytes,1,opt,name=name,proto3" json:"name,omitempty"`
	Path string `protobuf:"bytes,2,opt,name=path,proto3" json:"path,omitempty"`
	Ttl  string `protobuf:"bytes,3,opt,name=ttl,stdduration" json:"ttl,omitempty"`
}

func (m *LoadBalancerSpec) Reset()         { *m = LoadBalancerSpec{} }
func (m *LoadBalancerSpec) String() string { return proto.CompactTextString(m) }
func (*LoadBalancerSpec) ProtoMessage()    {}

// LoadBalancerSpec defines the desired state of LoadBalancer
type LoadBalancerSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Host string `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	//LoadBalancer *LoadBalancerSettings `protobuf:"bytes,1,opt,name=loadBalancer,json=loadBalancer" json:"loadBalancer,omitempty"`
	Simple         *LoadBalancerSettings_Simple         `json:"simple,omitempty" protobuf:"varint,1,opt,name=simple"`
	ConsistentHash *LoadBalancerSettings_ConsistentHash `json:"consistentHash,omitempty" protobuf:"varint,1,opt,name=consistentHash"`
}

// LoadBalancerStatus defines the observed state of LoadBalancer
type LoadBalancerStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

func (m *LoadBalancer) Reset()         { *m = LoadBalancer{} }
func (m *LoadBalancer) String() string { return proto.CompactTextString(m) }
func (*LoadBalancer) ProtoMessage()    {}

// +kubebuilder:object:root=true

// LoadBalancer is the Schema for the loadbalancers API
type LoadBalancer struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   LoadBalancerSpec   `json:"spec,omitempty"`
	Status LoadBalancerStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// LoadBalancerList contains a list of LoadBalancer
type LoadBalancerList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []LoadBalancer `json:"items"`
}

func init() {
	SchemeBuilder.Register(&LoadBalancer{}, &LoadBalancerList{})
}
