package migration

import (
	"github.com/fyuan1316/asm-operator/pkg/task/migration/tasks"
	"github.com/fyuan1316/operatorlib/manage/model"
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
