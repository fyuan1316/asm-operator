package options

import (
	"github.com/fyuan1316/asm-operator/pkg/oprlib/manage/model"
	"k8s.io/apimachinery/pkg/runtime"
)

type Option func(spec *model.OperatorManage)

func SetFinalizer(id string) Option {
	return func(spec *model.OperatorManage) {
		spec.FinalizerID = id
	}
}
func SetScheme(scheme *runtime.Scheme) Option {
	return func(spec *model.OperatorManage) {
		spec.Scheme = scheme
	}
}
