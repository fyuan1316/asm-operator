package task

import (
	"fmt"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/manage"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/resource"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/sync"
	"io/ioutil"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
)

type ProvisionCrdsTask struct {
	resource.SyncManager
}

func (p ProvisionCrdsTask) Run(om *manage.OperatorManage) error {
	fmt.Println("ProvisionCrdsTask Run")
	err := p.Sync(om)
	return err
}

var ProvisionCrds ProvisionCrdsTask

var ClusterAsmDir = "pkg/provision/cluster-asm"

func init() {
	ProvisionCrds = ProvisionCrdsTask{}
	//crdRes := resource.SyncResource{
	//	Object: &apiextensionsv1.CustomResourceDefinition{},
	//	Sync:   sync.CreateIFNotFound,
	//}
	files, err := ioutil.ReadDir(ClusterAsmDir + "/crds")
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		filepath := ClusterAsmDir + "/crds/" + file.Name()
		err := ProvisionCrds.LoadFile(
			filepath,
			resource.SyncResource{
				Object: &apiextensionsv1.CustomResourceDefinition{},
				Sync:   sync.CreateIFNotFound,
			})
		if err != nil {
			panic(err)
		}
	}
}

var _ manage.ExecuteItem = ProvisionCrdsTask{}
