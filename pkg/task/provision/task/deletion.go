package task

import (
	"fmt"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/manage/model"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/resource"
	"github.com/fyuan1316/asm-operator/pkg/task"
	"github.com/fyuan1316/asm-operator/pkg/task/data"
)

type DeleteResourcesTask struct {
	resource.SyncManager
}

func (p DeleteResourcesTask) GetStageName() string {
	return task.StageDeletion
}

func (p DeleteResourcesTask) Run(om *model.OperatorManage) error {
	fmt.Println("DeleteResourcesTask Run")
	err := p.Delete(om)
	return err
}

var DeleteResources DeleteResourcesTask

func SetUpDeletion() {
	DeleteResources = DeleteResourcesTask{
		resource.SyncManager{ChargeByOperator: resource.PointerTrue()},
	}
	files, err := resource.GetFilesInFolder(ClusterAsmResDir, resource.Suffix(".yaml"))
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		err := DeleteResources.LoadFile(
			file,
			&resource.SyncResource{},
			data.GetDefaults(),
		)
		if err != nil {
			panic(err)
		}
	}
}

var _ model.ExecuteItem = DeleteResourcesTask{}
