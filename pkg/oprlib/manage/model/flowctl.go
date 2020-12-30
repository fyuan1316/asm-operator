package model

import (
	"context"
	"github.com/fyuan1316/asm-operator/pkg/logging"
	pkgerrors "github.com/pkg/errors"
	ctrl "sigs.k8s.io/controller-runtime"
	"sigs.k8s.io/controller-runtime/pkg/reconcile"
	"time"
)

var (
	logger = logging.RegisterScope("controller.oprlib")
)

func (m *OperatorManage) Reconcile(provisionStages, deletionStages [][]ExecuteItem) (ctrl.Result, error) {
	if m.Options.FinalizerID != "" && m.CR.GetDeletionTimestamp().IsZero() {
		if !ContainsString(m.CR.GetFinalizers(), m.Options.FinalizerID) {
			finalizers := append(m.CR.GetFinalizers(), m.Options.FinalizerID)
			m.CR.SetFinalizers(finalizers)
			if err := m.K8sClient.Update(context.Background(), m.CR); err != nil {
				return ctrl.Result{}, pkgerrors.Wrap(err, "could not add finalizer to config")
			}
			return ctrl.Result{
				RequeueAfter: time.Second * 1,
			}, nil
		}
	} else if !m.CR.GetDeletionTimestamp().IsZero() {
		if len(deletionStages) > 0 {
			if err := m.ProcessStages(deletionStages); err != nil {
				return ctrl.Result{}, err
			}
		}
		f := RemoveString(m.CR.GetFinalizers(), m.Options.FinalizerID)
		m.CR.SetFinalizers(f)
		if err := m.K8sClient.Update(context.Background(), m.CR); err != nil {
			return reconcile.Result{}, pkgerrors.Wrap(err, "could not remove finalizer from config")
		}
		return ctrl.Result{}, nil
	}
	if len(provisionStages) == 0 {
		return ctrl.Result{}, nil
	}
	//sync
	if err := m.ProcessStages(provisionStages); err != nil {
		return ctrl.Result{}, err
	}
	if err := m.DoHealthCheck(provisionStages); err != nil {
		return ctrl.Result{}, err
	}
	return ctrl.Result{}, nil

}

func (m *OperatorManage) DoHealthCheck(stages [][]ExecuteItem) error {
	var readyCheckNum, readyNum, healthyCheckNum, healthyNum int
	for _, items := range stages {
		for _, item := range items {
			if ref, ok := CanDoHealthCheck(item); ok {
				logger.Debugf("run HealthCheck")
				readyCheckNum += 1
				if ref.IsReady(m.K8sClient) {
					readyNum += 1
				}
				healthyCheckNum += 1
				if ref.IsHealthy(m.K8sClient) {
					healthyNum += 1
				}
			}
		}
	}
	// if some task needs report its states, we update operator cr's status.state
	if readyCheckNum > 0 {
		if err := m.Options.StatusUpdater(m.CR, m.K8sClient)(readyCheckNum == readyNum, healthyCheckNum == healthyNum); err != nil {
			return err
		}
		return nil
	}
	return nil
}

func (m *OperatorManage) ProcessStages(stages [][]ExecuteItem) error {
	for _, items := range stages {
		for _, item := range items {
			if ref, ok := CanDoPreCheck(item); ok {
				logger.Debugf("run precheck")
				if err := loopUntil(context.Background(), 5*time.Second, 3, ref.PreCheck, m.K8sClient); err != nil {
					return err
				}
			}
		}
		for _, item := range items {
			if ref, ok := CanDoPreRun(item); ok {
				logger.Debugf("run prerun")
				if err := ref.PreRun(m.K8sClient); err != nil {
					return err
				}
			}
		}
		for _, item := range items {
			if ref, ok := CanDoRun(item); ok {
				logger.Debugf("execute run")
				if err := ref.Run(m); err != nil {
					return err
				}
			}
		}

		for _, item := range items {
			if ref, ok := CanDoPostRun(item); ok {
				logger.Debugf("run postrun")
				if err := ref.PostRun(m.K8sClient); err != nil {
					return err
				}
			}
		}
		for _, item := range items {
			if ref, ok := CanDoPostCheck(item); ok {
				logger.Debugf("run postcheck")
				if err := loopUntil(context.Background(), 5*time.Second, 3, ref.PostCheck, m.K8sClient); err != nil {
					return err
				}
			}
		}
	}
	return nil
}

func ContainsString(slice []string, s string) bool {
	for _, item := range slice {
		if item == s {
			return true
		}
	}
	return false
}
func RemoveString(slice []string, s string) (result []string) {
	for _, item := range slice {
		if item == s {
			continue
		}
		result = append(result, item)
	}
	return
}
