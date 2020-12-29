package provision

import (
	"github.com/fyuan1316/asm-operator/pkg/oprlib/manage/model"
	"github.com/fyuan1316/asm-operator/pkg/task/provision/tasks"
)

func GetStages() [][]model.ExecuteItem {
	tasks := [][]model.ExecuteItem{
		{
			tasks.ProvisionCrds,
		},
		{
			tasks.ProvisionResources,
		},
	}

	return tasks
}
func GetDeleteStages() [][]model.ExecuteItem {
	tasks := [][]model.ExecuteItem{
		{
			tasks.DeleteResources,
		},
	}

	return tasks
}
