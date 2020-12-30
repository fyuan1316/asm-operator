package model

import (
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type PreRun interface {
	PreRun(client.Client) error
}
type PostRun interface {
	PostRun(client.Client) error
}
type PreCheck interface {
	PreCheck(client.Client) (bool, error)
}
type PostCheck interface {
	PostCheck(client.Client) (bool, error)
}
type Runnable interface {
	Run(*OperatorManage) error
}

type HealthCheck interface {
	LiveNess(client.Client) bool
}

//
func CanDoPreCheck(inf interface{}) (PreCheck, bool) {
	if b, ok := inf.(PreCheck); ok {
		return b, true
	}
	return nil, false
}
func CanDoPostCheck(inf interface{}) (PostCheck, bool) {
	if b, ok := inf.(PostCheck); ok {
		return b, true
	}
	return nil, false
}
func CanDoPreRun(inf interface{}) (PreRun, bool) {
	if b, ok := inf.(PreRun); ok {
		return b, true
	}
	return nil, false
}
func CanDoPostRun(inf interface{}) (PostRun, bool) {
	if b, ok := inf.(PostRun); ok {
		return b, true
	}
	return nil, false
}
func CanDoHealthCheck(inf interface{}) (HealthCheck, bool) {
	if b, ok := inf.(HealthCheck); ok {
		return b, true
	}
	return nil, false
}
func CanDoRun(inf interface{}) (Runnable, bool) {
	if b, ok := inf.(Runnable); ok {
		return b, true
	}
	return nil, false
}

type Item interface {
	GetStageName() string
}
type ExecuteItem interface {
	Item
	Runnable
}
type OperationType string

var Operations = struct {
	Provision OperationType
	Deletion  OperationType
}{"Provision", "Deletion"}

type OverrideOperation interface {
	GetOperation() OperationType
	Override(OverrideOperation)
}

//type TypeObjectMeta struct {
//	metav1.TypeMeta
//	metav1.Object
//}
