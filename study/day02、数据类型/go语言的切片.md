## go语言的切片

因为数组的长度固定不变并且长度也是数组类型的一部分，所以它有很多局限性。这时有了一种新的数据类型。在go语言中叫做切片(`slice`)

### 简介

切片是一个拥有相同数据类型元素的可变长度的序列。它是基于数组类型做的一层封装，它非常灵活，支持自动扩容。

切片是一个引用类型，它的内部结构包含`地址`、`长度`和`容量`。切片一般用于快速地操作一块数据集合。

### 定义

基本语法：

```go
var 切片变量名 []切片元素类型
```

### 初始化

切片必须初始化或者`append`之后才能使用

```go
package main

import (
	"fmt"
)

func main() {
	slice := []float64{1.1, 2.2, 3.3}
	
	// 类型：[]float64	值：[]float64{1.1, 2.2, 3.3}
	fmt.Printf("类型：%T\t值：%#v\n", slice, slice)
	// [1.1 2.2 3.3]
	fmt.Println(slice)
}

```

### 生成切片

#### 基于数组

切片的底层就是数组，可以通过数组得到切片。

```go
package main

import (
	"fmt"
)

func main() {
	array1 := [5]float64{1.1, 2.2, 3.3, 4.4, 5.5}
	slice1 := array1[1:5]

	// 得到切片：[2.2 3.3 4.4 5.5]
	fmt.Println(slice1)
	// 类型：[]float64
	fmt.Printf("%T",slice1)
}

```

这里的切片和python的列表类似，顾头不顾尾，还支持省略头和尾

#### make函数

上面都是基于数组来创建的切片，如果需要动态的创建一个切片，我们就需要使用内置的`make()`函数，格式如下：

```go
make([]T, size, cap)
```

- T：切片元素类型
- size：切片中元素数量
- cap：切片的容量，设置后，不影响size，只是提前分配空间，降低多次分配空间造成的性能问题

```go
package main

import "fmt"

func main() {
	slice := make([]int, 2, 10)
	// [0 0]
	fmt.Println(slice)
	// 2
	fmt.Println(len(slice))
	// 10
	fmt.Println(cap(slice))
}

```

上面代码中`slice`的内部存储空间已经分配了10个，但实际上只用了2个。 容量并不会影响当前元素的个数，所以`len(slice)`返回2，`cap(slice)`则返回该切片的容量。

### 切片的长度和容量

切片拥有自己的长度和容量，我们可以通过使用内置的`len()`函数求长度，使用内置的`cap()`函数求切片的容量。

切片的底层是数组，长度指当前切片内的元素数目。容量指底层数组的容下的最大元素数

```go
package main

import "fmt"

func main() {
	array1 := [...]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	slice1 := array1[1:4]
	
	// 3
	fmt.Println("长度：", len(slice1))
	// 9
	fmt.Println("容量：", cap(slice1))
}

```

### 切片的常用操作

#### 添加

##### 单个元素

当切片容量不够时，可以通过`append()`扩容，最简单的扩容规律按两倍数扩充：1,2,4...

扩容规律还有其他情况，这里暂不讨论。

```go
package main

import "fmt"

func main() {
	var slice6 []int
	for i := 0; i < 20; i++ {
		slice6 = append(slice6, i)
		// 打印扩容情况
		fmt.Printf("长度：%d     cap：%d      指针：%p\n", len(slice6), cap(slice6), slice6)
	}
}

```

执行结果：

```go
长度：1     cap：1      指针：0xc000056080
长度：2     cap：2      指针：0xc0000560c0
长度：3     cap：4      指针：0xc0000540c0
长度：4     cap：4      指针：0xc0000540c0
长度：5     cap：8      指针：0xc0000680c0
长度：6     cap：8      指针：0xc0000680c0
长度：7     cap：8      指针：0xc0000680c0
长度：8     cap：8      指针：0xc0000680c0
长度：9     cap：16      指针：0xc000086080
长度：10     cap：16      指针：0xc000086080
长度：11     cap：16      指针：0xc000086080
长度：12     cap：16      指针：0xc000086080
长度：13     cap：16      指针：0xc000086080
长度：14     cap：16      指针：0xc000086080
长度：15     cap：16      指针：0xc000086080
长度：16     cap：16      指针：0xc000086080
长度：17     cap：32      指针：0xc00003c100
长度：18     cap：32      指针：0xc00003c100
长度：19     cap：32      指针：0xc00003c100
长度：20     cap：32      指针：0xc00003c100
```

##### 多个元素

```go
package main

import "fmt"

func main() {
	var slice6 []int
	
	slice6 = append(slice6, 1, 2, 3, 4)
	fmt.Println(slice6)
}

```

##### 添加切片

```go
package main

import "fmt"

func main() {
	var slice6 []int

	slice6 = append(slice6, 1, 2, 3, 4)
	// [1 2 3 4]
	fmt.Println(slice6)

	slice6 = append(slice6,[]int{5, 6, 7}...)
	// [1 2 3 4 5 6 7]
	fmt.Println(slice6)
}

```

#### 删除

go语言中并没有提供切片删除元素的方法，但是我们可以使用切片本身的特性来删除元素。

本质：以被删除的元素为分界点，将前后两个部分的内存重新连接起来。

```go
package main

import "fmt"

func main() {
	slice := []string{"a", "b", "c", "d"}
	// 删除c
	slice = append(slice[:2],slice[3:]...)
	fmt.Println(slice)
}

```

