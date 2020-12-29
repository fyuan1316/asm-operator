package sync

import (
	"context"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/manage/model"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var FnService = func(client client.Client, object model.Object) error {
	deploy := corev1.Service{}
	err := client.Get(context.Background(),
		types.NamespacedName{Namespace: object.GetNamespace(), Name: object.GetName()},
		&deploy,
	)
	if err != nil {
		if errors.IsNotFound(err) {
			errCreate := client.Create(context.Background(), object)
			if errCreate != nil {
				return errCreate
			}
			return nil
		}
		return err
	}
	//update
	wanted := object.(*corev1.Service)
	if !equality.Semantic.DeepDerivative(wanted.Spec, deploy.Spec) {
		deploy.Spec = wanted.Spec
		if errUpd := client.Update(context.Background(), &deploy); errUpd != nil {
			return errUpd
		}
	}
	return nil
}
