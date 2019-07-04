package main

import (
	"fmt"
	"reflect"
)

// 定义一个Enum类型
type Enum int

const (
	Zero Enum = 0
)

type Cat struct {
}

func testKind() {
	cat := new(Cat)
	var cat1 Cat
	cat2 := &Cat{}
	typeOfCat := reflect.TypeOf(cat)
	typeOfCat1 := reflect.TypeOf(cat1)
	typeOfCat2 := reflect.TypeOf(cat2)
	typeOfZero := reflect.TypeOf(Zero)
	fmt.Println(typeOfCat.Name(), typeOfCat.Kind())   //
	fmt.Println(typeOfCat1.Name(), typeOfCat1.Kind()) // Cat struct
	fmt.Println(typeOfCat2.Name(), typeOfCat2.Kind()) //      ptr
	fmt.Println(typeOfZero.Name(), typeOfZero.Kind()) // Enum int
}

func testEle() {
	// 定义一个狗的结构体
	type Dog struct {
	}

	// 实例化一个狗
	dog1 := new(Dog)
	dog2 := &Dog{}
	dog3 := Dog{}

	// 获取狗结构体实例的反射类型对象
	typeOfDog1 := reflect.TypeOf(dog1)
	typeOfDog2 := reflect.TypeOf(dog2)
	typeOfDog3 := reflect.TypeOf(dog3)

	// 显示反射类型对象的名称和种类
	fmt.Printf("种类：%v\t名称：%v\n", typeOfDog1.Kind(), typeOfDog1.Name())
	fmt.Printf("种类：%v\t名称：%v\n", typeOfDog2.Kind(), typeOfDog2.Name())
	fmt.Printf("种类：%v\t名称：%v\n", typeOfDog3.Kind(), typeOfDog3.Name())

	// 取指针指向的元素的类型，不可以通过一个非指针类型获取它的指针类型
	// typeOfDog3并不是一个指针类型。
	typeOfEle1 := typeOfDog1.Elem()
	typeOfEle2 := typeOfDog2.Elem()

	// 显示反射类型对象的名称和种类
	fmt.Printf("ele name：%v\tele kind：%v\n", typeOfEle1.Name(), typeOfEle1.Kind())
	fmt.Printf("ele name：%v\tele kind：%v\n", typeOfEle2.Name(), typeOfEle2.Kind())
}

func testField() {
	// 任意值获取反射对象信息后，如果它的类型是结构体，可以通过反射对象的NumField()和Field()获得结构体的成员详细信息。
	/*

		reflect.Type主要方法：
			Field(i int) StructField：根据索引，返回索引对应的结构体字段的信息。

			NumField() int：返回结构体成员字段数量，类型不是结构体或者索引越界会发生宕机。

			FieldByName(name string) (StructField, bool)：根据给定字符串返回结构体相应字符串的信息。返回值 bool和该字段类型

			FieldByIndex(index []int) StructField：多层成员访问

			FieldByNameFunc(match func(string) bool) (StructField, bool)：根据匹配函数匹配需要的字段。

		结构体字段类型
		type StructField struct {
			Name     	string     // 字段名
			PkgPath  	string 	   // 字段路径
			Type	 	Type	   // 字段反射类型对象
			Tag		 	StructTag  // 字段的结构体标签
			Index	 	[]int	   // Type.FieldByIndex中的返回的索引值
			Anonymous 	bool	   // 是否为匿名字段
		}
	*/

	// 获取成员反射信息
	// 声明一个人结构体
	type Person struct {
		Name string `a:1 b:"b"`
		Type int    `json:"type" id:"1"` // 带有结构体tag的字段
	}
	// 实例化一个实例
	p1 := Person{Name: "jiang_wei", Type: 1}

	// 获取p1的反射类型对象
	typeOfP1 := reflect.TypeOf(p1)
	fmt.Println(typeOfP1.Name(), typeOfP1.Kind())
	//遍历结构体所有成员
	for i := 0; i < typeOfP1.NumField(); i++ {
		// 获取每个成员的结构体字段类型
		filedType := typeOfP1.Field(i)

		// 输出成员名和Tag
		fmt.Printf("字段名称：%v\t 字段标签：%v\n", filedType.Name, filedType.Tag)
	}

	// 通过字段名，找到字段类型信息
	if p1Type, ok := typeOfP1.FieldByName("Type"); ok {
		// 从p1Type中取出需要的tag
		fmt.Println(p1Type.Tag.Get("json"), p1Type.Tag.Get("id"))
	}
}

func testStructTag() {
	// 结构体标签是对结构体字段的额外信息标签

	type Phone struct {
		Name string
		Type string `json:"type" id:"100"`
	}

	// 取
	phone1 := Phone{
		"苹果x",
		"iPhone",
	}
	typeOfPhone1 := reflect.TypeOf(phone1)
	if phoneType, ok := typeOfPhone1.FieldByName("Type"); ok {
		fmt.Println(phoneType.Name)
		fmt.Println(phoneType.Type)
		fmt.Println(phoneType.Tag.Get("json"))
		fmt.Println(phoneType.Tag.Get("id"))
	}
}

func main() {
	// 反射种类的定义

	// 种类(Kind)指的是对象归属的品种
	// Map、Slice、Chan 属于引用类型，使用起来类似于指针。但是在种类常量定义中仍然属于独立的种类，不属于Ptr。
	// type A struct{} 定义的结构体属于 Struct种类。*A属于Ptr，是指针类型。
	//testKind()

	// 指针与指针指向的元素
	//testEle()

	// 使用反射获取结构体的成员类型
	//testField()

	// 结构体标签
	//testStructTag()
}
