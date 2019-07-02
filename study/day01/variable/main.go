package main

import "fmt"

func declare() {

	// 1.1、标准格式
	// var 变量名 类型
	var a string
	var b int
	var c bool

	// 1.2、批量格式
	var (
		a1 string
		b1 int
		c1 bool
	)

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(a1)
	fmt.Println(b1)
	fmt.Println(c1)
}

func initVar() {
	// 2.1、标准格式
	// var 变量名 变量类型 = b表达式
	var name string = "jiang_wei"
	var age int = 23
	var money float64 = 520.1314

	fmt.Println(name, age, money)

	// 2.2、批量格式
	// 标准格式 左右两边冗余 ，左边的int可以省略
	var name1 = "jiang_zi_ya"
	var age1 = 22
	var money1 = 9.99

	fmt.Println(name1, age1, money1)

	// 2.3 短变量模式初始化
	// 使用该方式必须在函数内部，所以var常用在声明全局变量。声明并初始化变量时。变量名必须未定义过，否则会报错
	language := "Go"
	paiming := 1

	fmt.Println(language, paiming)

}

func sum() (int, int) {

	return 100, 200
}

func anonymous() {
	// 可以通过_来使用匿名变量。匿名变量不会分配内存，也不会因为多次声明而无法使用
	a, _ := sum()
	_, b := sum()

	fmt.Println(a, b)
}

func testConst() {
	// 常量固定不变
	const pi = 3.14159
	const e = 2.71

	// 批量声明
	const (
		pi1 = 1.1
		e1  = 2.2
	)

	// 枚举：go语言现阶段没有枚举，但是可以用常量配合iota来模拟枚举

	const (
		aa int = iota
		bb
		cc
	)
	fmt.Println(aa, bb, cc)

}

func main() {

	// 1、声明变量
	// declare()

	/*
		2、初始化变量
			- 整型和浮点型变量的默认值为0
			- 字符串变量的默认值为空字符串
			- 布尔型变量默认为 false
			- 切片、函数、指针变量的默认为 nil
		当然，依然可以在变量声明时赋予变量一个初始值
	*/
	// initVar()

	// 3、匿名变量
	// anonymous()

	// 4、常量
	testConst()

}
