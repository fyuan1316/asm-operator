package entry

import (
	migrationtasks "github.com/fyuan1316/asm-operator/pkg/task/migration/tasks"
	provisiontasks "github.com/fyuan1316/asm-operator/pkg/task/provision/tasks"
)

// 初始化任务数据
func SetUp() error {

	// for deploy stages
	// 1 migrations
	migrationtasks.SetUpMigShell()

	// 2 provisions
	//provisiontasks.SetUpCrds()
	provisiontasks.SetUpProvision()
	//-----------------
	//for delete stages
	provisiontasks.SetUpDeletion()
	return nil
}
