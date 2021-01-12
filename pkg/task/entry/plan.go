package entry

import (
	"github.com/fyuan1316/asm-operator/pkg/task/migration"
	"github.com/fyuan1316/asm-operator/pkg/task/provision"
	"github.com/fyuan1316/operatorlib/manage/model"
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
