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

type ConnectionPoolSettings_TCPSettings struct {
	// Maximum number of HTTP1 /TCP connections to a destination host. Default 1024.
	MaxConnections int32 `protobuf:"varint,1,opt,name=maxConnections,json=maxConnections,proto3" json:"maxConnections,omitempty"`
	// TCP connection timeout.
	ConnectTimeout string `protobuf:"bytes,2,opt,name=connectTimeout,json=connectTimeout,proto3" json:"connectTimeout,omitempty"`
}

func (m *ConnectionPoolSpec) Reset()         { *m = ConnectionPoolSpec{} }
func (m *ConnectionPoolSpec) String() string { return proto.CompactTextString(m) }
func (*ConnectionPoolSpec) ProtoMessage()    {}

// ConnectionPoolSpec defines the desired state of ConnectionPool
type ConnectionPoolSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file
	Host string                               `protobuf:"bytes,1,opt,name=host,proto3" json:"host,omitempty"`
	Tcp  *ConnectionPoolSettings_TCPSettings  `protobuf:"bytes,1,opt,name=tcp,proto3" json:"tcp,omitempty"`
	Http *ConnectionPoolSettings_HTTPSettings `protobuf:"bytes,2,opt,name=http,proto3" json:"http,omitempty"`
}

// Settings applicable to HTTP1.1/HTTP2/GRPC connections.
type ConnectionPoolSettings_HTTPSettings struct {
	// Maximum number of pending HTTP1 requests to a destination. Default 1024.
	HttpMaxPendingRequests int32 `protobuf:"varint,4,opt,name=httpMaxPendingRequests,json=httpMaxPendingRequests,proto3" json:"httpMaxPendingRequests,omitempty"`
	// Maximum number of HTTP2 requests to a backend. Default 1024.
	HttpMaxRequests int32 `protobuf:"varint,4,opt,name=httpMaxRequests,json=httpMaxRequests,proto3" json:"httpMaxRequests,omitempty"`
	// parameter 1 disables keep alive. Default 0, meaning "unlimited",
	MaxRequestsPerConnection int32 `protobuf:"varint,4,opt,name=maxRequestsPerConnection,json=maxRequestsPerConnection,proto3" json:"maxRequestsPerConnection,omitempty"`
	// cluster at a given time. Defaults to 3.
	MaxRetries int32 `protobuf:"varint,4,opt,name=maxRetries,json=maxRetries,proto3" json:"maxRetries,omitempty"`
}

// ConnectionPoolStatus defines the observed state of ConnectionPool
type ConnectionPoolStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// ConnectionPool is the Schema for the connectionpools API
type ConnectionPool struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ConnectionPoolSpec   `json:"spec,omitempty"`
	Status ConnectionPoolStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ConnectionPoolList contains a list of ConnectionPool
type ConnectionPoolList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ConnectionPool `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ConnectionPool{}, &ConnectionPoolList{})
}
