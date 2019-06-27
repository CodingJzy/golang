package main

import (
	"fmt"
	"math"
	"unicode/utf8"
)

func testNum() {
	// 整型

	/*
		1、按长度分：int8  int16  int32 int64
		2、按对应的无符号整型：uint8  uint16  uint32  uint64

		unit8 就是byte类型
		int16 就是c语言中的short
		int64 就是c语言中的long
	*/
	n1 := 100
	n2 := 200

	fmt.Println(n1, n2)

	// 浮点型
	// float32  float64

	fmt.Printf("%f\n", math.Pi)
	fmt.Printf("%.2f\n", math.Pi)

}

func testBool() {
	flag := true
	var isLogin bool
	isLogin = false

	fmt.Println(flag, isLogin)
}

func testString() {
	// 字符串用双引号表示
	s1 := "hello world"
	s2 := "你好，世界"
	var s3 string
	fmt.Println(s1)
	fmt.Println(s2)
	fmt.Println(s3 + "-----------")

	// 字符串转义
	fmt.Println("1\r2\n3\t45\"\\6")

	// 定义多行字符串
	const s4 = `
	第一行
	第二行
	第三行
	\r\n
	`
	fmt.Println(s4)

	// 字符
	// 字符串中的每一个元素叫做字符
	/*
		unit8：byte类型，代表ascii码的一个字符
		rune：int32类型，代表utf8字符。
	*/

	var a byte = 'a'
	var b = '你'
	fmt.Printf("%d 类型为：%T\n", a, a)
	fmt.Printf("%d 类型为：%T\n", b, b)

	// 长度
	// len()：返回值为int类型，表示字符串的ascii字符个数的字节长度
	// utf8.RuneCountInString()：表示unicode字符串长度
	str1 := "hello python"
	str2 := "你好"
	fmt.Println(len(str1))
	fmt.Println(utf8.RuneCountInString(str2))
	fmt.P

	// 遍历
	// ascii字符可以通过for循环根据索引来获取
	// unciode字符则需要通过for range来获取
	str3 := "hello 江子牙"
	str4 := "你好：hello"
	for i := 0; i < len(str1); i++ {
		fmt.Printf("ascii遍历：%c  %d\n", str3[i], str3[i])
	}

	for _, s := range str4 {
		fmt.Printf("unicode遍历：%c  %d\n", s, s)
	}

	for _, s := range str3 {
		fmt.Printf("%c\n", s)
	}

	// 格式化
	/*
		%v	原来的值输出
		%%	输出%
		%T	输出类型和值
		%b	二进制显示
		%d  十进制显示
		%o	八进制显示
		%x	十六进制显示
		%f	浮点数显示
		%p	指针，十六进制显示
		%U	unicode字符
	*/

	title := fmt.Sprintf("%v  %T", "谷歌", "谷歌")
	fmt.Println(title)
}

func swap(a, b *int) {
	// 交换两个数的值函数

	// 取出a指针的值，存到tmp变量
	tmp := *a

	// 取出b指针的值，赋值给a指向的变量
	*a = *b

	// 临时变量tmp 赋值给b指针指向的值
	*b = tmp
}

func testPtr() {
	/*
		取地址：&
		根据取出的地址取出地址指向的值：*
	*/

	// 准备一个字符串类型
	house := "北京市昌平区豆各庄新村1栋"

	// 对字符串进行取地址操作， ptr 类型为 *string
	ptr := &house

	// 打印ptr的类型
	fmt.Printf("ptr 类型为：%T\n", ptr)

	// 打印ptr取到的地址值
	fmt.Printf("ptr 地址值为：%p\n", ptr)

	// 对ptr指针变量取到的地址值进行取值操作
	value := *ptr

	fmt.Printf("%s 类型为：%T\n", value, value)
	fmt.Printf("%s 值为：%s\n", value, value)

	// 例子：用指针操作对两个数的值进行交换操作
	a, b := 1, 2
	fmt.Println("交换前", a, b)
	swap(&a, &b)
	fmt.Println("交换后", a, b)

}

func testAlias()  {
	// 类型别名和类型定义的区别

	// 为int类型取一个别名
	type intAlias = int

	// 定义一个类型为int的newInt类型
	type newInt int

	var n1 newInt = 1
	var n2 intAlias = 2
	fmt.Printf("%d 类型为：%T\n", n1, n1)
	fmt.Printf("%d 类型为：%T\n", n2, n2)
}

func main() {
	// 数字
	// testNum()

	// 布尔
	// testBool()

	// 字符串
	// testString()

	// 指针
	// testPtr()

	// 类型别名
	testAlias()
}
