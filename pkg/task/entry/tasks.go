package entry

import (
	"github.com/fyuan1316/asm-operator/pkg/oprlib/manage/model"
	"github.com/fyuan1316/asm-operator/pkg/task/provision"
	"github.com/fyuan1316/asm-operator/pkg/task/provision/tasks"
)

func GetAllProvisionStages() [][]model.ExecuteItem {
	//return mock.GetDeployStages()
	return provision.GetStages()
}

func GetDeployStages() [][]model.ExecuteItem {
	//return mock.GetDeployStages()
	//append(migration)
	return provision.GetStages()
}

func GetDeleteStages() [][]model.ExecuteItem {
	//return mock.GetDeployStages()
	return provision.GetDeleteStages()
}

func SetUp() error {
	tasks.SetUpCrds()
	tasks.SetUpResource()
	tasks.SetUpDeletion()
	return nil
}
