package model

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type OperatorManage struct {
	K8sClient client.Client
	CR        Object
	Options   *OperatorOptions
	//TODO fy event
}

func (m *OperatorManage) GetEditableCR() runtime.Object {
	return m.CR.DeepCopyObject()
}

type Object interface {
	runtime.Object
	metav1.Object
}
