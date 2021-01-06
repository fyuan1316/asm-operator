package test

import (
	"encoding/json"
	"fmt"
	"testing"
)

func Test_t1(t *testing.T) {
	/*
		cleanTask := &CleanTask{
			Task{},
		}
		cleanTask.executor = cleanTask // 实现hook函数的效果：由子类负责编写业务代码
		cleanTask.Start()
	*/
}

func Test_useDefaultProcess(t *testing.T) {
	/*
		cleanTask := &ChartTask{}
		cleanTask.executor = cleanTask
		cleanTask.Start()

	*/
}

func Test_convert(t *testing.T) {
	m := map[string]string{"k1": "v1", "k2": "v2"}
	var err error
	var mi map[string]interface{}
	var bytes []byte
	bytes, err = json.Marshal(m)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal([]byte(bytes), &mi)
	if err != nil {
		panic(err)
	}

	fmt.Println(mi)
}
