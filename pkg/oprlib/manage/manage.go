package manage

import (
	"context"
	"github.com/fyuan1316/asm-operator/pkg/logging"
	"k8s.io/apimachinery/pkg/runtime"
	"sigs.k8s.io/controller-runtime/pkg/client"
	"time"
)

var (
	logger = logging.RegisterScope("controller.oprlib")
)

type OperatorManage struct {
	K8sClient client.Client
	Object    Object
	Scheme    *runtime.Scheme
}

func NewOperatorManage(client client.Client, object Object, scheme *runtime.Scheme) *OperatorManage {
	return &OperatorManage{
		K8sClient: client,
		Object:    object,
		Scheme:    scheme,
	}
}

func (m *OperatorManage) Reconcile(stages [][]ExecuteItem) error {
	//delete
	//if !m.Object.GetDeletionTimestamp().IsZero(){
	//
	//}
	//sync
	if err := m.ProcessStages(stages); err != nil {
		return err
	}
	if err := m.HealthCheck(stages); err != nil {
		return err
	}
	return nil

}

func (m *OperatorManage) HealthCheck(stages [][]ExecuteItem) error {
	var (
		total, success int
	)
	for _, items := range stages {
		for _, item := range items {
			if ref, ok := CanDoHealthCheck(item); ok {
				logger.Debugf("run HealthCheck")
				total += 1
				if ref.LiveNess(m.K8sClient) {
					success += 1
				}
			}
		}
	}
	if success == total {

		//aa:= m.Object(v1alpha1.AsmSpec)
		m.K8sClient.Status().Update(context.Background(), m.Object.DeepCopyObject())
	}
	return nil
}

func (m *OperatorManage) ProcessStages(stages [][]ExecuteItem) error {
	for _, items := range stages {
		for _, item := range items {
			if ref, ok := CanDoPreCheck(item); ok {
				logger.Debugf("run precheck")
				if err := loopUntil(context.Background(), 5*time.Second, 10, ref.PreCheck, m.K8sClient); err != nil {
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
			logger.Debugf("execute run")
			if err := item.Run(m); err != nil {
				return err
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
				if err := loopUntil(context.Background(), 5*time.Second, 10, ref.PostCheck, m.K8sClient); err != nil {
					return err
				}
			}
		}
	}
	return nil
}
