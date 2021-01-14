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
	"fmt"
	asmerrors "github.com/fyuan1316/asm-operator/pkg/errors"
	"github.com/fyuan1316/asm-operator/pkg/task/entry"
	"github.com/fyuan1316/operatorlib/equality"
	"github.com/fyuan1316/operatorlib/event"
	"github.com/fyuan1316/operatorlib/manage"
	"github.com/fyuan1316/operatorlib/manage/model"
	"k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	"k8s.io/client-go/tools/record"
	"k8s.io/client-go/util/retry"
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
	//DynamicClient dynamic.Interface
	//Config        *rest.Config
	Log      logr.Logger
	Scheme   *runtime.Scheme
	Recorder record.EventRecorder
}

var once = sync.Once{}
var mgr *manage.OperatorManage
var (
	provisionTasks [][]model.ExecuteItem
	deletionTasks  [][]model.ExecuteItem
)

// +kubebuilder:rbac:groups=operator.alauda.io,resources=asms,verbs=get;list;watch;create;update;patch;delete
// +kubebuilder:rbac:groups=operator.alauda.io,resources=asms/status,verbs=get;update;patch
const finalizerID = "asms.operator.alauda.io"

func (r *AsmReconciler) Reconcile(req ctrl.Request) (ctrl.Result, error) {
	var err error
	log := r.Log.WithValues("asm", req.NamespacedName)
	log.Info(fmt.Sprintf("Starting reconcile loop for %v", req.NamespacedName))
	defer log.Info(fmt.Sprintf("Finish reconcile loop for %v", req.NamespacedName))

	instance := &operatorv1alpha1.Asm{}
	err = r.Get(context.Background(), req.NamespacedName, instance)
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
			r.Client,
			manage.SetScheme(r.Scheme),
			manage.SetRecorder(r.Recorder),
			manage.SetFinalizer(finalizerID),
			manage.SetStatusUpdater(asmOperatorStatusUpdater3))

		provisionTasks, deletionTasks = entry.GetOperatorStages()
	})
	result, err := mgr.Reconcile(instance, provisionTasks, deletionTasks)
	if err != nil {
		log.Error(err, "Reconcile err")
		r.Recorder.Event(instance, event.WarningEvent, asmerrors.ReconcileError, err.Error())
		return result, err
	}

	return result, nil
}

var asmOperatorStatusUpdater3 = func(reqCtx *model.OperatorContext, statusCtx *model.StatusContext) error {
	asm := &operatorv1alpha1.Asm{}
	client := reqCtx.K8sClient
	wantedStatus := statusCtx.GetOperatorStatus()
	return retry.RetryOnConflict(retry.DefaultBackoff, func() error {
		err := client.Get(context.Background(), types.NamespacedName{
			Name:      reqCtx.Instance.GetName(),
			Namespace: reqCtx.Instance.GetNamespace(),
		}, asm)
		if err != nil {
			return err
		}
		current := asm.DeepCopy()
		if !equality.OperatorStatusSemantic.DeepDerivative(wantedStatus, current.Status.GetOperatorStatus()) {
			current.Status.State = wantedStatus.State
			if wantedStatus.InstallConditions != nil {
				current.Status.InstallConditions = wantedStatus.InstallConditions
			}
			if wantedStatus.DeleteConditions != nil {
				current.Status.DeleteConditions = wantedStatus.DeleteConditions
			}
			if errUpd := client.Status().Update(context.Background(), current); errUpd != nil {
				return errUpd
			}
		}
		return nil
	})
}

/*
var asmOperatorStatusUpdater = func(current runtime.Object,  client client.Client) error {
	//var asm *operatorv1alpha1.Asm
	//var ok bool
	//if asm, ok = obj.(*operatorv1alpha1.Asm); !ok {
	//	return pkgerrors.New("asmOperatorStatusUpdate cast model.Object to operatorv1alpha1.Asm error")
	//}

	c := types.NamespacedName{}
	_ = client.Get(context.Background(), c, obj)

	return retry.RetryOnConflict(retry.DefaultBackoff, func() error {
		copy := obj.DeepCopyObject()
		if !equality.OperatorStatusSemantic.DeepDerivative(status, copy.Status.GetOperatorStatus()) {
			copy.Status.State = status.State
			if status.InstallConditions != nil {
				copy.Status.InstallConditions = status.InstallConditions
			}
			if status.DeleteConditions != nil {
				copy.Status.DeleteConditions = status.DeleteConditions
			}
			if errUpd := client.Status().Update(context.Background(), copy); errUpd != nil {
				return errUpd
			}
		}
		return nil
	})
	//
	//asmCopy := asm.DeepCopy()
	//if !equality.OperatorStatusSemantic.DeepDerivative(status, asmCopy.Status.GetOperatorStatus()) {
	//	asmCopy.Status.State = status.State
	//	if status.InstallConditions != nil {
	//		asmCopy.Status.InstallConditions = status.InstallConditions
	//	}
	//	if status.DeleteConditions != nil {
	//		asmCopy.Status.DeleteConditions = status.DeleteConditions
	//	}
	//	if errUpd := client.Status().Update(context.Background(), asmCopy); errUpd != nil {
	//		return errUpd
	//	}
	//}
	//return nil
}
*/

func (r *AsmReconciler) SetupWithManager(mgr ctrl.Manager) error {
	return ctrl.NewControllerManagedBy(mgr).
		For(&operatorv1alpha1.Asm{}).
		Complete(r)
}
