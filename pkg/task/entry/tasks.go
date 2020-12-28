package entry

import (
	"github.com/fyuan1316/asm-operator/pkg/oprlib/manage/model"
	"github.com/fyuan1316/asm-operator/pkg/task/provision"
	"github.com/fyuan1316/asm-operator/pkg/task/provision/task"
)

func GetAllProvisionStages() [][]model.ExecuteItem {
	//return mock.GetDeployStages()
	return provision.GetStages()
}

func GetDeployStages() [][]model.ExecuteItem {
	//return mock.GetDeployStages()
	return provision.GetStages()
}

func GetDeleteStages() [][]model.ExecuteItem {
	//return mock.GetDeployStages()
	return provision.GetDeleteStages()
}

func SetUp() error {
	task.SetUpCrds()
	task.SetUpResource()
	task.SetUpDeletion()
	return nil
}
