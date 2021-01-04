package tasks

import (
	"context"
	"fmt"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/manage/model"
	"github.com/fyuan1316/asm-operator/pkg/task"
	"io/ioutil"
	apiextensionsv1 "k8s.io/apiextensions-apiserver/pkg/apis/apiextensions/v1"
	"k8s.io/apimachinery/pkg/api/errors"
	"sigs.k8s.io/yaml"
)

var _ model.ExecuteItem = CustomTask{}

type CustomTask struct {
}

var CustomCrdTask CustomTask

func init() {
	CustomCrdTask = CustomTask{}

}

func (m CustomTask) Name() string {
	return task.StageMigration
}

func (m CustomTask) LiveNess() bool {
	panic("implement me")
}

/**
自定义task实现
*/
func (m CustomTask) Run(oCtx *model.OperatorContext) error {

	fmt.Println("ChangeCrdTask Run")
	shFilePath := "files/migration/crds/1-vs.yaml"
	crdList := &apiextensionsv1.CustomResourceDefinitionList{}
	err := oCtx.K8sClient.List(context.Background(), crdList)
	if err != nil {
		return err
	}
	bytes, err := ioutil.ReadFile(shFilePath)
	if err != nil {
		return err
	}
	crd := apiextensionsv1.CustomResourceDefinition{}
	err = yaml.Unmarshal(bytes, &crd)
	if err != nil {
		return err
	}
	err = oCtx.K8sClient.Create(context.Background(), &crd)
	if errors.IsAlreadyExists(err) {
		err = nil
	}
	return err
}

func (m CustomTask) PreRun(oCtx *model.OperatorContext) error {
	fmt.Println("ChangeCrdTask prerun")
	crdList := &apiextensionsv1.CustomResourceDefinitionList{}
	err := oCtx.K8sClient.List(context.Background(), crdList)
	fmt.Println(len(crdList.Items), err)
	return nil
}

func (m CustomTask) PostRun(oCtx *model.OperatorContext) error {
	fmt.Println("ChangeCrdTask PostRun")
	return nil
}

func (m CustomTask) PreCheck(oCtx *model.OperatorContext) (bool, error) {
	fmt.Println("ChangeCrdTask PreCheck")
	return true, nil
}

func (m CustomTask) PostCheck(oCtx *model.OperatorContext) (bool, error) {
	fmt.Println("ChangeCrdTask PostCheck")
	return true, nil
}
