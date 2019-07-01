package main

import (
	"fmt"
)

func testSuanShu() {
	n1 := 100
	n2 := 200
	fmt.Println("相加", n1+n2)
	fmt.Println("相减", n1-n2)
	fmt.Println("相乘", n1*n2)
	fmt.Println("相除", n1/n2)
	n1++
	fmt.Println("自增", n1)
	n2--
	fmt.Println("自减", n2)
}

func testGuanxi() {
	n1 := 1
	n2 := 2

	fmt.Println(n1 == n2)
	fmt.Println(n1 > n2)
	fmt.Println(n1 < n2)
	fmt.Println(n1 >= n2)
	fmt.Println(n1 <= n2)
	fmt.Println(n1 != n2)
}

func testLuoji() {
	b1 := true
	b2 := false

	fmt.Println(b1 && b2)
	fmt.Println(b1 || b2)
	fmt.Println(!b1)
	fmt.Println(!b2)
}

func main() {

	// 算术运算符
	// testSuanShu()

	// 关系运算符
	// testGuanxi()

	// 逻辑运算符
	testLuoji()

}
