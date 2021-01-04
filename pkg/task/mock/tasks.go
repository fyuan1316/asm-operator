package mock

import (
	"context"
	"fmt"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/manage/model"
	tasks2 "github.com/fyuan1316/asm-operator/pkg/task/migration/tasks"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
)

func GetDeployStages() [][]model.ExecuteItem {
	tasks := [][]model.ExecuteItem{
		{
			tasks2.CustomCrdTask,
			PatchTask{},
		},
		{
			deployTask,
		},
	}

	return tasks
}

type PatchTask struct {
}

func (m PatchTask) Name() string {
	panic("implement me")
}

func (m PatchTask) LiveNess() bool {
	panic("implement me")
}

func (m PatchTask) Run(oCtx *model.OperatorContext) error {
	fmt.Println("PatchTask Run")
	client := oCtx.K8sClient
	ns := corev1.Namespace{}
	err := client.Get(context.Background(), types.NamespacedName{Name: "default"}, &ns)
	if err != nil {
		return err
	}
	if len(ns.Labels) == 0 {
		ns.Labels = make(map[string]string)
	}
	ns.Labels["asm-opr-patch"] = "test-fy"
	if err := client.Update(context.Background(), &ns); err != nil {
		return err
	}
	return nil
}

var _ model.ExecuteItem = PatchTask{}

func (m PatchTask) PreRun(oCtx *model.OperatorContext) error {
	fmt.Println("PatchTask prerun")
	return nil
}

func (m PatchTask) PostRun(oCtx *model.OperatorContext) error {
	fmt.Println("PatchTask PostRun")
	return nil
}

func (m PatchTask) PreCheck(oCtx *model.OperatorContext) (bool, error) {
	fmt.Println("PatchTask PreCheck")
	return true, nil
}

func (m PatchTask) PostCheck(oCtx *model.OperatorContext) (bool, error) {
	fmt.Println("PatchTask PostCheck")
	return true, nil
}
