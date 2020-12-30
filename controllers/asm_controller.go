/*


Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package controllers

import (
	"context"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/manage"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/manage/model"
	"github.com/fyuan1316/asm-operator/pkg/task/entry"
	pkgerrors "github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/dynamic"
	"k8s.io/client-go/rest"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"sync"

	operatorv1alpha1 "github.com/fyuan1316/asm-operator/api/v1alpha1"
	"github.com/go-logr/logr"
	"k8s.io/apimachinery/pkg/runtime"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

// AsmReconciler reconciles a Asm object
type AsmReconciler struct {
	client.Client
	DynamicClient dynamic.Interface
	Config        *rest.Config
	Log           logr.Logger
	Scheme        *runtime.Scheme
}

var once = sync.Once{}
var mgr *model.OperatorManage
var (
	provisionTasks [][]model.ExecuteItem
	deletionTasks  [][]model.ExecuteItem
)

// +kubebuilder:rbac:groups=operator.alauda.io,resources=asms,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=operator.alauda.io,resources=asms/status,verbs=get;update;patch
const finalizerID = "asms.operator.alauda.io"

func (r *AsmReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	var err error
	_ = context.Background()
	log := r.Log.WithValues("asm", req.NamespacedName)
	_ = log

	instance := &operatorv1alpha1.Asm{}
	err = r.Get(context.TODO(), req.NamespacedName, instance)
	if err != nil {
		if errors.IsNotFound(err) {
			// CR not found, return.  Created objects are automatically garbage collected.
			// For additional cleanup logic use finalizers.
			return reconcile.Result{}, nil
		}
		// Error reading the object - requeue the request.
		return reconcile.Result{}, err
	}
	once.Do(func() {
		mgr = manage.NewOperatorManage(
			r.Client, instance,
			model.SetScheme(r.Scheme),
			model.SetFinalizer(finalizerID),
			model.SetStatusUpdater(asmOperatorStatusUpdate))

		provisionTasks, deletionTasks = entry.GetOperatorStages()
	})
	result, err := mgr.Reconcile(provisionTasks, deletionTasks)
	if err != nil {
		log.Error(err, "Reconcile err")
	}

	return result, nil
}

var asmOperatorStatusUpdate = func(obj runtime.Object, client client.Client) func(isReady, isHealthy bool) error {
	return func(isReady, isHealthy bool) error {
		var asm *operatorv1alpha1.Asm
		var ok bool
		if asm, ok = obj.(*operatorv1alpha1.Asm); !ok {
			return pkgerrors.New("asmOperatorStatusUpdate cast model.Object to operatorv1alpha1.Asm error")
		}
		asmCopy := asm.DeepCopy()
		asmCopy.Status.SetState(isReady, isHealthy)
		if asm.Status.State != asmCopy.Status.State {
			if updErr := client.Status().Update(context.Background(), asmCopy); updErr != nil {
				if errors.IsConflict(updErr) {
					cur := &operatorv1alpha1.Asm{}
					if err := client.Get(
						context.Background(),
						types.NamespacedName{
							Namespace: asmCopy.GetNamespace(),
							Name:      asmCopy.GetName(),
						},
						cur,
					); err != nil {
						return err
					}
					retryObj := cur.DeepCopy()
					retryObj.Status.SetState(isReady, isHealthy)
					if updErr2 := client.Status().Update(context.Background(), retryObj); updErr2 != nil {
						return pkgerrors.Wrap(updErr2, "reUpdate AsmStatus error")
					}
					return nil
				}
				return pkgerrors.Wrap(updErr, "update AsmStatus error")
			}
		}
		return nil
	}

}

func (r *AsmReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&operatorv1alpha1.Asm{}).
		Complete(r)
}
