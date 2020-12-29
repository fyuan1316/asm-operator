package sync

import (
	"context"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/manage/model"
	"k8s.io/apimachinery/pkg/api/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

var FnDelete = func(client client.Client, object model.Object) error {
	err := client.Delete(context.Background(), object)
	if err != nil && !errors.IsNotFound(err) {
		return err
	}
	return nil
}
