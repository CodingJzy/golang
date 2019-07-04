package main

import (
	"fmt"
)

func printType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("int类型")
	case string:
		fmt.Println("string类型")
	case bool:
		fmt.Println("bool类型")
	default:
		fmt.Println("不知名类型")
	}
}

// 支付宝方式结构体
type Alipy struct {
}

// 现金方式方式结构体
type Cash struct {
}

// 刷脸的接口
type useFace interface {
	Face(string)
}

// 刷人民币值为100的接口
type useOneHundred interface {
	OneHundred(string)
}

// 为支付宝提供人脸支付的方法
func (a Alipy) Face(s string) {
	fmt.Println(s)
}

// 为现金支付提供支付100元的方法
func (c Cash) OneHundred(s string) {
	fmt.Println(s)
}

func printPay(payMethod interface{}) {
	var pay = payMethod
	switch payMethod.(type) {
	case useFace:
		face := pay.(*Alipy)
		face.Face("刷脸")
	case useOneHundred:
		cash := pay.(*Cash)
		cash.OneHundred("支付100毛爷爷")
	}
}

func main() {
	fmt.Println("类型分支---批量判断接口中变量的类型")

	// go语言的switch不仅可以像其他语言一样实现数值、字符的判断。还有一直特殊的用途---判断一个接口内保存或实现的类型。

	// 使用类型分支判断基本类型
	printType("str")
	printType(1)
	printType(true)
	printType([]string{})

	//使用类型分支判断接口类型
	printPay(new(Alipy))
	printPay(new(Cash))

}
