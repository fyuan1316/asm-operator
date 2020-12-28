package provision

import (
	"github.com/fyuan1316/asm-operator/pkg/oprlib/manage/model"
	"github.com/fyuan1316/asm-operator/pkg/task/provision/task"
)

func GetStages() [][]model.ExecuteItem {
	tasks := [][]model.ExecuteItem{
		{
			task.ProvisionCrds,
		},
		{
			task.ProvisionResources,
		},
	}

	return tasks
}
func GetDeleteStages() [][]model.ExecuteItem {
	tasks := [][]model.ExecuteItem{
		{
			task.DeleteResources,
		},
	}

	return tasks
}
