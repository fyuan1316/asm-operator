package tasks

import (
	"github.com/fyuan1316/asm-operator/pkg/task"
	"github.com/fyuan1316/operatorlib/manage/model"
	"github.com/fyuan1316/operatorlib/task/shell"
)

var _ model.ExecuteItem = ShellTask{}

type ShellTask struct {
	shell.ScriptManager
}

var ShellTasks ShellTask

func (t ShellTask) Name() string {
	return task.StageMigration + "-task"
}

var MigrationShellDir = "files/migration/shell"

func SetUpMigShell() {
	ShellTasks = ShellTask{}
	if err := ShellTasks.LoadDir(MigrationShellDir); err != nil {
		panic(err)
	}
}
