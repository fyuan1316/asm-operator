package tasks

import (
	"fmt"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/manage/model"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/processor/resource"
	resource2 "github.com/fyuan1316/asm-operator/pkg/oprlib/resource"
	"github.com/fyuan1316/asm-operator/pkg/task"
	"github.com/fyuan1316/asm-operator/pkg/task/data"
)

type ProvisionResourcesTask struct {
	resource.Task
}

var ProvisionResources ProvisionResourcesTask
var _ model.Operation = ProvisionResourcesTask{}

//var _ model.ExecuteItem = ProvisionResourcesTask{}

// 子类需要实现的接口，可以统一合并为一个大的接口定义。
func (p ProvisionResourcesTask) GetOperation() model.OperationType {
	return model.Operations.Provision
}

func (p ProvisionResourcesTask) GetStageName() string {
	return task.StageProvision
}

//func (p ProvisionResourcesTask) Run(om *model.OperatorManage) error {
//	fmt.Println("ProvisionCrdsTask Run")
//	err := p.Sync(om)
//	return err
//}

var ClusterAsmResDir = "pkg/task/provision/cluster-asm/resources"

func SetUpResource() {
	ProvisionResources = ProvisionResourcesTask{
		resource.Task{
			TemplateValues: data.GetDefaults(),
			// 增加自定义的mapping操作
			//ResourceMappings:
		},
	}
	ProvisionResources.Implementor = ProvisionResources

	files, err := resource2.GetFilesInFolder(ClusterAsmResDir, resource2.Suffix(".yaml"))
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
