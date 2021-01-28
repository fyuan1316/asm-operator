package provision

import (
	"github.com/fyuan1316/operatorlib/manage/model"
	"gitlab-ce.alauda.cn/micro-service/asm-operator/pkg/task/provision/tasks"
)

func GetStages() [][]model.ExecuteItem {
	return [][]model.ExecuteItem{
		{
			tasks.ProvisionResources,
		},
	}
}
func GetDeleteStages() [][]model.ExecuteItem {
	return [][]model.ExecuteItem{
		{
			tasks.DeleteResources,
		},
	}
}
