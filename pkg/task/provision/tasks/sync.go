package tasks

import (
	"github.com/fyuan1316/asm-operator/pkg/task"
	"github.com/fyuan1316/operatorlib/manage/model"
	"github.com/fyuan1316/operatorlib/task/chart"
)

type ProvisionResourcesTask struct {
	*chart.ChartTask
}

var ProvisionResources ProvisionResourcesTask
var _ model.OverrideOperation = ProvisionResourcesTask{}
var _ model.HealthCheck = ProvisionResourcesTask{}

func (p ProvisionResourcesTask) IsReady(oCtx *model.OperatorContext) bool {
	return true
}

func (p ProvisionResourcesTask) IsHealthy(oCtx *model.OperatorContext) bool {
	return false
}

func (p ProvisionResourcesTask) GetOperation() model.OperationType {
	return model.Operations.Provision
}

func (p ProvisionResourcesTask) Name() string {
	return task.StageProvision + "-task"
}

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
