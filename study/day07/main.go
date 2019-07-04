package main

import (
	"fmt"
	"golang/study/day07/mylib"
)

// 1、全局变量
var a1 = 9

// 2、init()
func init() {
	fmt.Println("开始执行init()")
	a2 := 10
	fmt.Printf("输出局部变量%d\n", a2)
	fmt.Printf("在init()里输出全局变量%d\n", a1)
	fmt.Println("准备执行main()")
}

// 3、main()
func main() {
	fmt.Println("开始执行main")
	fmt.Printf("在main()里输出全局变量%d\n", a1)

	// 调用mylib包下的Add()加法函数
	print(mylib.Add(10, 10))
}
