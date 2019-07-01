package main

import "fmt"

func main() {
	//  在go语言中，对于引用类型的变量，使用的时候不仅要提前声明，还要为他分配内存空间，不然会引起宕机。对于值类型，声明的时候默认为他分配了默认值，分配内存有两种方式：make、new

	/*
		new:
			func new(Type) *Type

		Type：表示new函数只接受一个参数，并且是接收一个类型。
		*Type：表示new函数返回一个指针类型，指向该类型内存地址的指针。
	*/

	a := new(int)
	b := new(bool)

	// 输出两个变量的类型
	fmt.Printf("%T \n", a) // *int
	fmt.Printf("%T \n", b) // *bool
	// 根据指向该内存地址的指针获取值
	fmt.Println(*a) // 0
	fmt.Println(*b) // false

	/*
		make：
			func make(t Type, size ...IntegerType) Type

		使用切片、map、channel时都需要用make来进行初始化。
	*/
	mp1 := make(map[string]int, 4)
	fmt.Println(mp1)
	mp1["a"] = 1
	mp1["b"] = 2
	mp1["c"] = 3
	mp1["d"] = 4
	mp1["e"] = 5
	fmt.Println(mp1)

	/*
		区别：
			1、二者都是用来分配内存
			2、make只用于slice、map以及channel的初始化，返回的还是这三个引用类型本身
			3、new用于类型的内存分配，并且内存对应的值为类型零值，返回的是指向类型的指针。
	*/
}
