package migration

import (
	"github.com/fyuan1316/operatorlib/manage/model"
	"gitlab-ce.alauda.cn/micro-service/asm-operator/pkg/task/migration/tasks"
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
