package tasks

import (
	"context"
	"github.com/fyuan1316/operatorlib/manage/model"
	"github.com/fyuan1316/operatorlib/task/chart"
	pkgerrors "github.com/pkg/errors"
	corev1 "k8s.io/api/core/v1"
	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ProvisionResourcesTask struct {
	*chart.ChartTask
}

func (p ProvisionResourcesTask) GetName() string {
	return "asm-controlplane-install"
}

var ProvisionResources ProvisionResourcesTask
var _ model.OverrideOperation = ProvisionResourcesTask{}
var _ model.HealthCheck = ProvisionResourcesTask{}
var _ model.PreCheck = ProvisionResourcesTask{}

func (p ProvisionResourcesTask) PreCheck(ctx *model.OperatorContext) (bool, error) {
	key := client.ObjectKey{Name: ctx.InstalledNamespace}
	current := corev1.Namespace{}
	current.Name = ctx.InstalledNamespace
	err := ctx.K8sClient.Get(context.TODO(), key, &current)
	if err != nil {
		if !apierrors.IsNotFound(err) {
			return false, pkgerrors.WithStack(err)
		}
		if cErr := ctx.K8sClient.Create(context.TODO(), &current); cErr != nil {
			return false, pkgerrors.WithStack(err)
		}
		return true, nil
	}
	return true, nil
}

func (p ProvisionResourcesTask) IsReady(oCtx *model.OperatorContext) bool {
	return true
}

func (p ProvisionResourcesTask) IsHealthy(oCtx *model.OperatorContext) bool {
	return false
}

func (p ProvisionResourcesTask) GetOperation() model.OperationType {
	return model.Operations.Provision
}

//func (p ProvisionResourcesTask) Name() string {
//	return task.StageProvision + "-task"
//}

var ClusterAsmResDir = "files/provision/cluster-asm"

func SetUpProvision() {
	ProvisionResources = ProvisionResourcesTask{
		&chart.ChartTask{
			Dir: ClusterAsmResDir,
		},
	}
	ProvisionResources.Init()
	ProvisionResources.Override(ProvisionResources)
}
