package main

import (
	"fmt"
)

func testDefer() {
	// 类似于栈，先进后出。先defer的语句最后执行，后defer的语句最先执行
	// 延时语句会在所在函数结束时进行。函数结束可以是正常返回时，也可以是发生宕机时。

	fmt.Println("defer begin")

	// 定义一个defer语句
	defer fmt.Println(1)

	// 再此定义一个defer语句
	defer fmt.Println(2)

	// 定义最后一个defer语句
	defer fmt.Println(3)

	fmt.Println("defet end")
	// 使用defer可以很好的处理函数结束或退出时资源未释放的问题。比如打开文件、接收请求、回复请求、加锁、解锁等
}

func testPanic() {
	// 可以手动触发宕机。让程序崩溃。开发者能及时的发现错误，同时减少可能的损失
	// panic("crash")

	// 在宕机时触发延迟执行语句
	defer fmt.Println("宕机要做的第二件事")
	defer fmt.Println("宕机要做的第一件事")

	panic("发生宕机了")
}

func restRevover() {
	// 无论是代码运行错误还说抛出的宕机错误。还是主动触发的宕机错误，都可以配合defer和recover实现错误捕捉和恢复，让代码崩溃后继续运行

	// 延时语句捕捉宕机
	defer func() {
		fmt.Println("开始捕捉")
		err := recover()
		fmt.Println("捕捉成功，错误为：", err)
	}()

	panic("发生宕机了，之后执行defer语句，defer语句调用err()函数，函数内部recove()捕捉错误，继续执行")

}

func main() {
	// fmt.Println("函数高级")
	// 延迟执行语句(defer)
	// testDefer()

	// 处理运行时发生的错误

	// 宕机(panic)---程序终止运行
	// testPanic()

	// 宕机恢复(recover)---防止程序崩溃
	restRevover()
	fmt.Println("捕捉错误成功，程序继续向下执行")
}
