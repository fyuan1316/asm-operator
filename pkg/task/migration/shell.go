package migration

import (
	"fmt"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/manage/model"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/processor/shell"
	resource2 "github.com/fyuan1316/asm-operator/pkg/oprlib/resource"
	"github.com/fyuan1316/asm-operator/pkg/task"
	"sigs.k8s.io/controller-runtime/pkg/client"
)

type ShellTask struct {
	shell.ScriptManager
}

func (s ShellTask) GetStageName() string {
	return task.StageMigration
}

var shellTask ShellTask
var MigrationShellDir = "pkg/task/migration/shell"

func SetUpMigShell() {
	shellTask = ShellTask{}
	files, err := resource2.GetFilesInFolder(MigrationShellDir, resource2.Suffix(".sh"))
	if err != nil {
		panic(err)
	}
	for _, file := range files {
		err := shellTask.Load(file)
		if err != nil {
			panic(err)
		}
	}
}

func (m ShellTask) PreRun(client client.Client) error {
	fmt.Println("ShellTask prerun")

	return nil
}
func (s ShellTask) Run(manage *model.OperatorManage) error {
	return nil
}

var _ model.ExecuteItem = ShellTask{}
