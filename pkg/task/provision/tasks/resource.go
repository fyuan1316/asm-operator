package tasks

import (
	"fmt"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/manage/model"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/processor/resource"
	resource2 "github.com/fyuan1316/asm-operator/pkg/oprlib/resource"
	"github.com/fyuan1316/asm-operator/pkg/task"
	"github.com/fyuan1316/asm-operator/pkg/task/data"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ProvisionResourcesTask struct {
	*resource.Task
}

func (p ProvisionResourcesTask) IsReady(client client.Client) bool {
	return true
}

func (p ProvisionResourcesTask) IsHealthy(client client.Client) bool {
	return false
}

var ProvisionResources ProvisionResourcesTask
var _ model.OverrideOperation = ProvisionResourcesTask{}
var _ model.HealthCheck = ProvisionResourcesTask{}

// 子类需要实现的接口，可以统一合并为一个大的接口定义。
func (p ProvisionResourcesTask) GetOperation() model.OperationType {
	return model.Operations.Provision
}

func (p ProvisionResourcesTask) GetStageName() string {
	return task.StageProvision
}

var ClusterAsmResDir = "pkg/task/provision/cluster-asm/resources"

func SetUpResource() {
	ProvisionResources = ProvisionResourcesTask{
		&resource.Task{
			TemplateValues: data.GetDefaults(),
			// 增加自定义的mapping操作
			//ResourceMappings:
		},
	}
	ProvisionResources.Override(ProvisionResources)
	//files, err := resource2.GetFilesInFolder(ClusterAsmResDir, resource2.Suffix(".yaml"))
	test := "/Users/yuan/Dev/GolangProjects/charts/chart-cluster-asm-copy/chart"
	files, err := resource2.GetChartResources(test)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		err := ProvisionResources.LoadFile(
			file,
			// 当需要保留此资源时设置，优先级高于task 设置
			//resource.KeepResourceAfterOperatorDeleted(),
		)
		if err != nil {
			panic(err)
		}
	}
	fmt.Println()
}
