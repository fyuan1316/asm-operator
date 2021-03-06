// +build !ignore_autogenerated

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

// Code generated by controller-gen. DO NOT EDIT.

package v1alpha1

import (
	"github.com/fyuan1316/asm-operator/api/dep/flagger/v1beta1"
	"github.com/fyuan1316/asm-operator/api/dep/istio/v1alpha3"
	"k8s.io/apimachinery/pkg/runtime"
)

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *AllowListItem) DeepCopyInto(out *AllowListItem) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new AllowListItem.
func (in *AllowListItem) DeepCopy() *AllowListItem {
	if in == nil {
		return nil
	}
	out := new(AllowListItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanaryAnalysis) DeepCopyInto(out *CanaryAnalysis) {
	*out = *in
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = make([]CanaryMetric, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Webhooks != nil {
		in, out := &in.Webhooks, &out.Webhooks
		*out = make([]CanaryWebhook, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Match != nil {
		in, out := &in.Match, &out.Match
		*out = make([]v1alpha3.HTTPMatchRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanaryAnalysis.
func (in *CanaryAnalysis) DeepCopy() *CanaryAnalysis {
	if in == nil {
		return nil
	}
	out := new(CanaryAnalysis)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanaryCondition) DeepCopyInto(out *CanaryCondition) {
	*out = *in
	in.LastUpdateTime.DeepCopyInto(&out.LastUpdateTime)
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanaryCondition.
func (in *CanaryCondition) DeepCopy() *CanaryCondition {
	if in == nil {
		return nil
	}
	out := new(CanaryCondition)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanaryDelivery) DeepCopyInto(out *CanaryDelivery) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	in.Status.DeepCopyInto(&out.Status)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanaryDelivery.
func (in *CanaryDelivery) DeepCopy() *CanaryDelivery {
	if in == nil {
		return nil
	}
	out := new(CanaryDelivery)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CanaryDelivery) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanaryDeliveryAnalysis) DeepCopyInto(out *CanaryDeliveryAnalysis) {
	*out = *in
	if in.Match != nil {
		in, out := &in.Match, &out.Match
		*out = make([]v1alpha3.HTTPMatchRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanaryDeliveryAnalysis.
func (in *CanaryDeliveryAnalysis) DeepCopy() *CanaryDeliveryAnalysis {
	if in == nil {
		return nil
	}
	out := new(CanaryDeliveryAnalysis)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanaryDeliveryList) DeepCopyInto(out *CanaryDeliveryList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CanaryDelivery, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanaryDeliveryList.
func (in *CanaryDeliveryList) DeepCopy() *CanaryDeliveryList {
	if in == nil {
		return nil
	}
	out := new(CanaryDeliveryList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CanaryDeliveryList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanaryDeliverySpec) DeepCopyInto(out *CanaryDeliverySpec) {
	*out = *in
	out.TargetRef = in.TargetRef
	if in.Analysis != nil {
		in, out := &in.Analysis, &out.Analysis
		*out = new(CanaryDeliveryAnalysis)
		(*in).DeepCopyInto(*out)
	}
	if in.Metrics != nil {
		in, out := &in.Metrics, &out.Metrics
		*out = make([]v1beta1.CanaryMetric, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Records != nil {
		in, out := &in.Records, &out.Records
		*out = make([]*CanaryRecord, len(*in))
		for i := range *in {
			if (*in)[i] != nil {
				in, out := &(*in)[i], &(*out)[i]
				*out = new(CanaryRecord)
				(*in).DeepCopyInto(*out)
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanaryDeliverySpec.
func (in *CanaryDeliverySpec) DeepCopy() *CanaryDeliverySpec {
	if in == nil {
		return nil
	}
	out := new(CanaryDeliverySpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanaryDeliveryStatus) DeepCopyInto(out *CanaryDeliveryStatus) {
	*out = *in
	if in.TrackedConfigs != nil {
		in, out := &in.TrackedConfigs, &out.TrackedConfigs
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
	in.LastTransitionTime.DeepCopyInto(&out.LastTransitionTime)
	if in.Conditions != nil {
		in, out := &in.Conditions, &out.Conditions
		*out = make([]CanaryCondition, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanaryDeliveryStatus.
func (in *CanaryDeliveryStatus) DeepCopy() *CanaryDeliveryStatus {
	if in == nil {
		return nil
	}
	out := new(CanaryDeliveryStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanaryEvent) DeepCopyInto(out *CanaryEvent) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanaryEvent.
func (in *CanaryEvent) DeepCopy() *CanaryEvent {
	if in == nil {
		return nil
	}
	out := new(CanaryEvent)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanaryMetric) DeepCopyInto(out *CanaryMetric) {
	*out = *in
	if in.ThresholdRange != nil {
		in, out := &in.ThresholdRange, &out.ThresholdRange
		*out = new(CanaryThresholdRange)
		(*in).DeepCopyInto(*out)
	}
	if in.TemplateRef != nil {
		in, out := &in.TemplateRef, &out.TemplateRef
		*out = new(CrossNamespaceObjectReference)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanaryMetric.
func (in *CanaryMetric) DeepCopy() *CanaryMetric {
	if in == nil {
		return nil
	}
	out := new(CanaryMetric)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanaryRecord) DeepCopyInto(out *CanaryRecord) {
	*out = *in
	if in.Events != nil {
		in, out := &in.Events, &out.Events
		*out = make([]CanaryEvent, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanaryRecord.
func (in *CanaryRecord) DeepCopy() *CanaryRecord {
	if in == nil {
		return nil
	}
	out := new(CanaryRecord)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanaryService) DeepCopyInto(out *CanaryService) {
	*out = *in
	out.TargetPort = in.TargetPort
	if in.Gateways != nil {
		in, out := &in.Gateways, &out.Gateways
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.Hosts != nil {
		in, out := &in.Hosts, &out.Hosts
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
	if in.TrafficPolicy != nil {
		in, out := &in.TrafficPolicy, &out.TrafficPolicy
		*out = new(v1alpha3.TrafficPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.Match != nil {
		in, out := &in.Match, &out.Match
		*out = make([]v1alpha3.HTTPMatchRequest, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
	if in.Rewrite != nil {
		in, out := &in.Rewrite, &out.Rewrite
		*out = new(v1alpha3.HTTPRewrite)
		**out = **in
	}
	if in.Retries != nil {
		in, out := &in.Retries, &out.Retries
		*out = new(v1alpha3.HTTPRetry)
		**out = **in
	}
	if in.Headers != nil {
		in, out := &in.Headers, &out.Headers
		*out = new(v1alpha3.Headers)
		(*in).DeepCopyInto(*out)
	}
	if in.CorsPolicy != nil {
		in, out := &in.CorsPolicy, &out.CorsPolicy
		*out = new(v1alpha3.CorsPolicy)
		(*in).DeepCopyInto(*out)
	}
	if in.Backends != nil {
		in, out := &in.Backends, &out.Backends
		*out = make([]string, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanaryService.
func (in *CanaryService) DeepCopy() *CanaryService {
	if in == nil {
		return nil
	}
	out := new(CanaryService)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanaryTemplate) DeepCopyInto(out *CanaryTemplate) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanaryTemplate.
func (in *CanaryTemplate) DeepCopy() *CanaryTemplate {
	if in == nil {
		return nil
	}
	out := new(CanaryTemplate)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CanaryTemplate) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanaryTemplateList) DeepCopyInto(out *CanaryTemplateList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]CanaryTemplate, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanaryTemplateList.
func (in *CanaryTemplateList) DeepCopy() *CanaryTemplateList {
	if in == nil {
		return nil
	}
	out := new(CanaryTemplateList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *CanaryTemplateList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanaryTemplateSpec) DeepCopyInto(out *CanaryTemplateSpec) {
	*out = *in
	out.TargetRef = in.TargetRef
	if in.AutoscalerRef != nil {
		in, out := &in.AutoscalerRef, &out.AutoscalerRef
		*out = new(CrossNamespaceObjectReference)
		**out = **in
	}
	if in.IngressRef != nil {
		in, out := &in.IngressRef, &out.IngressRef
		*out = new(CrossNamespaceObjectReference)
		**out = **in
	}
	in.Service.DeepCopyInto(&out.Service)
	if in.Analysis != nil {
		in, out := &in.Analysis, &out.Analysis
		*out = new(CanaryAnalysis)
		(*in).DeepCopyInto(*out)
	}
	if in.CanaryAnalysis != nil {
		in, out := &in.CanaryAnalysis, &out.CanaryAnalysis
		*out = new(CanaryAnalysis)
		(*in).DeepCopyInto(*out)
	}
	if in.ProgressDeadlineSeconds != nil {
		in, out := &in.ProgressDeadlineSeconds, &out.ProgressDeadlineSeconds
		*out = new(int32)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanaryTemplateSpec.
func (in *CanaryTemplateSpec) DeepCopy() *CanaryTemplateSpec {
	if in == nil {
		return nil
	}
	out := new(CanaryTemplateSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanaryThresholdRange) DeepCopyInto(out *CanaryThresholdRange) {
	*out = *in
	if in.Min != nil {
		in, out := &in.Min, &out.Min
		*out = new(float64)
		**out = **in
	}
	if in.Max != nil {
		in, out := &in.Max, &out.Max
		*out = new(float64)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanaryThresholdRange.
func (in *CanaryThresholdRange) DeepCopy() *CanaryThresholdRange {
	if in == nil {
		return nil
	}
	out := new(CanaryThresholdRange)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanaryWebhook) DeepCopyInto(out *CanaryWebhook) {
	*out = *in
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = new(map[string]string)
		if **in != nil {
			in, out := *in, *out
			*out = make(map[string]string, len(*in))
			for key, val := range *in {
				(*out)[key] = val
			}
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanaryWebhook.
func (in *CanaryWebhook) DeepCopy() *CanaryWebhook {
	if in == nil {
		return nil
	}
	out := new(CanaryWebhook)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CanaryWebhookPayload) DeepCopyInto(out *CanaryWebhookPayload) {
	*out = *in
	if in.Metadata != nil {
		in, out := &in.Metadata, &out.Metadata
		*out = make(map[string]string, len(*in))
		for key, val := range *in {
			(*out)[key] = val
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CanaryWebhookPayload.
func (in *CanaryWebhookPayload) DeepCopy() *CanaryWebhookPayload {
	if in == nil {
		return nil
	}
	out := new(CanaryWebhookPayload)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionPool) DeepCopyInto(out *ConnectionPool) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionPool.
func (in *ConnectionPool) DeepCopy() *ConnectionPool {
	if in == nil {
		return nil
	}
	out := new(ConnectionPool)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConnectionPool) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionPoolList) DeepCopyInto(out *ConnectionPoolList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ConnectionPool, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionPoolList.
func (in *ConnectionPoolList) DeepCopy() *ConnectionPoolList {
	if in == nil {
		return nil
	}
	out := new(ConnectionPoolList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ConnectionPoolList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionPoolSettings_HTTPSettings) DeepCopyInto(out *ConnectionPoolSettings_HTTPSettings) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionPoolSettings_HTTPSettings.
func (in *ConnectionPoolSettings_HTTPSettings) DeepCopy() *ConnectionPoolSettings_HTTPSettings {
	if in == nil {
		return nil
	}
	out := new(ConnectionPoolSettings_HTTPSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionPoolSettings_TCPSettings) DeepCopyInto(out *ConnectionPoolSettings_TCPSettings) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionPoolSettings_TCPSettings.
func (in *ConnectionPoolSettings_TCPSettings) DeepCopy() *ConnectionPoolSettings_TCPSettings {
	if in == nil {
		return nil
	}
	out := new(ConnectionPoolSettings_TCPSettings)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionPoolSpec) DeepCopyInto(out *ConnectionPoolSpec) {
	*out = *in
	if in.Tcp != nil {
		in, out := &in.Tcp, &out.Tcp
		*out = new(ConnectionPoolSettings_TCPSettings)
		**out = **in
	}
	if in.Http != nil {
		in, out := &in.Http, &out.Http
		*out = new(ConnectionPoolSettings_HTTPSettings)
		**out = **in
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionPoolSpec.
func (in *ConnectionPoolSpec) DeepCopy() *ConnectionPoolSpec {
	if in == nil {
		return nil
	}
	out := new(ConnectionPoolSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ConnectionPoolStatus) DeepCopyInto(out *ConnectionPoolStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ConnectionPoolStatus.
func (in *ConnectionPoolStatus) DeepCopy() *ConnectionPoolStatus {
	if in == nil {
		return nil
	}
	out := new(ConnectionPoolStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *CrossNamespaceObjectReference) DeepCopyInto(out *CrossNamespaceObjectReference) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new CrossNamespaceObjectReference.
func (in *CrossNamespaceObjectReference) DeepCopy() *CrossNamespaceObjectReference {
	if in == nil {
		return nil
	}
	out := new(CrossNamespaceObjectReference)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *DestMatch) DeepCopyInto(out *DestMatch) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new DestMatch.
func (in *DestMatch) DeepCopy() *DestMatch {
	if in == nil {
		return nil
	}
	out := new(DestMatch)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ErrorCheck) DeepCopyInto(out *ErrorCheck) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ErrorCheck.
func (in *ErrorCheck) DeepCopy() *ErrorCheck {
	if in == nil {
		return nil
	}
	out := new(ErrorCheck)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ErrorCheck) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ErrorCheckList) DeepCopyInto(out *ErrorCheckList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]ErrorCheck, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ErrorCheckList.
func (in *ErrorCheckList) DeepCopy() *ErrorCheckList {
	if in == nil {
		return nil
	}
	out := new(ErrorCheckList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *ErrorCheckList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ErrorCheckSpec) DeepCopyInto(out *ErrorCheckSpec) {
	*out = *in
	if in.ErrList != nil {
		in, out := &in.ErrList, &out.ErrList
		*out = make([]ErrorItem, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ErrorCheckSpec.
func (in *ErrorCheckSpec) DeepCopy() *ErrorCheckSpec {
	if in == nil {
		return nil
	}
	out := new(ErrorCheckSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ErrorCheckStatus) DeepCopyInto(out *ErrorCheckStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ErrorCheckStatus.
func (in *ErrorCheckStatus) DeepCopy() *ErrorCheckStatus {
	if in == nil {
		return nil
	}
	out := new(ErrorCheckStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *ErrorItem) DeepCopyInto(out *ErrorItem) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ErrorItem.
func (in *ErrorItem) DeepCopy() *ErrorItem {
	if in == nil {
		return nil
	}
	out := new(ErrorItem)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in ErrorItemSlice) DeepCopyInto(out *ErrorItemSlice) {
	{
		in := &in
		*out = make(ErrorItemSlice, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new ErrorItemSlice.
func (in ErrorItemSlice) DeepCopy() ErrorItemSlice {
	if in == nil {
		return nil
	}
	out := new(ErrorItemSlice)
	in.DeepCopyInto(out)
	return *out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OutlierDetection) DeepCopyInto(out *OutlierDetection) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	out.Spec = in.Spec
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OutlierDetection.
func (in *OutlierDetection) DeepCopy() *OutlierDetection {
	if in == nil {
		return nil
	}
	out := new(OutlierDetection)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OutlierDetection) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OutlierDetectionList) DeepCopyInto(out *OutlierDetectionList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]OutlierDetection, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OutlierDetectionList.
func (in *OutlierDetectionList) DeepCopy() *OutlierDetectionList {
	if in == nil {
		return nil
	}
	out := new(OutlierDetectionList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *OutlierDetectionList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OutlierDetectionSpec) DeepCopyInto(out *OutlierDetectionSpec) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OutlierDetectionSpec.
func (in *OutlierDetectionSpec) DeepCopy() *OutlierDetectionSpec {
	if in == nil {
		return nil
	}
	out := new(OutlierDetectionSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *OutlierDetectionStatus) DeepCopyInto(out *OutlierDetectionStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new OutlierDetectionStatus.
func (in *OutlierDetectionStatus) DeepCopy() *OutlierDetectionStatus {
	if in == nil {
		return nil
	}
	out := new(OutlierDetectionStatus)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WhiteList) DeepCopyInto(out *WhiteList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ObjectMeta.DeepCopyInto(&out.ObjectMeta)
	in.Spec.DeepCopyInto(&out.Spec)
	out.Status = in.Status
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WhiteList.
func (in *WhiteList) DeepCopy() *WhiteList {
	if in == nil {
		return nil
	}
	out := new(WhiteList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WhiteList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WhiteListList) DeepCopyInto(out *WhiteListList) {
	*out = *in
	out.TypeMeta = in.TypeMeta
	in.ListMeta.DeepCopyInto(&out.ListMeta)
	if in.Items != nil {
		in, out := &in.Items, &out.Items
		*out = make([]WhiteList, len(*in))
		for i := range *in {
			(*in)[i].DeepCopyInto(&(*out)[i])
		}
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WhiteListList.
func (in *WhiteListList) DeepCopy() *WhiteListList {
	if in == nil {
		return nil
	}
	out := new(WhiteListList)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyObject is an autogenerated deepcopy function, copying the receiver, creating a new runtime.Object.
func (in *WhiteListList) DeepCopyObject() runtime.Object {
	if c := in.DeepCopy(); c != nil {
		return c
	}
	return nil
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WhiteListSpec) DeepCopyInto(out *WhiteListSpec) {
	*out = *in
	out.DestMatch = in.DestMatch
	if in.AllowList != nil {
		in, out := &in.AllowList, &out.AllowList
		*out = make([]AllowListItem, len(*in))
		copy(*out, *in)
	}
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WhiteListSpec.
func (in *WhiteListSpec) DeepCopy() *WhiteListSpec {
	if in == nil {
		return nil
	}
	out := new(WhiteListSpec)
	in.DeepCopyInto(out)
	return out
}

// DeepCopyInto is an autogenerated deepcopy function, copying the receiver, writing into out. in must be non-nil.
func (in *WhiteListStatus) DeepCopyInto(out *WhiteListStatus) {
	*out = *in
}

// DeepCopy is an autogenerated deepcopy function, copying the receiver, creating a new WhiteListStatus.
func (in *WhiteListStatus) DeepCopy() *WhiteListStatus {
	if in == nil {
		return nil
	}
	out := new(WhiteListStatus)
	in.DeepCopyInto(out)
	return out
}
