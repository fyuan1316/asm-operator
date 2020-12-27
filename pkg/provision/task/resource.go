package task

import (
	"fmt"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/manage"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/resource"
	"github.com/fyuan1316/asm-operator/pkg/task/data"
	"os"
	"path/filepath"
)

type ProvisionResourcesTask struct {
	resource.SyncManager
}

func (p ProvisionResourcesTask) Run(om *manage.OperatorManage) error {
	fmt.Println("ProvisionCrdsTask Run")
	err := p.Sync(om)
	return err
}

var ProvisionResources ProvisionResourcesTask

var ClusterAsmResDir = "pkg/provision/cluster-asm/resources"

func init() {
	ProvisionResources = ProvisionResourcesTask{
		resource.SyncManager{SetOwnerReference: resource.PointerTrue()},
	}
	//files, err := ioutil.ReadDir(ClusterAsmResDir)
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
		//filepath := ClusterAsmResDir + "/" + file
		err := ProvisionResources.LoadFile(
			file,
			&resource.SyncResource{
				//Object: &apiextensionsv1.CustomResourceDefinition{},
				//Sync: sync.FnCrd,
			},
			data.GetDefaults(),
		)
		if err != nil {
			panic(err)
		}
	}
}

var _ manage.ExecuteItem = ProvisionCrdsTask{}
