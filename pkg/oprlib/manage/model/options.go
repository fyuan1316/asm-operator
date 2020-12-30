package model

import (
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type StatusUpdaterFunc func(obj Object, client client.Client) func(isReady, isHealthy bool) error
type Option func(spec *OperatorOptions)
type OperatorOptions struct {
	Scheme        *runtime.Scheme
	FinalizerID   string
	StatusUpdater StatusUpdaterFunc
}

func SetFinalizer(id string) Option {
	return func(spec *OperatorOptions) {
		spec.FinalizerID = id
	}
}
func SetScheme(scheme *runtime.Scheme) Option {
	return func(spec *OperatorOptions) {
		spec.Scheme = scheme
	}
}
func SetStatusUpdater(updater StatusUpdaterFunc) Option {
	return func(spec *OperatorOptions) {
		spec.StatusUpdater = updater
	}
}
