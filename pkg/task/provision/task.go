package provision

import (
	"github.com/fyuan1316/asm-operator/pkg/oprlib/manage/model"
	"github.com/fyuan1316/asm-operator/pkg/task/provision/tasks"
)

func GetStages() [][]model.ExecuteItem {
	return [][]model.ExecuteItem{
		{
			tasks.ProvisionCrds,
		},
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
