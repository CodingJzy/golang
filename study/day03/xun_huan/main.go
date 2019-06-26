package main

import (
	"fmt"
)

func testFor1() {
	// 初始化语句是在第一次循环前执行的语句，一般用于初始化变量，如果变量在此处被声明，其作用域将被局限在这个for循环内。
	// 初始化语句可以被忽略，分号不可以省略。

	// 打印 0-9 数字
	var num int

	for ; num < 10; num++ {
		fmt.Println(num)
	}
}

func main() {
	fmt.Println("循环")

	/*
		格式如下：
			for 初始语句;条件表达式;结束语句{
				循环体代码
			}

	*/

	// for 循环的初始语句---开始循环时执行的语句
	testFor1()
	// testFor2()
	// testFor3()
}
