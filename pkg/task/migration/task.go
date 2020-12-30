package migration

import (
	"github.com/fyuan1316/asm-operator/pkg/oprlib/manage/model"
	"github.com/fyuan1316/asm-operator/pkg/task/migration/tasks"
)

func GetStages() [][]model.ExecuteItem {
	return [][]model.ExecuteItem{
		//{
		//	tasks.CustomCrdTask,
		//},
		{
			tasks.ShellTasks,
		},
	}
}
