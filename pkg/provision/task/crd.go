package task

import (
	"fmt"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/manage"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/resource"
	"github.com/fyuan1316/asm-operator/pkg/task/data"
	"io/ioutil"
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

var ClusterAsmCrdDir = "pkg/provision/cluster-asm/crds"

func init() {
	ProvisionCrds = ProvisionCrdsTask{}
	files, err := ioutil.ReadDir(ClusterAsmCrdDir)
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		filepath := ClusterAsmCrdDir + "/" + file.Name()
		err := ProvisionCrds.LoadFile(
			filepath,
			&resource.SyncResource{},
			data.GetDefaults(),
		)
		if err != nil {
			panic(err)
		}
	}
}

var _ manage.ExecuteItem = ProvisionCrdsTask{}
