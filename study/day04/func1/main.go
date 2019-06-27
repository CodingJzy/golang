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

// 定义一个函数，返回值为一个匿名函数，匿名函数与原函数形成闭包，把name 、hp 返回
func palyerGen(name string) func() (string, int) {

	hp := 150

	return func() (string, int) {
		return name, hp
	}
}

func testClosure() {
	// 闭包 = 函数 + 引用环境
	// 函数是编译期静态的概念，而闭包是运行期动态的概念

	// 闭包内部修改引用的变量
	str1 := "hello world"
	fmt.Printf("修改之前---%s\n", str1)
	// 创建一个匿名函数,并直接调用
	func() {
		str1 = "hello go"
	}()
	fmt.Printf("修改后---%s\n", str1)
	// 匿名函数中并没有定义str1，也不是通过参数传递的方式。就算通过传递参俗话的方式。由于字符串不可变。不会对原有字符串发生改变。
	// str1 的定义在匿名函数之前。此时，str1就被引用到匿名函数中形成了闭包

	// 闭包实现生成器
	// 调用函数
	gen := palyerGen("天使彦")
	// 返回天使彦的名字和血量
	name, hp := gen()
	// 输出
	fmt.Println(name, hp)
	// 可以看出闭包还具有一定的封装性，血量是无法从外部修改的，与面向对象的封装类似。限制外部内内部的访问权限
}

func studentInfo1(name string, age int, args ...string) {
	fmt.Println("固定参数", name, age)
	fmt.Println("不固定参数", args)

	// 通过遍历获取每一个参数
	for index, arg := range args {
		fmt.Printf("索引：%d  参数值：%s\n", index, arg)
	}
}

func studentInfo2(name string, args ...interface{}) {
	fmt.Println("固定参数", name)
	fmt.Println("不固定参数", args)

	// 通过遍历获取每一个参数的类型
	for k, v := range args {

		// 对参数的值进行格式化
		value := fmt.Sprintf("%v", v)

		fmt.Printf("索引：%d  \t参数的类型：%T\t参数值：%s\n", k, v, value)
	}

}

func printInfo(slices ...interface{}) {

	for index, arg := range slices {
		value := fmt.Sprintf("%v", arg)
		fmt.Printf("索引：%d  参数值：%s\n", index, value)
	}
}

func studentInfo3(slices ...interface{}) {
	// 将slices 可变参数切片完整的传递给下一个打印函数
	printInfo(slices...)
}

func testArgs() {
	// 有时候我们不知道我们传的参数有多少个，可以通过不固定参数来解决。类似python中的 *args **kwargs

	/*
		格式：
			func 函数名 (固定参数列表, v ... T) (返回参数列表) {
				函数体
			}
		v:
			为可变参数变量，类型为[]T ，也就是拥有多个元素的类型切片 ，v 和 t 之间用 ...
		T:
			为可变参数类型。当类型为interface{}的时候，传入的值可以是任意值
	*/

	// 可变参数为多个字符串元素的切片类型。并循环遍历每个参数
	// studentInfo1("江子牙", 21, "江西省", "萍乡市", "莲花县", "坊楼镇", "小江村")

	// 可变参数为interface{}类型。并循环遍历每个参数的类型
	// studentInfo2("江子牙", 22, 99.9, "sex", true, map[string]string{
	// 	"className": "高三(1)班",
	// 	"score":     "99",
	// })

	// 可变参数本身就是一个切片类型，当我们想传多个参数列表的切片的时候，可以这样

	studentInfo3("江子牙", 22, 99.9, "sex", true, map[string]string{
		"className": "高三(1)班",
		"score":     "99",
	})

}

// 声明一个命令行全局变量
var skillParam = flag.String("skill", "", "skill to perform")

func main() {
	fmt.Println("函数基础")

	// 声明函数
	// testDeclare()

	// 函数变量---把函数作为值保存到变量中
	// testSave()

	// 匿名函数
	// testAnonymous()

	// 函数类型实现接口---把函数作为接口来调用

	// 闭包(Closure)---引用了外部变量的匿名函数
	// testClosure()

	// 可变参数---参数数量不固定的函数形式
	// testArgs()
}
