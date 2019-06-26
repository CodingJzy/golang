package main

import (
	"fmt"
)

func testIf1() {
	num := 10
	if num < 9 {
		fmt.Println("小于")
	} else if num == 10 {
		fmt.Println("等于")
	} else {
		fmt.Println("大于")
	}
}

func connect() []string {
	var slices []string
	slices = append(slices, "Go")
	return slices[:]
}

func testIf2() {
	if err := connect(); err != nil {
		fmt.Println("切片不为空", err)
		return
	}
}

func main() {
	fmt.Println("条件判断if")

	/*
		格式：
			if 表达式1 ｛
				分支1
			｝ else if 表达式2 ｛
				分支2
			｝ else {
				分支3
			}
	*/

	// 常用写法
	// testIf1()

	// 特殊写法：可以在if表达式之前添加一个执行语句，再根据变量值进行判断
	testIf2()
}
