package main

import (
	"fmt"
)

func testDesc() {
	str1 := `
go语言通过自定义的方式形成新的类型, 结构体是类型中带有成员的复合类型。
go语言使用结构体和结构体成员来描述真实世界的实体和实体对应的各种属性。
结构体成员是由一系列成员变量构成，这些成员变量也被称为字段。字段有以下特征：
	1、字段拥有自己的类型和值
	2、字段名必须唯一
	3、字段的类型也可以是结构体、甚至是字段所在结构体的类型
go语言中没有类的概念，也不支持类的继承等面向对象的概念, 但在结构体的内嵌配合接口比面向对象更具更高的扩展性和灵活性。
go语言不仅认为结构体能拥有方法，且每种自定义类型也可以拥有自己的方法。
	`
	fmt.Println(str1)
}

func testDeclare() {
	// 关键字type可以将各种基本数据类型定义为自定义类型。基本类型包括整型、字符串、布尔等，结构体是一种复合的基本类型。通过type定义为自定义类型后，使结构体更便于使用。
	/*
		格式如下：
			type 类型名 struct {
				字段1  字段1类型
				字段2  字段2类型
				字段3  字段3类型
			}

		类型名：
			表示自定义结构体的名称，在同一个包内不能重复。

		struct {}：
			表示结构体类型，可以理解为把该结构体定义为类型名的类型。

		字段1、字段2：
			表示结构体字段名。结构题中的名字必须唯一。
	*/

	// 简单的结构体
	type Point struct {
		X int
		Y int
	}

	// 同种类型的字段可以放一行，颜色发红绿蓝
	type Color struct {
		R, G, B byte
	}

}

func testNew() {
	// 结构体的定义只是一种内存布局的描述, 此时并不会分配内存。只有实例化的时候才会真正分配内存
	// 所以，在定义结构体并实例化才能使用结构体字段，有多重实例方式

	// 基本的实例化形式：结构体本身是一种类型，可以像声明变量一样，以 var 的方式声明结构体即可完成实例化
	// 格式：var ins T

	// 只是定义结构体
	type Point struct {
		X, Y int
	}

	// 实例化结构体
	var p Point
	p.X = 10
	p.Y = 10

	// 创建指针类型的结构体：ins := new(T)
	type Person struct {
		name string
		age  int
	}

	per := new(Person)
	per.name = "江子牙"
	per.age = 22
	// 类型为 *Person
	fmt.Printf("类型为：%T", per)

	// 取结构体地址的实例化
	// 定义一个命令行指令，指令包含名称、变量、注释等。

	// 定义Command结构体， 表示命令行指令
	type Command struct {
		Name string
		// 命令绑定的变量，使用整型指针绑定一个指针。指针的值可以与绑定的值随时保持同步
		// 根据取出来的指针取出该指针指向的值
		Var     *int
		Comment string
	}

	version := 1
	// 对结构体进行取地址实例化
	cmd := &Command{}
	// 初始化成员字段
	cmd.Name = "version"
	// 对变量进行取指针的操作
	cmd.Var = &version
	cmd.Comment = "show version"
}

func testInit() {
	// 键值对初始化结构体的书写格式
	// 可以理解为python中的关键字传参。

	type Book struct {
		Name   string
		Price  float64
		Author []string
	}
	b := Book{
		Name:   "Go 入门到进阶",
		Price:  99.99,
		Author: []string{"江子牙", "小喵喵", "小豹子"},
	}
	fmt.Println(b)

	type People struct {
		Name  string
		Child *People
	}

	p := People{
		Name: "爷爷",
		Child: &People{
			Name: "爸爸",
			Child: &People{
				Name: "儿子",
			},
		},
	}
	fmt.Println(p.Child.Child.Name)

	// 使用多个值的列表初始化结构体书写格式
	// 可以理解为python中的函数位置传参，顺序要一一对应。不能漏参，不能混用

	type Address struct {
		Country  string
		Province string
		City     string
		County   string
		Town     string
		Zipcode  int
		Username string
		Phone    string
	}

	add := Address{
		"中国",
		"江西省",
		"萍乡市",
		"莲花县",
		"坊楼镇",
		337103,
		"江子牙",
		"13006293101",
	}
	fmt.Println(add)
}

func testAnonymous() {
	// 就是没有类型名称的结构体 无需用type声明就可以直接使用

	//匿名结构体定义格式并初始化
	u := struct {
		name string
		age  int
	}{
		"江子牙",
		21,
	}
	fmt.Println(u)
}

func main() {
	fmt.Println("结构体")

	// 结构体简介
	// testDesc()

	// 定义结构体
	// testDeclare()

	// 实例化结构体---为结构体分配内存并初始化
	// testNew()

	// 初始化结构体的成员变量
	// testInit()

	// 初始化匿名结构体
	// testAnonymous()

}
