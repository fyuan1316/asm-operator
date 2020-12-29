package sync

import (
	"context"
	depv1beta1 "github.com/fyuan1316/asm-operator/api/dep/v1beta1"
	depv1beta2 "github.com/fyuan1316/asm-operator/api/dep/v1beta2"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/manage/model"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var FnCreateCaseMonitorv1beta1 = func(client client.Client, object model.Object) error {
	deploy := depv1beta1.CaseMonitor{}
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
	return nil
}

var FnCreateCaseMonitorv1beta2 = func(client client.Client, object model.Object) error {
	deploy := depv1beta2.CaseMonitor{}
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
	return nil
}
