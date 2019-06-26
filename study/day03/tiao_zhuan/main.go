package main

import (
	"fmt"
)

func testGoto() {

	for i := 0; i < 10; i++ {
		for j := 0; j < 20; j++ {
			if j == 14 {
				fmt.Println("开始跳转")
				goto Here
			}
		}
	}

	fmt.Println("----------测试跳转-------------")

	// 手动返回，避免代码正常执行进入此标签
	return
Here:
	fmt.Println("跳我这干嘛")
}

func firstCheckError() []string {
	var arr []string
	return arr
}
func secondCheckError() []string {
	var slices = []string{"1"}
	return slices
}

func errorGoto() {
	err := firstCheckError()
	err2 := secondCheckError()
	if err != nil {
		fmt.Println("出错了1，退出吧！！！")
		goto Exit
	}

	if err2 != nil {
		fmt.Println("出错了2，退出吧！！！")
		goto Exit
	}
	return
Exit:
	fmt.Println("遇到错误，退出", err)
}

func main() {
	// goto 语句通过标进行代码间无条件的跳转，可以快速跳出循环，避免重复退出有一定的帮助。

	// testGoto()

	// 统一错误处理的跳转例子
	errorGoto()

}
