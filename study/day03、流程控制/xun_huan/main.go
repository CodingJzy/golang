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

func testFor2() {
	// 对每次循环前开始前计算的表达式，如果bool值为true，则继续往下执行。否则结束循环。
	// 条件表达式可以忽略，但是循环默认成了无限循环

	// 打印 5-10 的数字
	num := 10
	for ; num >= 5; num-- {
		fmt.Println(num)
	}

	// 无限循环
	for ; ; num++ {
		fmt.Println(num)
		if num == 10 {
			break
		}
	}

	fmt.Printf("-----------------%d---------------------\n", num)

	// 更美观的写法：省略初始语句、条件表达式、结束语句
	for {
		fmt.Println(num)
		if num == 100 {
			break
		}

		num++
	}

	num = 0
	// if 和 for 进一步整合
	for num != 10 {
		fmt.Println(num)
		num++
	}

}

func testFor3() {
	// 结束语句：在每次循环结束之前执行的语句。但是当循环碰break、goto、return、panic等语句强制退出时，结束语句不会执行。
}

func jiujiu() {
	for i := 1; i < 10; i++ {
		for j := 1; j <= i; j++ {
			fmt.Printf("%d * %d = %d\t", i, j, i*j)
		}
		fmt.Println()
	}
}

func forArray() {
	array1 := [...]string{"Python", "Go", "Java", "Linux"}

	for k, v := range array1 {
		fmt.Println(k, v)
	}

	slice1 := []string{"Python", "Go", "Java", "Linux"}
	slice1 = append(slice1, "C++")
	for k, v := range slice1 {
		fmt.Println(k, v)
	}
}

func forString() {
	str1 := "Life Is Short You Need Python --- 江子牙"

	// v 变量其实是rune类型， 也就是int32 可以用%c占位符来显示
	for k, v := range str1 {
		fmt.Printf("index: %d  val: %c\n", k, v)
	}

}

func forMap() {
	mp := map[string]int{
		"a": 1,
		"b": 2,
		"c": 3,
	}

	for k, v := range mp {
		fmt.Println(k, v)
	}
}

func testKV() {
	// 遍历数组、切片
	// forArray()

	// 遍历字符串
	// forString()

	// 遍历map
	forMap()

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
	// testFor1()

	// for 循环的条件表达式---控制是否循环的开关
	// testFor2()

	// for 循环的结束语句---每次循环结束时的执行语句
	// testFor3()

	// 九九乘法表
	// jiujiu()

	// 键值循环
	testKV()

}
