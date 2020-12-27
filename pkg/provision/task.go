package provision

import (
	"github.com/fyuan1316/asm-operator/pkg/oprlib/manage"
	"github.com/fyuan1316/asm-operator/pkg/provision/task"
)

func GetStages() [][]manage.ExecuteItem {
	tasks := [][]manage.ExecuteItem{
		{
			task.ProvisionCrds,
		},
		{
			task.ProvisionResources,
		},
	}

	return tasks
}
func GetDeleteStages() [][]manage.ExecuteItem {
	tasks := [][]manage.ExecuteItem{
		{
			task.DeleteResources,
		},
	}

	return tasks
}


