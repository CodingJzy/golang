package main

import (
	"fmt"
)

func main() {
	fmt.Println("空接口类型(interface{})----能保存所有值的类型")

	// 空接口是接口类型的特殊形式，空接口没有任何方法。因此任何类型都无须实现口，从实现的角度来看。任何值都满足这个接口的需求。因此空接口类型可以保存任何值，也可以从中取值。

	// 将各种数据类型的值保存到空接口

	var any interface{}

	any = 1
	fmt.Println(any)
	any = 0.99
	fmt.Println(any)
	any = "Hello Gp"
	fmt.Println(any)
	any = true
	fmt.Println(any)
	any = []string{}
	fmt.Println(any)
	any = [...]string{}
	fmt.Println(any)
	any = map[string]int{}
	fmt.Println(any)

	// 从空接口获取值
	// 虽然可以任意存值，但是：比如我们有变量a := 1，把他存到空接口中 var kong interface{} = a，这也没错。此时我们声明变量b := kong。则保错
	// 虽然存到里面的a的值是1，类型为int。但是kong依旧是一个interface{}类型的变量。
	// 就好比集装箱。他可以装烟草，也可以装茶叶等等等。但是集装箱始终是箱子。金属做的，无法改变。
	a := 1
	var kong interface{} = a
	fmt.Println(kong)
	// 根据类型断言，可解决上诉问题
	b := kong.(int)
	fmt.Println("根据类型断言解决了", b)

	// 空接口的值的比较：空接口可以保存不同的值后，可以和其他变量值一样使用"=="
	// 类型不同的空接口的比较：
	var kong1 interface{}
	kong1 = "hello"

	var kong2 interface{}
	kong2 = 1

	var kong3 interface{}
	kong3 = "hello"
	fmt.Println(kong1 == kong2, kong1 == kong3)

	// 不能比较空接口中的动态值：会触发错误
	// 不可比：字典、切片
	// 可比：通道、数组、结构体、函数

}
