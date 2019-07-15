## go语言的结构体

### 简介

go语言中没有类的概念，也不支持类的继承等面向对象的概念, 但在结构体的内嵌配合接口比面向对象更具更高的扩展性和灵活性。

Go语言中的基础数据类型可以表示一些事物的基本属性，但是当我们想表达一个事物的全部或部分属性时，这时候再用单一的基本数据类型明显就无法满足需求了，Go语言提供了一种自定义数据类型，可以封装多个基本数据类型，这种数据类型叫结构体`struct`，通过结构体来实现面向对象。

### 定义

使用`type`和`struct`关键字来定义结构体，格式如下：

```go
type 类型名 struct {
    字段名 字段类型
    字段名 字段类型
    …
}
```

其中：

- 类型名：表示自定义结构体的名称，在同一个包内不能重复。
- 字段名：表示结构体字段名，结构体中的名字必须唯一。
- 字段类型：表示结构体字段的具体类型。

举个例子，定义一个坐标结构体：

```go
type Point struct {
    X int
    Y int
}
```

同种类型的字段也可以放一行，定义一个颜色的三原色结构体：

```go
type Color struct {
    R, G, B byte
}
```

### 结构体的实例化

结构体的定义只是一种内存布局的描述，此时并不会分配内存，只有实例化的时候才会真正分配内存。

在定义结构体并实例化才能使用结构体字段，有多重实例方式。

#### 基本实例化

结构体本身是一种类型，可以像声明变量一样，以 var 的方式声明结构体即可完成实例化。

```go
var 结构体实例 结构体类型
```

举个例子：

```go
package main

import (
	"fmt"
)

type Person struct {
	name string
	city string
	age  int
}

func main() {
	var p1 Person
	p1.name = "江子牙"
	p1.city = "江西"
	p1.age = 23

	// {江子牙 江西 23}
	fmt.Println(p1)
	
	// 类型：main.Person       值：main.Person{name:"江子牙", city:"江西", age:23}
	fmt.Printf("类型：%T\t值：%#v", p1, p1)
}
```

我们可以通过`.`来对结构体的字段进行赋值和取值操作。

#### 创建指针类型的结构体

我们还可以通过使用`new`关键字对结构体进行实例化，得到的是结构体的地址。 格式如下：

```go
package main

import (
	"fmt"
)

type Person struct {
	name string
	city string
	age  int
}

func main() {
	p1 := new(Person)

	// &{  0} 得到的是一个结构体指针，字段的值对应该字段类型的零值
	fmt.Println(p1)
	// 指针：0xc00006e2d0      类型：*main.Person      值：&main.Person{name:"", city:"", age:0}
	fmt.Printf("指针：%p\t类型：%T\t值：%#v", p1, p1, p1)
}

```

尽管得到的是一个结构体指针，但是go语言中同样可以通过`.`来进行对字段进行操作：

```go
package main

import (
	"fmt"
)

type Person struct {
	name string
	city string
	age  int
}

func main() {
	p1 := new(Person)

	p1.name = "江子牙"
	p1.city = "江西"
	p1.age = 23

	// &{江子牙 江西 23}
	fmt.Println(p1)
	// 指针：0xc00006e2d0      类型：*main.Person      值：&main.Person{name:"江子牙", city:"江西", age:23}
	fmt.Printf("指针：%p\t类型：%T\t值：%#v", p1, p1, p1)
}

```

#### 取结构体的地址实例化

使用`&`对结构体进行取地址操作相当于对结构体进行了一次`new`操作。

```go
package main

import (
	"fmt"
)

type Person struct {
	name string
	city string
	age  int
}

func main() {
	p2 := &Person{}
	p2.name = "江子牙"
	p2.city = "江西"
	p2.age = 23

	// &{江子牙 江西 23}
	fmt.Println(p2)
	// 指针：0xc00006e2d0      类型：*main.Person      值：&main.Person{name:"江子牙", city:"江西", age:23}
	fmt.Printf("指针：%p\t类型：%T\t值：%#v", p2, p2, p2)
}

```



