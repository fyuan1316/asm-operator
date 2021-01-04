package tasks

import (
	"context"
	"errors"
	"fmt"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/manage/model"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/processor/resource"
	resource2 "github.com/fyuan1316/asm-operator/pkg/oprlib/resource"
	"github.com/fyuan1316/asm-operator/pkg/task"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"

	apierrors "k8s.io/apimachinery/pkg/api/errors"
	"k8s.io/apimachinery/pkg/types"
	utilerrors "k8s.io/apimachinery/pkg/util/errors"
)

type ProvisionCrdsTask struct {
	resource.FileTask
}

func (p ProvisionCrdsTask) Name() string {
	return task.StageProvision
}

func (p ProvisionCrdsTask) PreCheck(oCtx *model.OperatorContext) (bool, error) {
	var isExistsFn = func(name string) bool {
		crd := apiextensionsv1.CustomResourceDefinition{}
		if getErr := oCtx.K8sClient.Get(context.Background(),
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
	var pass = true
	for _, crd := range crdList {
		if !isExistsFn(crd) {
			pass = false
			errs = append(errs, errors.New("crd "+crd+" not found"))
		}
	}

	return pass, utilerrors.NewAggregate(errs)
}

func (p ProvisionCrdsTask) Run(ctx *model.OperatorContext) error {
	fmt.Println("ProvisionCrdsTask Run")
	err := p.Sync(ctx)
	return err
}

var ProvisionCrds ProvisionCrdsTask

var ClusterAsmCrdDir = "files/provision/crds"

func SetUpCrds() {
	ProvisionCrds = ProvisionCrdsTask{
		FileTask: resource.FileTask{
			//TemplateValues: data.GetDefaults(),
		},
	}
	files, err := resource2.GetFilesInFolder(ClusterAsmCrdDir, resource2.Suffix(".yaml"))
	if err != nil {
		panic(err)
	}

	for path := range files {
		err := ProvisionCrds.LoadFile(
			path,
		)
		if err != nil {
			panic(err)
		}
	}
}

var _ model.ExecuteItem = ProvisionCrdsTask{}
var _ model.PreCheck = ProvisionCrdsTask{}
