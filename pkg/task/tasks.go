package task

import (
	"github.com/fyuan1316/asm-operator/pkg/oprlib/manage"
	"github.com/fyuan1316/asm-operator/pkg/provision"
)

func GetDeployStages() [][]manage.ExecuteItem {
	//return mock.GetDeployStages()
	return provision.GetStages()
}

func GetDeleteStages() [][]manage.ExecuteItem {
	//return mock.GetDeployStages()
	return provision.GetDeleteStages()
}
