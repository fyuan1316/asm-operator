package task

import (
	"context"
	"errors"
	"fmt"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/manage/model"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/resource"
	"github.com/fyuan1316/asm-operator/pkg/task"
	"github.com/fyuan1316/asm-operator/pkg/task/data"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ProvisionCrdsTask struct {
	resource.SyncManager
}

func (p ProvisionCrdsTask) GetStageName() string {
	return task.StageProvision
}

func (p ProvisionCrdsTask) PreCheck(client client.Client) (pass bool, err error) {
	var isExistsFn = func(name string) bool {
		crd := apiextensionsv1.CustomResourceDefinition{}
		if getErr := client.Get(context.Background(),
			types.NamespacedName{Name: name},
			&crd,
		); getErr != nil && apierrors.IsNotFound(getErr) {
			return false
		}
		return true
	}
	crdList := []string{
		"podmonitors.monitoring.coreos.com",
		"servicemonitors.monitoring.coreos.com",
	}
	var errs []error
	for _, crd := range crdList {
		if !isExistsFn(crd) {
			pass = false
			errs = append(errs, errors.New("crd "+crd+" not found"))
		}
	}

	return pass, utilerrors.NewAggregate(errs)
}

func (p ProvisionCrdsTask) Run(om *model.OperatorManage) error {
	fmt.Println("ProvisionCrdsTask Run")
	err := p.Sync(om)
	return err
}

var ProvisionCrds ProvisionCrdsTask

var ClusterAsmCrdDir = "pkg/provision/cluster-asm/crds"

func SetUpCrds() {
	ProvisionCrds = ProvisionCrdsTask{}

	files, err := resource.GetFilesInFolder(ClusterAsmCrdDir, resource.Suffix(".yaml"))
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		err := ProvisionCrds.LoadFile(
			file,
			&resource.SyncResource{},
			data.GetDefaults(),
		)
		if err != nil {
			panic(err)
		}
	}
}

var _ model.ExecuteItem = ProvisionCrdsTask{}
var _ model.PreCheck = ProvisionCrdsTask{}
