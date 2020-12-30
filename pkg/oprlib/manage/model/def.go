package model

import (
	metav1 "k8s.io/apimachinery/pkg/apis/meta/v1"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type OperatorManage struct {
	K8sClient client.Client
	CR        Object
	//Scheme        *runtime.Scheme
	//FinalizerID   string
	//statusUpdater func(obj Object, client client.Client) func(isReady, isHealthy bool) error
	Options *OperatorOptions
}

type Object interface {
	runtime.Object
	metav1.Object
}
