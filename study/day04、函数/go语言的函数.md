# go语言的函数

[TOC]



## 函数基础

### 简介

在任何编程语言中，函数就是对功能的封装，是组织好的，可重复使用，用于执行指定功能的代码块。在go语言中，函数属于一等公民。

### 定义

在go语言当中，用`func`定义一个函数，格式：

```go
func 函数名 (参数) (返回值) {
    函数体
}
```

- 函数名：由数字、字母、下划线组成。但是不能以数字开头，在同一个包内，函数名不能重复。
- 参数：一个参数由参数变量和参数类型组成，多个参数之间用`,`分隔。
- 返回值：可以是返回值类型，也可以是返回值变量和返回值类型组成。多个返回值用`()`包裹。并`,`分隔。
- 函数体：实现指定功能的代码块。

定义一个求两数之和的函数并调用：

```go
package main

import (
	"fmt"
)

// 定义一个求两数之和的函数
func add(a int, b int) int {
	return a + b
}

func main() {
    // 调用
	res := add(100, 200)
	// 300
	fmt.Println(res)
}
```

### 参数

#### 参数简写

当两个相邻参数的类型相同时，可以简写，例如：

```go
func add(a, b int) int {
	return a + b
}
```

#### 可变参数

有时候我们不知道我们传的参数有多少个，可以通过不固定参数来解决。在参数名后面加`...`来标识。类似python中的` *args` 、`**kwargs`。

格式：

```go
func 函数名 (固定参数, v ... T) (返回值) {
				函数体
			}
```

- v：可变参数变量，类型为`[]T` ，也就是拥有多个元素的类型切片 ，v 和 t 之间用 `...`。
- T：可变参数变量类型，当类型为`interface{}`的时候，传入的值可以是任意值。

##### 可变参数为多个字符串

```go
package main

import (
	"fmt"
)

func studentInfo(name string, age int, args ...string) {
	fmt.Println("固定参数：", name, age)
	fmt.Println("不固定参数：", args)

	// 通过遍历获取每一个参数
	for index, arg := range args {
		fmt.Printf("索引：%d  参数值：%s\n", index, arg)
	}
}
func main() {
	studentInfo("江子牙", 21, "江西省", "萍乡市", "莲花县", "坊楼镇", "小江村")
}
```

执行结果：

```go
固定参数： 江子牙 21
不固定参数： [江西省 萍乡市 莲花县 坊楼镇 小江村]
索引：0  参数值：江西省
索引：1  参数值：萍乡市
索引：2  参数值：莲花县
索引：3  参数值：坊楼镇
索引：4  参数值：小江村
```

##### 可变参数为空接口

```go
package main

import (
	"fmt"
)

func studentInfo(name string, args ...interface{}) {
	fmt.Println("固定参数", name)
	fmt.Println("不固定参数", args)

	// 通过遍历获取每一个参数的类型
	for k, v := range args {

		fmt.Printf("索引：%d  \t参数的类型：%T\t参数值：%v\n", k, v, v)
	}

}
func main() {
	studentInfo("江子牙", 22, 99.9, "sex", true, map[string]string{
		"className": "高三(1)班",
		"score":     "99",
	})
}
```

执行结果：

```go
固定参数 江子牙
不固定参数 [22 99.9 sex true map[className:高三(1)班 score:99]]
索引：0  	参数的类型：int	参数值：22
索引：1  	参数的类型：float64	参数值：99.9
索引：2  	参数的类型：string	参数值：sex
索引：3  	参数的类型：bool	参数值：true
索引：4  	参数的类型：map[string]string	参数值：map[className:高三(1)班 score:99]
```

本质上，可变参数是根据切片来实现的。

### 返回值

#### 多返回值

go语言是支持多个返回值，如果有多个返回值就要用`()`包裹。例如：

```
package main

import "fmt"

// 定义一个求两数之和、两数之差的函数
func cacl(a, b int) (int, int) {
	sum := a + b
	sub := a - b
	return sum, sub
}

func main() {
	sum, sub := cacl(10, 10)
	// 两数之和 20
	fmt.Println("两数之和", sum)
	// 两数之差 0
	fmt.Println("两数之差", sub)
}
```

#### 返回值命名

返回值可以和参数一样，由返回值变量和返回值变量类型组成。最后由`return`关键字返回。

```go
package main

import "fmt"

// 定义一个求两数之和、两数之差的函数
func cacl(a, b int) (sum, sub int) {
	sum = a + b
	sub = a - b
	return
}

func main() {
	sum, sub := cacl(10, 10)
	// 两数之和 20
	fmt.Println("两数之和", sum)
	// 两数之差 0
	fmt.Println("两数之差", sub)
}
```

## 函数进阶

### 变量作用域

#### 全局变量

全局变量定义在函数外部的变量。它在程序整个运行周期内都有效，在函数中可以访问到全局变量。

```go
package main

import "fmt"

// 定义两个全局变量a 、b
var (
	a = 10
	b = 10
)

// 定义一个求两数之和、两数之差的函数
func cacl() (sum, sub int) {
	fmt.Println("两个全局变量：",a, b)
	sum = a + b
	sub = a - b
	return
}

func main() {
	sum, sub := cacl()
	// 两数之和 20
	fmt.Println("两数之和", sum)
	// 两数之差 0
	fmt.Println("两数之差", sub)
}

```

#### 局部变量

##### 函数内部定义的局部变量函数外部无法访问

```go
package main

import "fmt"

func testLocalVar() {
	local := "我是函数内部的局部变量，函数外部无法访问我"
	fmt.Println(local)
}

func main() {
	testLocalVar()
	//fmt.Println(local) 访问不到
}

```

##### 全局变量和局部变量同时存在

如果局部变量和全局变量同名，优先访问局部变量，就近原则。

```go
package main

import "fmt"

//定义全局变量num
var num int64 = 10

func testNum() {
	num := 100
	//函数中优先使用局部变量
	fmt.Println("局部变量的num：", num)
}
func main() {
	testNum()
	fmt.Println("全局变量的num：", num)
}
```

### 函数类型与函数变量

函数也是一种类型，也可以赋值给一个变量保存起来。

##### 函数变量

```go
package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func main() {
	f:= add
	fmt.Printf("类型：%T\n", f)
	sum := f(1,2)
	fmt.Println(sum)
}
```

可以看出函数也是一种类型，就像上面的`add`函数，用变量名保存起来之后，打印它的类型，`func(int, int) int`,函数变量名加`()`也可以调用。

执行结果：

```go
类型：func(int, int) int
3
```

##### 函数类型

既然函数也是一种类型，那么也可以声明一个函数类型。

```go
package main

import "fmt"

// 声明一个cacl函数类型
type cacl func(int, int) int

func add(a, b int) int {
	return a + b
}

func main() {
	var c cacl
	c = add
	fmt.Printf("类型：%T\n", c)
	sum := c(1, 2)
	fmt.Println(sum)
}

```

执行结果：

```go
类型：main.cacl
3
```

### 高阶函数

满足其中一个条件为高阶函数

- 函数作为参数传入
- 函数作为返回值返回

#### 函数作为参数

```go
package main

import "fmt"

func add(a, b int) int {
	return a + b
}

func cacl(a, b int, f func(int, int) int) int {
	sum := f(a, b)
	return sum
}

func main() {
	res := cacl(10, 10, add)
	fmt.Println(res)
}

```

#### 函数作为返回值

```go
package main

import (
	"errors"
	"fmt"
)

func sum(a, b int) int {
	return a + b
}

func sub(a, b int) int {
	return a - b
}

// 定义一个计算的函数，根据传入的操作字符串执行对应的加减
func cacl(do string) (func(int, int) int, error) {
	switch do {
	case "加":
		return sum, nil
	case "减":
		return sub, nil
	default:
		err := errors.New("不支持该操作")
		return nil, err
	}
}

func main() {
	do1, _ := cacl("加")
	sum := do1(1, 20)
	fmt.Println(sum)

	do2, _ := cacl("减")
	sub := do2(10, 8)
	fmt.Println(sub)

	do2, err := cacl("乘")
	fmt.Println(err)
}

```

执行结果：

```go
21
2
不支持该操作
```

### 匿名函数

在其他语言中，函数内部还可以定义函数，即函数嵌套。但是go语言函数内部不能再像以前那样定义函数了，只能定义匿名函数。顾名思义就是没有函数名字的函数，常用于回调函数和闭包。

格式：

```go
func (参数) (返回值) {
				函数体
			}

// 可以当做没有函数名字的普通函数定义
```

因为没有函数名了，不能和普通函数那样直接调用。匿名函数需要保存到某个变量或者立即作为函数执行。

##### 保存为变量

定义一个匿名函数之后，保存到变量中。

```go
package main

import "fmt"

func main() {
	add := func(a, b int) int {
		sum := a + b
		return sum
	}
	sum := add(1, 2)
	fmt.Println(sum)
}

```

##### 立即执行

定义一个匿名函数之后，后面直接加`()，`立即执行。

```go
package main

import "fmt"

func main() {
	sum := func(a, b int) int {
		sum := a + b
		return sum
	}(1, 2)
	fmt.Println(sum)
}

```

### 闭包

函数是编译期静态的概念，闭包是运行期动态的概念。

`闭包 = 函数 + 引用环境`

##### 闭包示例1

```go
package main

import "fmt"

func adder() func(int) int {
	var x int
	return func(y int) int {
		x += y
		return x
	}
}
func main() {
    // f是一个函数，应用了其外部作用域的x变量，此时f就是一个闭包，在f的声明周期内，x变量一直有效
	f := adder()
	fmt.Println(f(10)) //10
	fmt.Println(f(20)) //30
	fmt.Println(f(30)) //60

	f1 := adder()
	fmt.Println(f1(40)) //40
	fmt.Println(f1(50)) //90
}

```

##### 闭包示例2

```go
package main

import "fmt"

func main() {
	// 定义一个字符串变量
	str1 := "hello world"
	fmt.Printf("修改之前---%s\n", str1)

	// 定义一个匿名函数
	func() {
		str1 = "hello go"
	}()
	fmt.Printf("修改后---%s\n", str1)

}
```

匿名函数中并没有定义`str1`，也不是通过参数传递的方式。就算通过传递参数话的方式。由于字符串不可变。不会对原有字符串发生改变。但是却对`str1`进行了修改。

`str1` 的定义在匿名函数之前，此时，`str1`就被引用到匿名函数中形成了闭包。

##### 闭包示例3

```go
package main

import "fmt"

// 定义一个函数，返回值为一个匿名函数，匿名函数与原函数形成闭包，把name 、hp 返回
func Gen(name string) func() (string, int) {
	hp := 150
	return func() (string, int) {
		return name, hp
	}
}

func main() {
	// 调用函数
	gen := Gen("天使彦")
	// 通过匿名函数和闭包，返回天使彦的名字和血量
	name, hp := gen()
	fmt.Println(name, hp)
}

```

可以看出闭包还具有一定的封装性，血量是无法从外部修改的，与面向对象的封装类似，限制外部内内部的访问权限。

闭包很灵活，记住一句话，`闭包 = 函数 + 引用环境`。

### defer语句

这是go语言独有的特性，延时执行语句。

延时语句会在所在函数结束时进行。函数结束可以是正常返回时，也可以是发生宕机时。

由于`defer`语句延迟调用的特性，所以`defer`语句能非常方便的处理资源释放问题。比如：资源清理、文件关闭、解锁及记录时间等。

类似于栈，先进后出。先defer的语句最后执行，后defer的语句最先执行。

```go
package main

import "fmt"

func main() {
	fmt.Println("----start----")
	defer fmt.Println(1)
	defer fmt.Println(2)
	defer fmt.Println(3)
	fmt.Println("----end----")
}

```

执行结果：

```go
----start----
----end----
3
2
1
```

### 宕机(panic)和宕机恢复(recover)

- `panic`：终止程序运行
- `recover`：防止程序崩溃

go语言目前还没有异常机制，但使用`panic/recover`模式来处理错误，`panic`可以在任何地方引发，但`recover`只在`defer`调用的函数中有效。

可以手动触发宕机。让程序崩溃。开发者能及时的发现错误，同时减少可能的损失。

#### 示例1：手动触发宕机

```go
package main

import "fmt"

func testPanic()  {
	fmt.Println("start")
	panic("宕机")
	fmt.Println("end")
}

func main() {
	testPanic()
}

```

执行结果：

```go
start
panic: 宕机

goroutine 1 [running]:
main.testPanic()
	D:/Study/Go_Study/src/golang/study/day04、函数/test.go:7 +0x9d
main.main()
	D:/Study/Go_Study/src/golang/study/day04、函数/test.go:12 +0x27

Process finished with exit code 2
```

#### 示例2：宕机后执行defer语句

```go
package main

import "fmt"

func testPanic()  {
	fmt.Println("start")
	// 在宕机时触发延迟执行语句
	defer fmt.Println("宕机要做的第二件事")
	defer fmt.Println("宕机要做的第一件事")
	panic("宕机")
}

func main() {
	testPanic()
}

```

执行结果：

```go
start
宕机要做的第一件事
宕机要做的第二件事
panic: 宕机

goroutine 1 [running]:
main.testPanic()
	D:/Study/Go_Study/src/golang/study/day04、函数/test.go:10 +0x151
main.main()
	D:/Study/Go_Study/src/golang/study/day04、函数/test.go:14 +0x27

Process finished with exit code 2
```

#### 示例3：宕机恢复

无论是代码运行错误，抛出的宕机错误，还是主动触发的宕机错误，都可以配合defer和recover实现错误捕捉和恢复，让代码崩溃后继续运行。

```go
package main

import "fmt"

func testRecover() {

	// 延时语句捕捉宕机
	defer func() {
		fmt.Println("开始捕捉")
		err := recover()
		fmt.Println("捕捉成功，错误为：", err)
	}()

	panic("发生宕机了，之后执行defer语句，defer语句执行匿名函数，函数内部recover()捕捉错误，继续执行")

}

func main() {
	testRecover()
}

```

执行结果：

```go
开始捕捉
捕捉成功，错误为： 发生宕机了，之后执行defer语句，defer语句执行匿名函数，函数内部recover()捕捉错误，继续执行
```

注意：

- `recover()`必须搭配`defer`使用。
- `defer`一定要在可能引发`panic`的语句之前定义。