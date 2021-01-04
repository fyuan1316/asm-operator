package tasks

import (
	"github.com/fyuan1316/asm-operator/pkg/oprlib/manage/model"
	"github.com/fyuan1316/asm-operator/pkg/oprlib/processor/shell"
	resource2 "github.com/fyuan1316/asm-operator/pkg/oprlib/resource"
	"github.com/fyuan1316/asm-operator/pkg/task"
)

var _ model.ExecuteItem = ShellTask{}

type ShellTask struct {
	shell.ScriptManager
}

var ShellTasks ShellTask

func (t ShellTask) Name() string {
	return task.StageMigration
}

//var MigrationShellDir = "pkg/task/migration/shell"
var MigrationShellDir = "files/migration/shell"

func SetUpMigShell() {
	ShellTasks = ShellTask{}
	files, err := resource2.GetFilesInFolder(MigrationShellDir, resource2.Suffix(".sh"))
	if err != nil {
		panic(err)
	}
	if err = ShellTasks.LoadFiles(files); err != nil {
		panic(err)
	}
	//for _, file := range files {
	//	err := ShellTasks.Load(file)
	//	if err != nil {
	//		panic(err)
	//	}
	//}
}
