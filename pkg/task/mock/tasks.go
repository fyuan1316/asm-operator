package mock

import (
	"context"
	"fmt"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/manage/model"
	tasks2 "github.com/fyuan1316/asm-operator/pkg/task/migration/tasks"
	corev1 "k8s.io/api/core/v1"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"
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

func (m PatchTask) GetStageName() string {
	panic("implement me")
}

func (m PatchTask) LiveNess() bool {
	panic("implement me")
}

func (m PatchTask) Run(manage *model.OperatorManage) error {
	fmt.Println("PatchTask Run")
	client := manage.K8sClient
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

func (m PatchTask) PreRun(client client.Client) error {
	fmt.Println("PatchTask prerun")
	return nil
}

func (m PatchTask) PostRun(client client.Client) error {
	fmt.Println("PatchTask PostRun")
	return nil
}

func (m PatchTask) PreCheck(client client.Client) (bool, error) {
	fmt.Println("PatchTask PreCheck")
	return true, nil
}

func (m PatchTask) PostCheck(client client.Client) (bool, error) {
	fmt.Println("PatchTask PostCheck")
	return true, nil
}
