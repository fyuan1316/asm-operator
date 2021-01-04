package tasks

import (
	"fmt"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/manage/model"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/processor/resource"
	resource2 "github.com/fyuan1316/asm-operator/pkg/oprlib/resource"
	"github.com/fyuan1316/asm-operator/pkg/task"
)

type ProvisionResourcesTask struct {
	*resource.ChartTask
}

func (p ProvisionResourcesTask) IsReady(oCtx *model.OperatorContext) bool {
	return true
}

func (p ProvisionResourcesTask) IsHealthy(oCtx *model.OperatorContext) bool {
	return false
}

var ProvisionResources ProvisionResourcesTask
var _ model.OverrideOperation = ProvisionResourcesTask{}
var _ model.HealthCheck = ProvisionResourcesTask{}

// 子类需要实现的接口，可以统一合并为一个大的接口定义。
func (p ProvisionResourcesTask) GetOperation() model.OperationType {
	return model.Operations.Provision
}

func (p ProvisionResourcesTask) Name() string {
	return task.StageProvision
}

var ClusterAsmResDir = "pkg/task/provision/cluster-asm/resources"

func SetUpResource() {
	ClusterAsmResDir = "files/provision/cluster-asm"
	ProvisionResources = ProvisionResourcesTask{
		&resource.ChartTask{
			ChartDir: ClusterAsmResDir,
			//TemplateValues: data.GetDefaults(),
			// 增加自定义的mapping操作
			//ResourceMappings:

		},
	}
	ProvisionResources.Override(ProvisionResources)
	//files, err := resource2.GetFilesInFolder(ClusterAsmResDir, resource2.Suffix(".yaml"))
	//test := "/Users/yuan/Dev/GolangProjects/charts/cluster-asm-cluster-asm-copy/cluster-asm"
	//files, err := resource2.GetChartResources(test)
	//if err != nil {
	//	panic(err)
	//}
	//for _, file := range files {
	//	err := ProvisionResources.LoadFile(
	//		file,
	//		// 当需要保留此资源时设置，优先级高于task 设置
	//		//resource.KeepResourceAfterOperatorDeleted(),
	//	)
	//	if err != nil {
	//		panic(err)
	//	}
	//}
	var (
		files map[string]string
		err   error
	)

	if files, err = resource2.GetChartResources(ClusterAsmResDir, nil); err != nil {
		panic(err)
	}
	if err = ProvisionResources.LoadResources(files); err != nil {
		panic(err)
	}
	fmt.Println()
}
