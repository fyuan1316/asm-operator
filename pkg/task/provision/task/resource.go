package task

import (
	"fmt"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/manage/model"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/resource"
	"github.com/fyuan1316/asm-operator/pkg/task"
	"github.com/fyuan1316/asm-operator/pkg/task/data"
)

type ProvisionResourcesTask struct {
	resource.SyncManager
}

func (p ProvisionResourcesTask) GetStageName() string {
	return task.StageProvision
}

func (p ProvisionResourcesTask) Run(om *model.OperatorManage) error {
	fmt.Println("ProvisionCrdsTask Run")
	err := p.Sync(om)
	return err
}

var ProvisionResources ProvisionResourcesTask

var ClusterAsmResDir = "pkg/provision/cluster-asm/resources"

func SetUpResource() {
	ProvisionResources = ProvisionResourcesTask{
		resource.SyncManager{ChargeByOperator: resource.PointerTrue()},
	}
	files, err := resource.GetFilesInFolder(ClusterAsmResDir, resource.Suffix(".yaml"))
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		err := ProvisionResources.LoadFile(
			file,
			&resource.SyncResource{},
			data.GetDefaults(),
		)
		if err != nil {
			panic(err)
		}
	}
}

var _ model.ExecuteItem = ProvisionResourcesTask{}
