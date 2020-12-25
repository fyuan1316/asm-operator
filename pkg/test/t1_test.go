package test

import (
	"fmt"
	"testing"
)

func Test_t1(t *testing.T){
	use1 := TestStruct{}

	//if h,ok := use1.(Human);ok{
	//	fmt.Println()
	//}else{
	//	fmt.Println("not a human")
	//}
	we_test(use1)

}
func we_test(i interface{} ) {
	if h,ok := i.(Human);ok{
		fmt.Println(h.HasName())
	}else{
		fmt.Println("not a human")
	}
}