package task

import (
	"fmt"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/manage"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/resource"
	"github.com/fyuan1316/asm-operator/pkg/task/data"
	"os"
	"path/filepath"
)

type DeleteResourcesTask struct {
	resource.SyncManager
}

func (p DeleteResourcesTask) Run(om *manage.OperatorManage) error {
	fmt.Println("DeleteResourcesTask Run")
	err := p.Delete(om)
	return err
}

var DeleteResources DeleteResourcesTask

//var ClusterAsmResDir = "pkg/provision/cluster-asm/resources"

func init() {
	DeleteResources = DeleteResourcesTask{
		resource.SyncManager{SetOwnerReference: resource.PointerTrue()},
	}
	var files []string
	err := filepath.Walk(ClusterAsmResDir, func(path string, info os.FileInfo, err error) error {
		if filepath.Ext(path) != ".yaml" {
			return nil
		}
		files = append(files, path)
		return nil
	})
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

var _ manage.ExecuteItem = DeleteResourcesTask{}
