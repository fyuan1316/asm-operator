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
	"sort"
	"strconv"
	"time"

	flaggerv1beta1 "github.com/fyuan1316/asm-operator/api/dep/flagger/v1beta1"
	istiov1alpha3 "github.com/fyuan1316/asm-operator/api/dep/istio/v1alpha3"
	corev1 "k8s.io/api/core/v1"
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
)

// EventName is a label for the condition of a event at the current time
type EventName string

const (
	//
	EventNameNewRevision EventName = "NewRevision"
	//
	EventNameRoutingAll EventName = "RoutingAll"
	//
	EventNamePromotionCompleted EventName = "PromotionCompleted"
	//
	EventNameStartingCanaryAnalysis EventName = "StartingCanaryAnalysis"
	//
	EventNameSwitchingWeight EventName = "SwitchingWeight"
	//
	EventNameUpdatePrimary EventName = "UpdatePrimary"
	//
	EventNameStopTrafficMirroring EventName = "StopTrafficMirroring"
	//
	EventNameSkipAnalysis EventName = "SkipAnalysis"
	//
	EventNameInitialized EventName = "Initialized"
	//
	EventNameConfirmRolloutPassed EventName = "ConfirmRolloutPassed"
	//
	EventNameConfirmPromotionPassed EventName = "ConfirmPromotionPassed"
	//
	EventNamePreRolloutPassed EventName = "PreRolloutPassed"
	//
	EventNamePostRolloutPassed EventName = "PostRolloutPassed"
	//
	EventNameRollbackSkip EventName = "RollbackSkip"
	//
	EventNameManualWebhookInvoked EventName = "ManualWebhookInvoked"
	//
	EventNameProgressDeadlineExceeded EventName = "ProgressDeadlineExceeded"
	//
	EventNameWaitingForApproval EventName = "WaitingForApproval"
	//
	EventNameWaitingForPromotionApproval EventName = "WaitingForPromotionApproval"
	//
	EventNamePreRolloutFailed EventName = "PreRolloutFailed"
	//
	EventNameNOIstioMetrics EventName = "NOIstioMetrics"
	//Canary failed
	EventNameCanaryFailed EventName = "CanaryFailed"

	//Canary failed
	EventNameInitializingGenerationFailed EventName = "GeneretionErr"
)

// IsNew just check if the new revision found
func (event *CanaryEvent) IsNew() bool {
	return event.EventName == EventNameNewRevision
}

// IsFinished just check if canary is completed
func (event *CanaryEvent) IsFinished() bool {
	return event.EventName == EventNamePromotionCompleted
}

// IsRunningCanary just check if canary is progressing
func (event *CanaryEvent) IsRunningCanary() bool {
	return event.EventName == EventNameSwitchingWeight
}

// IsInitilization just check if initilization is finished
func (event *CanaryEvent) IsInitilization() bool {
	return event.EventName == EventNameInitialized
}

// IsCanaryFailed just check if canary failed
func (event *CanaryEvent) IsCanaryFailed() bool {
	return event.EventName == EventNameCanaryFailed
}

// IsGenaretionErr
func (event *CanaryEvent) IsGenaretionErr() bool {
	return event.EventName == EventNameInitializingGenerationFailed
}

// EDIT THIS FILE!  THIS IS SCAFFOLDING FOR YOU TO OWN!
// NOTE: json tags are required.  Any new fields you add must have json tags for the fields to be serialized.

// CanaryDeliverySpec defines the desired state of CanaryDelivery
type CanaryDeliverySpec struct {
	// INSERT ADDITIONAL SPEC FIELDS - desired state of cluster
	// Important: Run "make" to regenerate code after modifying this file

	// TargetRef references a target resource
	TargetRef flaggerv1beta1.CrossNamespaceObjectReference `json:"targetRef"`

	// Analysis defines the validation process of a release
	Analysis *CanaryDeliveryAnalysis `json:"analysis,omitempty"`

	// Metric check list for this canary analysis
	// +optional
	Metrics []flaggerv1beta1.CanaryMetric `json:"metrics,omitempty"`

	// CanaryRecord list for canary
	// +optional
	Records []*CanaryRecord `json:"records,omitempty"`

	// current max record id
	// +optional
	CurrentRecordID int `json:"currentRecourdId,omitempty"`
}

// CanaryRecord is used to record how the canary migration happend
type CanaryRecord struct {
	// Id for this record
	Id int `json:"id,omitempty"`

	Name string `json:"name,omitempty"`

	Namespace string `json:"namespace"`

	StartTime string `json:"startTime"`

	EndTime       string `json:"endTime,omitempty"`
	CurrentWeight int    `json:"currentWeight,omitempty"`
	Status        string `json:"status,omitempty"`

	Events []CanaryEvent `json:"events,omitempty"`
}

// CanaryEvent is used for events of canary
type CanaryEvent struct {

	// name for the event
	// +optional
	EventName EventName `json:"eventName,omitempty"`

	// Phase
	// +optional
	Phase string `json:"phase,omitempty"`

	// A/B testing HTTP header match conditions
	// +optional
	EventType string `json:"eventType,omitempty"`

	// A/B testing HTTP header match conditions
	// +optional
	Timestamp string `json:"timestamp,omitempty"`

	// A/B testing HTTP header match conditions
	// +optional
	Message string `json:"message,omitempty"`

	// A/B testing HTTP header match conditions
	// +optional
	Weight int `json:"weight,omitempty"`
}

// CanaryDeliveryAnalysis is used to describe how the analysis should be done
type CanaryDeliveryAnalysis struct {
	// Schedule interval for this canary analysis
	Interval string `json:"interval"`

	// Number of checks to run for A/B Testing and Blue/Green
	// +optional
	Iterations int `json:"iterations,omitempty"`

	//Enable traffic mirroring for Blue/Green
	// +optional
	Mirror bool `json:"mirror,omitempty"`

	// Max traffic percentage routed to canary
	// +optional
	MaxWeight int `json:"maxWeight,omitempty"`

	// Incremental traffic percentage step
	// +optional
	StepWeight int `json:"stepWeight,omitempty"`

	// Max number of failed checks before the canary is terminated
	// +optional
	Threshold int `json:"threshold,omitempty"`

	// A/B testing HTTP header match conditions
	// +optional
	Match []istiov1alpha3.HTTPMatchRequest `json:"match,omitempty"`
}

// CanaryCondition is a status condition for a Canary
type CanaryCondition struct {
	// Type of this condition
	Type string `json:"type"`

	// Status of this condition
	Status corev1.ConditionStatus `json:"status"`

	// LastUpdateTime of this condition
	LastUpdateTime metav1.Time `json:"lastUpdateTime,omitempty"`

	// LastTransitionTime of this condition
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`

	// Reason for the current status of this condition
	Reason string `json:"reason,omitempty"`

	// Message associated with this condition
	Message string `json:"message,omitempty"`
}

// CanaryDeliveryStatus defines the observed state of CanaryDelivery
type CanaryDeliveryStatus struct {
	// INSERT ADDITIONAL STATUS FIELD - define observed state of cluster

	Phase        flaggerv1beta1.CanaryPhase `json:"phase"`
	FailedChecks int                        `json:"failedChecks"`
	CanaryWeight int                        `json:"canaryWeight"`
	Iterations   int                        `json:"iterations"`
	// +optional
	TrackedConfigs *map[string]string `json:"trackedConfigs,omitempty"`
	// +optional
	LastAppliedSpec string `json:"lastAppliedSpec,omitempty"`
	// +optional
	LastPromotedSpec string `json:"lastPromotedSpec,omitempty"`
	// +optional
	LastTransitionTime metav1.Time `json:"lastTransitionTime,omitempty"`

	// +optional
	Conditions []CanaryCondition `json:"conditions,omitempty"`
}

func (cd *CanaryDelivery) MergeRecrods() {
	records := cd.Spec.Records
	sort.SliceStable(records, func(i, j int) bool {
		return records[i].Id > records[j].Id
	})

	if len(records) > 10 {
		cd.Spec.Records = records[:9]
	}

	for _, record := range cd.Spec.Records {
		events := record.Events
		if len(events) > 0 {

			sort.SliceStable(events, func(i, j int) bool {

				return events[i].Timestamp < events[j].Timestamp
			})
		}

	}
}

func (cd *CanaryDelivery) FinishingPreRecrod() {
	records := cd.Spec.Records
	currentrecordId := cd.Spec.CurrentRecordID
	var foundrecord *CanaryRecord

	for _, record := range records {
		if record.Id == currentrecordId {
			foundrecord = record
			continue
		}
	}

	if foundrecord != nil {
		latestEvent := foundrecord.findLatestEvent()
		if len(foundrecord.EndTime) == 0 {
			if latestEvent != nil {
				foundrecord.EndTime = latestEvent.Timestamp
			} else {
				tm := time.Now()
				foundrecord.EndTime = strconv.FormatInt(tm.Unix(), 10)
			}
			foundrecord.CurrentWeight = 0
			foundrecord.Status = (string)(flaggerv1beta1.CanaryPhaseFailed)
		}

		if len(foundrecord.Status) == 0 {

			foundrecord.CurrentWeight = 0
			foundrecord.Status = (string)(flaggerv1beta1.CanaryPhaseSucceeded)

		}

	}

}

func (record *CanaryRecord) findLatestEvent() *CanaryEvent {
	events := record.Events
	if len(events) > 0 {
		latestEvent := events[len(events)-1]

		return &latestEvent
	}
	return nil
}

func (cd *CanaryDelivery) HasLabels(labelKey string) bool {

	if cd.Labels == nil {
		return false
	}
	_, ok := cd.Labels[labelKey]
	return ok
}

// +kubebuilder:object:root=true

// CanaryDelivery is the Schema for the Canarydeliveries API
type CanaryDelivery struct {
	metav1.TypeMeta   `json:",inline"`
	metav1.ObjectMeta `json:"metadata,omitempty"`

	Spec   CanaryDeliverySpec   `json:"spec,omitempty"`
	Status CanaryDeliveryStatus `json:"status,omitempty"`
}

// +kubebuilder:object:root=true

// CanaryDeliveryList contains a list of CanaryDelivery
type CanaryDeliveryList struct {
	metav1.TypeMeta `json:",inline"`
	metav1.ListMeta `json:"metadata,omitempty"`
	Items           []CanaryDelivery `json:"items"`
}

func init() {
	SchemeBuilder.Register(&CanaryDelivery{}, &CanaryDeliveryList{})
}
