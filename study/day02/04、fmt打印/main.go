package main

import (
	"fmt"
)

func testPrintln() {
	s := "jiang_wei"
	i := 23
	weight := 55.5
	isJieHun := false
	fmt.Println("自动换行")
	fmt.Println("我的名字是", s, ", 我今年", i, "岁, ", "体重", weight, "kg, 婚姻：", isJieHun, "\n")
	fmt.Println("另起两行")
}

func testPrintf() {

	s := "life is short let's go"
	i := 100
	f := 3.1415926
	b := true
	slice := []int{1, 2, 3, 4}
	mp := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	// 通用型
	// %v：按照值的默认格式表示
	fmt.Printf("字符串：%v   整数：%v   小数：%v   布尔：%v   切片：%v   字典：%v\n", s, i, f, b, slice, mp)
	// %T：值的类类型
	fmt.Printf("字符串：%T   整数：%T   小数：%T   布尔：%T  切片：%T   字典：%T\n", s, i, f, b, slice, mp)
	// %%：打印百分号
	fmt.Printf("今天90%%会下雨\n")
	// %p :值得内存地址，指针地址

	// 布尔型
	fmt.Printf("the bool is %t\n", b)

	// 数字型
	i1 := -2
	i2 := 15
	fmt.Printf("带符号：%+d\n", i)
	fmt.Printf("十进制：%d\n", i1)
	fmt.Printf("十进制左填充：%8d\n", i1)
	fmt.Printf("十进制右填充：%-8d\n", i1)
	fmt.Printf("二进制：%b\n", i2)
	fmt.Printf("十六进制：%x\n", i2)
	fmt.Printf("unicode字符：%U\n", s[1])
	fmt.Printf("浮点数(默认6位)：%f\n", f)
	fmt.Printf("浮点数保留小数位：%.2f\n", f)
	fmt.Printf("最多几位表示：%.4g\n", f)

	// 字符串
	fmt.Printf("this is a str：%s\n", s)

}

func main() {
	fmt.Println("fmt模块的打印")

	// 按原来的数据类型输出。是什么就打印什么。每一个println就是一行。
	testPrintln()

	// 占位符打印
	testPrintf()
}
