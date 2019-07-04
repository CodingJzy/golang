package main

import (
	"fmt"
)

// 定义一个鸭子类型接口
type Duck interface {
	// 鸭子叫
	GaGa()
	// 鸭子走
	YouYong()
}

// 定义一个小鸡类型的结构体
type Chicken struct {
}

// 为小鸡结构体实现一个方法(method)
func (c Chicken) GaGa() {
	fmt.Println("我是小鸡，但我也会嘎嘎叫")
}

// 为小鸡结构体实现一个方法(method)
func (c Chicken) YouYong() {
	fmt.Println("我是小鸡，但我也会游泳")
}

func main() {
	fmt.Println("实现接口的条件")
	// 接口定义后，需要实现接口，调用方才能正确编译通过并使用接口。接口的实现必须遵循两条规则才能让接口可用

	// 1、接口的方法与实现接口的类型方法格式一致
	// 2、接口的方法与实现接口的类型方法数目保持一致。即接口中的所有方法均被实现。

	// 实例化一个小鸡
	c := new(Chicken)

	// 声明一个小鸡小鸭类型
	var duckChicken Duck = c

	// 小鸡小鸭合体
	// 小鸡小鸭开始嘎嘎
	duckChicken.GaGa()
	// 小鸡小鸭开始游泳
	duckChicken.YouYong()

}
