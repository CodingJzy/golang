## go语言的指针

### 简介

在c语言中，指针是很难理解的一个知识。在go语言中，相关操作也简化了好多。

基本数据类型，变量存的就是值，也就是值类型。值类型都有对应的指针类型，形式为*值类型。

### 概念

- 指针类型：`*int`、`*int64`、`*string`、`*struct`。
- 指针地址：对变量进行取地址操作用`&`符合。
- 指针取值：根据取到的地址取其对应的值用`*`符号。

例如：

```go
package main

import "fmt"

func main() {

	// 初始化一个int的变量num
	var num int = 10
	// 对num进行取地址操作
	var ptr1 *int = &num

	// ptr1的类型：*int	ptr1的值 ：0xc00 0086010
	fmt.Printf("ptr1的类型：%T\tptr1的值：%p\n", ptr1, ptr1)

	// 根据取到的地址取其对应的值
	tmp := *ptr1
	// 10
	fmt.Println(tmp)
}

```

### 示例1

```go
package main

import "fmt"

func swap(a, b *int) {

	// 临时变量存取a指向的值
	tmp := *a

	// b指向的值赋值给a
	*a = *b

	// 临时变量赋值给b
	*b = tmp

}

func main() {
	// 利用指针交换两个数的值

	a := 100
	b := 200

	fmt.Println("交换前：", a, b)
	swap(&a, &b)
	fmt.Println("交换后：", a, b)
}
```



