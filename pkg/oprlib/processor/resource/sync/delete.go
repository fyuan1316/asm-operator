package sync

import (
	"context"
	"fmt"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/manage/model"
	pkgerrors "github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/api/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

const DeletionErr = "deletion error in resource task"

var FnDelete = func(client client.Client, object model.Object) error {
	//var err error
	//obj := object.DeepCopyObject()
	//err = client.Get(context.Background(), types.NamespacedName{
	//	Namespace: object.GetNamespace(),
	//	Name:      object.GetName(),
	//}, obj)
	//if err != nil && !errors.IsNotFound(err) {
	//	return err
	//}
	err := client.Delete(context.Background(), object)
	if err != nil && !errors.IsNotFound(err) {
		return pkgerrors.Wrap(err,
			fmt.Sprintf("%s [kind=%s,namespace=%s,name=%s]",
				DeletionErr,
				object.GetObjectKind(),
				object.GetNamespace(),
				object.GetName()))
	}
	return nil
}
