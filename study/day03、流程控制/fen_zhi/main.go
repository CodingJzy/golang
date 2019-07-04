package main

import (
	"fmt"
)

func test1() {

	str1 := "how"
	switch str1 {
	case "how":
		fmt.Println(1)
	case "are":
		fmt.Println(2)
	default:
		fmt.Println(3)
	}
}

func test2() {
	str1 := "go"
	switch str1 {
	case "python", "go":
		fmt.Println("haha")
	case "java":
		fmt.Println("laji")
	}
}

func test3() {

	// 这种情况下，swtich 不跟判断变量，连判断的目标都没有了

	num := 98
	switch {
	case num > 80 && num < 100:
		fmt.Println("优秀")
	case num > 60 && num < 80:
		fmt.Println("一般")
	case num < 60:
		fmt.Println("不及格。去死吧。")
	}
}

func main() {
	//  分支选择(switch)--- 拥有多个条件分支的判断

	// 标准写法
	// test1()

	// 一分支多值
	// test2()

	// 分支表达式
	// test3()

}
