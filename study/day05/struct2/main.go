package main

import (
	"fmt"
)

func genName(name string) *Cat {
	return &Cat{
		Name: name,
	}
}

type Cat struct {
	Name  string
	Color string
}

type BlackCat struct {
	Cat
}

type Bag struct {
	items []int
}

func testReolad() {
	c1 := genName("小豹子")
	fmt.Println(c1)
}

func genBlackCat(color string) *BlackCat {
	cat := &BlackCat{}
	cat.Color = color
	return cat
}

func testParent() {
	c1 := genBlackCat("豹纹色")
	fmt.Println(c1)
}

func insert1(b *Bag, itemid int) {
	b.items = append(b.items, itemid)
}

// (b *Bag)表示接收器
// 后面与函数写法类似
// 每个方法只能有一个接收器
func (b *Bag) insert2(itemid int) {
	b.items = append(b.items, itemid)
}

func (b Bag) insert3(itemid int) {
	fmt.Println("内部会将结构体的数据复制一份：", b.items)
	b.items = append(b.items, itemid)
	after := fmt.Sprintf("%v", b.items)
	fmt.Printf("修改成功：%s 但是方法结束后不生效，", after)
}

func testMethod() {
	// go 语言中的方法method是一种作用于特定类型变量的函数，这种特定类型变量叫做接收器recevier
	// 如果特定类型理解为结构体或类时，尽管go语言没有面向对象。接收器的概念就类似于其他语言的 this 或者 self

	// 面向过程实现
	bag1 := new(Bag)
	insert1(bag1, 1001)
	fmt.Println(bag1)

	// 结构体method实现
	bag2 := &Bag{}
	bag2.insert2(1002)
	fmt.Println(bag2)

	// 接收器---方法作用的目标
	/*
		格式如下：
			func (接收器变量 接收器类型) 方法名 (参数列表) (返回值列表) {
				函数体
			}
		接收器变量：
			官方建议，该接收器变量最好和接收器类型的第一个字母保持一致。如 Cat c 、Bag b 、Dog d。。。
		接收器类型：
			可以是指针类型和非指针类型
				指针类型：
					指针类型的接收器是一个结构体的指针，更接近于面向对象的this or self。
					由于指针的特性，调用方法时，修改接收器指针的任意成员变量，在方法调用结束后，已经修改过得都是有效的。
					上个例子的接收器*Bag就是一个指针接收器，所以才能append之后保持修改。
				非指针类型：
					当方法作用与非指针接收器时，go语言内部会在代码运行的时候将接收器的数据复制一份，在非指针接收器的方法可以获取，也可以修改，只是修改无效而已。
	*/

	// 非指针接收器修改数据无效
	bag3 := Bag{}
	bag3.insert3(1003)
	fmt.Println("是不生效：", bag3)

	/*
		总结：
			计算机中，小对象由于复制时速度比较傲快，所以适用非指针接收器。大对象因为复制性能较低。使用指针接收器。在接收器和参数间传递时不进行复制，只是传递指针。
	*/

}

// 自定义一个类型
type MyInt int

// 为 MyInt 添加是否为0的方法
func (m MyInt) IsZero() bool {
	return m == 0
}

// 为 MyInt 添加add()方法
func (m MyInt) Add(other int) int {
	return int(m) + other
}

func testTypeMethod() {
	// go语言可以对任何一中类型添加方法。给一种类型添加方法就像给结构体添加方法一样。

	// 定义一个Myint类型 b = 0
	var b MyInt
	fmt.Println(b.IsZero())
	b = 100
	// 测试是否为0的方法
	fmt.Println(b.IsZero())
	// 测试两数相加方法
	fmt.Println(b.Add(150))

}

func main() {
	fmt.Println("结构体2")
	// 多种方式创建和初始化结构体---模拟构造函数重载
	// testReolad()

	// 带有父子关系的结构体的构造和初始化
	// testParent()

	// 方法
	// testMethod()

	// 为类型添加方法
	testTypeMethod()

	// 使用事件系统实现事件的响应和处理

	// 类型内嵌和结构体内嵌

}
