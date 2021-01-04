package test

import "time"

type Executable interface {
	Start()
	Execute()
}
type Task struct {
	executor Executable // 实现hook函数的效果：由子类负责编写业务代码
}

func (t *Task) Start() {
	println("ChartTask.Start()")
	// 复用父类代码
	ticker := time.NewTicker(5 * time.Second)
	for range ticker.C {
		t.executor.Execute() // 实现hook函数的效果：由子类负责编写业务代码
	}
}
func (t *Task) Execute() {
	println("ChartTask.Execute()")
}

type CleanTask struct {
	Task
}

func (ct *CleanTask) Execute() {
	println("CleanTask.Execute()")
}
