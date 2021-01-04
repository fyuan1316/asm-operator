package t2

import (
	"testing"
)

func Test_t1(t *testing.T) {
	cleanTask := &CleanTask{}
	//cleanTask.executor = cleanTask // 实现hook函数的效果：由子类负责编写业务代码
	cleanTask.Start()

}

func Test_useDefaultProcess(t *testing.T) {
	/*
		cleanTask := &ChartTask{}
		//cleanTask.executor = cleanTask
		cleanTask.Start()

	*/
}
