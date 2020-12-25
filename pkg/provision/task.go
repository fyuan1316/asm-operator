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
		//{
		//	provisionResources,
		//},
	}

	return tasks
}
