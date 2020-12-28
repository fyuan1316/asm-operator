package tasks

import (
	"github.com/fyuan1316/asm-operator/pkg/oprlib/manage/model"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/processor/resource"
	resource2 "github.com/fyuan1316/asm-operator/pkg/oprlib/resource"
	"github.com/fyuan1316/asm-operator/pkg/task"
	"github.com/fyuan1316/asm-operator/pkg/task/data"
)

type DeleteResourcesTask struct {
	resource.Task
}

var DeleteResources DeleteResourcesTask
var _ model.Operation = DeleteResourcesTask{}

func (p DeleteResourcesTask) GetOperation() model.OperationType {
	return model.Operations.Deletion
}

func (p DeleteResourcesTask) GetStageName() string {
	return task.StageDeletion
}

//func (p DeleteResourcesTask) Run(om *model.OperatorManage) error {
//	fmt.Println("DeleteResourcesTask Run")
//	err := p.Delete(om)
//	return err
//}

func SetUpDeletion() {
	DeleteResources = DeleteResourcesTask{
		resource.Task{
			TemplateValues: data.GetDefaults(),
			// 增加自定义的mapping操作
			//ResourceMappings:
			KeepResourceAfterOperatorDeleted: resource.PointerFalse(),
		},
	}
	DeleteResources.Implementor = DeleteResources

	files, err := resource2.GetFilesInFolder(ClusterAsmResDir, resource2.Suffix(".yaml"))
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		err := DeleteResources.LoadFile(file)
		if err != nil {
			panic(err)
		}
	}
}
