package provision

import (
	"github.com/fyuan1316/asm-operator/pkg/task/provision/tasks"
	"github.com/fyuan1316/operatorlib/manage/model"
)

func GetStages() [][]model.ExecuteItem {
	return [][]model.ExecuteItem{
		//{
		//	tasks.ProvisionCrds,
		//},
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
