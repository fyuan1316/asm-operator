package sync

import (
	"context"
	v1 "github.com/coreos/prometheus-operator/pkg/apis/monitoring/v1"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/manage/model"
	"k8s.io/apimachinery/pkg/api/equality"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var FnServiceMonitor = func(client client.Client, object model.Object) error {
	deploy := v1.ServiceMonitor{}
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
	wanted := object.(*v1.ServiceMonitor)
	if !equality.Semantic.DeepDerivative(wanted.Spec, deploy.Spec) {
		deploy.Spec = wanted.Spec
		if errUpd := client.Update(context.Background(), &deploy); errUpd != nil {
			return errUpd
		}
	}
	return nil
}

var FnPodMonitor = func(client client.Client, object model.Object) error {
	deploy := v1.PodMonitor{}
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
	wanted := object.(*v1.PodMonitor)
	if !equality.Semantic.DeepDerivative(wanted.Spec, deploy.Spec) {
		deploy.Spec = wanted.Spec
		if errUpd := client.Update(context.Background(), &deploy); errUpd != nil {
			return errUpd
		}
	}
	return nil
}
