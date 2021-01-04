package entry

import (
	"github.com/fyuan1316/asm-operator/pkg/oprlib/manage/model"
	"github.com/fyuan1316/asm-operator/pkg/task/migration"
	migrationtasks "github.com/fyuan1316/asm-operator/pkg/task/migration/tasks"
	"github.com/fyuan1316/asm-operator/pkg/task/provision"
	provisiontasks "github.com/fyuan1316/asm-operator/pkg/task/provision/tasks"
)

func GetOperatorStages() ([][]model.ExecuteItem, [][]model.ExecuteItem) {
	return getDeployStages(), getDeleteStages()
}

func getDeployStages() [][]model.ExecuteItem {
	//stages contains migration and deploy
	stages := append(migration.GetStages(), provision.GetStages()...)

	// just test 测试migration
	//stages := migration.GetStages()

	return stages
}

func getDeleteStages() [][]model.ExecuteItem {
	return provision.GetDeleteStages()
}

// 初始化任务数据
func SetUp() error {

	// for deploy stages
	// 1 migrations
	migrationtasks.SetUpMigShell()
	// 2 provisions
	provisiontasks.SetUpCrds()
	provisiontasks.SetUpResource()
	//-----------------
	//for delete stages
	provisiontasks.SetUpDeletion()
	return nil
}
