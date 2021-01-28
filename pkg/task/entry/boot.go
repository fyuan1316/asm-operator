package entry

import (
	migrationtasks "gitlab-ce.alauda.cn/micro-service/asm-operator/pkg/task/migration/tasks"
	provisiontasks "gitlab-ce.alauda.cn/micro-service/asm-operator/pkg/task/provision/tasks"
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
