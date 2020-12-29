package t2

import "time"

type Executable interface {
	Start()
	Execute()
}
type Task struct {
	//executor Executable // 实现hook函数的效果：由子类负责编写业务代码
}

func (t *Task) Start() {
	println("Task.Start()")
	// 复用父类代码
	ticker := time.NewTicker(5 * time.Second)
	for range ticker.C {
		t.Execute() // 实现hook函数的效果：由子类负责编写业务代码
	}
}
func (t *Task) Execute() {
	println("Task.Execute()")
}

type CleanTask struct {
	Task
}

func (ct *CleanTask) Execute() {
	println("CleanTask.Execute()")
}
func (ct *CleanTask) Start() {
	println("CleanTask.Start()")
}
