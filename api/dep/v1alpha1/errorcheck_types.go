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
	"encoding/json"
	"github.com/mattbaird/jsonpatch"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"k8s.io/apimachinery/pkg/types"
)

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

type ErrorItem struct {
	ErrorDesc   string `json:"errordesc,omitempty"`
	ErrorCode   string `json:"errorcode,omitempty"`
	ErrorLevel  string `json:"errorlevel,omitempty"`
	ResoureName string `json:"resourename,omitempty"`
	RelateName  string `json:"relatename,omitempty"`
}

type ErrorItemSlice []ErrorItem

func (r ErrorCheck) Type() types.PatchType {
	return types.JSONPatchType
}

// Data is the raw data representing the patch.
func (r ErrorCheck) Data(obj runtime.Object) ([]byte, error) {
	odata, err := json.Marshal(&obj)
	if err != nil {
		return nil, err
	}

	mdata, err := json.Marshal(&r)
	if err != nil {
		return nil, err
	}

	jdata, err := jsonpatch.CreatePatch(odata, mdata)
	if err != nil {
		return nil, err
	}
	return json.Marshal(&jdata)
}

// ErrorCheckSpec defines the desired state of ErrorCheck
type ErrorCheckSpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	ErrList []ErrorItem `json:"errlist"`
}

// ErrorCheckStatus defines the observed state of ErrorCheck
type ErrorCheckStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster
	// Important: Run "make" to regenerate code after modifying this file
}

// +kubebuilder:object:root=true

// ErrorCheck is the Schema for the errorchecks API
type ErrorCheck struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   ErrorCheckSpec   `json:"spec,omitempty"`
	Status ErrorCheckStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// ErrorCheckList contains a list of ErrorCheck
type ErrorCheckList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []ErrorCheck `json:"items"`
}

func init() {
	SchemeBuilder.Register(&ErrorCheck{}, &ErrorCheckList{})
}
