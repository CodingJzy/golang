package main

import (
	"flag"
	"fmt"
)

func add(a, b int) int {
	return a + b
}

func add1(a, b int) (int, int) {
	return a, b
}

func add2() (a int, b int) {

	a = 520
	b = 1314
	return a, b

	// 当函数使用命名返回值时也，return 可与不写返回值列表，就像下面，效果是一样的
	// return
}

func add3() (a int, b int) {

	a = 1
	return a, 2
}

func testDeclare() {
	/*
		结构：
			func 函数名 （参数列表） （返回值列表） {
				函数体
			}

		函数名：
			由字母、数字、下划线组成，第一个字符不能为数字。在同一个包内，函数名不可以重复

		参数列表：一个参数由参数变量和参数类型组成，例如：
			func foo(a int, b string) {

			}

		返回值列表：可以是返回值类型列表，也可以是类似参数列表的组合。函数声明有返回值时，函数体必须return提供返回值列表

		函数体：
			能够被重复调用的代码片段
	*/

	// 参数类型的简写：当两个参数的类型相同的时候，具体看add函数
	sum := add(1, 3)
	fmt.Println(sum)

	// 函数的返回值：go 语言支持返回多个值，具体看add1函数
	x, y := add1(1, 100)
	fmt.Println(x, y)

	// 同一种类型的返回值：如果返回值是同种类型，具体看add1函数

	// 纯类型的返回值对代码可读性不是很友好，其实它可以和参数列表一样，有变量名。
	x, y = add2()
	fmt.Println(x, y)
	x, y = add3()
	fmt.Println(x, y)

	// 函数的调用：和其他编程语言一样，都是函数名()来进行调用

}

func testSave() {
	// 函数也是一种类型，可以和其他类型一样保存在变量中。

	f := add3

	// 根据声明的函数名保存的变量名来调用
	x, y := f()
	fmt.Println(x, y)
}

func visit(list []int) {
	for _, v := range list {
		fmt.Println(v)
	}
}

func visit1(list []int, f func(int)) {
	for _, v := range list {
		f(v)
	}
}

func testAnonymous() {
	// 匿名函数就是没有函数名的函数，只有函数体。常用于回调函数和闭包

	// 定义一个匿名函数
	/*
		结构：
			func (参数列表) (返回值列表) {
				函数体
			}

		可以当做没有函数名字的普通函数定义
	*/

	// 在定义时调用匿名函数
	func(str string) {
		fmt.Println(str)
	}("我是匿名函数，可以在定义时末尾加括号直接调用。。。")

	// 整个匿名函数也可以赋值给一个变量
	nm := func(s string) {
		fmt.Println(s)
	}
	//调用此变量名
	nm("通过声明匿名函数的变量名来调用")

	// 匿名函数作为回调函数。其实还是有点不明白该例子的意义。我用一个函数直接调用不是更省事吗？为啥非要加个匿名函数。
	visit([]int{1, 2, 3, 4, 5})
	fmt.Println("----------------------------------")
	visit1([]int{5, 4, 3, 2, 1}, func(v int) {
		fmt.Println(v)
	})

	// 使用匿名函数进行封装
	// 定义一个命令行参数 skillParam(全局变量)
	// 命令行参数解析
	flag.Parse()
	// 定义一个map， skill为key, value 为执行该技能的方法
	skill := map[string]func(){
		"run": func() {
			fmt.Println("我是野马会奔跑")
		},
		"sleep": func() {
			fmt.Println("我是猪会睡觉")
		},
		"fly": func() {
			fmt.Println("我是大鹏会飞翔")
		},
	}

	// 判断 通过对skill的指针地址取值
	if f, ok := skill[*skillParam]; ok {
		f()
	} else {
		fmt.Println("skill not found")
	}
	// 运行命令：go run main.go   --skill=run
}

// 声明一个命令行全局变量
var skillParam = flag.String("skill", "", "skill to perform")

func main() {
	fmt.Println("函数")

	// 声明函数
	// testDeclare()

	// 函数变量---把函数作为值保存到变量中
	// testSave()

	// 匿名函数
	testAnonymous()

	// 函数类型实现接口---把函数作为接口来调用

	// 闭包---引用了外部变量的匿名函数

	// 可变参数---参数数量不固定的函数形式

	// 延迟执行语句(defer)

	// 处理运行时发生的错误

	// 宕机(panic)---程序终止运行

	// 宕机恢复(recover)---防止程序崩溃
}
