# go语言的结构体

## 简介

go语言中没有类的概念，也不支持类的继承等面向对象的概念, 但在结构体的内嵌配合接口比面向对象更具更高的扩展性和灵活性。

Go语言中的基础数据类型可以表示一些事物的基本属性，但是当我们想表达一个事物的全部或部分属性时，这时候再用单一的基本数据类型明显就无法满足需求了，Go语言提供了一种自定义数据类型，可以封装多个基本数据类型，这种数据类型叫结构体`struct`，通过结构体来实现面向对象。

## 定义

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

## 结构体的实例化

结构体的定义只是一种内存布局的描述，此时并不会分配内存，只有实例化的时候才会真正分配内存。

在定义结构体并实例化才能使用结构体字段，有多重实例方式。

### 基本实例化

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

### 创建指针类型的结构体

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

### 取结构体的地址实例化

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

## 结构体的初始化

### 键值对初始化

类似于Python的关键字传参

```go
package main

import (
	"fmt"
)

type Person struct {
	name string
	age int
	isLogin bool
} 

func main() {

	p1 := Person{
		name:"江子牙",
		age:23,
		isLogin:true,
	}
	
	// {江子牙 23 true}
	fmt.Println(p1)
	
}
```

### 值的列表初始化

类似于Python的位置参数

```go
package main

import (
	"fmt"
)

type Person struct {
	name string
	age int
	isLogin bool
}

func main() {

	p2 := Person{
		"江子牙", 23, true,
	}
	// {江子牙 23 true}
	fmt.Println(p2)
}
```

使用此方式初始化时，值得注意的是：

- 必须初始化结构体的所有字段
- 初始化顺序必须和结构体声明时保持一致
- 两种方式的初始化不能混用

### 匿名结构体的初始化

顾名思义，匿名结构体就是没有名字的结构体，无需用`type`关键字声明就可以直接使用。

```go
package main

import (
	"fmt"
)

func main() {
	s1 := struct {
		name string
		age  int
	}{
		"江子牙",
		23,
	}
    
	// 江子牙
	fmt.Println(s1.name)
	// 23
	fmt.Println(s1.age)
	// {江子牙 23}
	fmt.Println(s1)
}
```

## 构造函数

结构体没有构造函数，我们可以自己实现。因为结构体是值类型，如果结构体较复杂的话，值拷贝的性能开销会比较大，所以我们的构造函数返回结构体指针类型。

```go
package main

import "fmt"

type Cat struct {
	Name  string
	Color string
}

func NewCat(name, color string) *Cat {
	return &Cat{
		Name:  name,
		Color: color,
	}
}

func main() {
	
    // 调用构造函数实例化一个小豹子
	cat := NewCat("小豹子", "豹纹")
	// &{小豹子 豹纹}
	fmt.Println(cat)
	// 小豹子
	fmt.Println(cat.Name)
	// 豹纹
	fmt.Println(cat.Color)
}

```

## 方法Method

go语言中的方法method是一种作用于特定类型变量的函数，叫做`receiver`，可以理解为其他语言的`this`和`self`。

格式如下：

```go
func (接收器变量 接收器类型) 方法名 (参数列表) (返回值列表) {
    函数体
}
```

- 接收器变量：官方建议，该接收器变量最好和接收器类型的第一个字母保持一致。如 Cat c 、Bag b 、Dog d。
- 接收者类型：可以是指针类型和非指针类型。
- 方法名、参数列表、返回值：与函数定义相同。

```go
package main

import "fmt"

// 定义一个包的结构体
type Bag struct {
	items []interface{}
}

// 为结构体定义一个Insert方法
func (b *Bag) Insert(thing interface{}) {
	b.items = append(b.items, thing)
}

func main() {

	// 实例化一个包
	bag := new(Bag)
	// 放入一个东西
	bag.Insert("口红")

	// [口红]
	fmt.Println(bag.items)
	// &{[口红]}
	fmt.Println(bag)


	// 放入一个手机
	bag.Insert("手机")
	// [口红 手机]
	fmt.Println(bag.items)
	// &{[口红 手机]}
	fmt.Println(bag)

}
```

## 接收者

### 指针类型的接收者

指针类型的接收器是一个结构体的指针，更接近于面向对象的this or self。

由于指针的特性，调用方法时，修改接收器指针的任意成员变量，在方法调用结束后，已经修改过得都是有效的。  

上个例子的接收器*Bag就是一个指针接收器，所以才能append之后保持修改。

### 值类型的接收者

当方法作用与非指针接收器时，go语言内部会在代码运行的时候将接收器的数据复制一份，在非指针接收器的方法可以获取，也可以修改，只是修改无效而已。

```go
package main

import "fmt"

type Bag struct {
	name string
}

func (b *Bag) SetName(name string) {
	b.name = name
}

func (b Bag) SetName1(name string) {
	fmt.Println(name, "修改不生效")
	b.name = name
}

func main() {

	// 实例化一个包
	bag := new(Bag)

	// 设置
	bag.SetName("指针类型的接收者")
	// 生效了：指针类型的接收者
	fmt.Println(bag.name)
	// 再设置
	bag.SetName1("值类型的接收者")
	// 不生效：指针类型的接收者
	fmt.Println(bag.name)

}
```

### 总结

- 需要修改接收者中的值

- 接收者是拷贝代价比较大的大对象
- 保证一致性，如果有某个方法使用了指针接收者，那么其他的方法也应该使用指针接收者。

## 为任意类型添加方法

go语言可以对任何一种类型添加方法。给一种类型添加方法就像给结构体添加方法一样。

```go
package main

import "fmt"

// 定义一个MyInt类型
type MyInt int

// 为 MyInt 添加是否为0的方法
func (m MyInt) IsZero() bool {
	return m == 0
}

// 为 MyInt 添加add()方法
func (m MyInt) Add(other int) int {
	return int(m) + other
}

func main() {
	var b MyInt
	// 定义之后为零值
	fmt.Println(b.IsZero())
	
	b = 100
	// 赋值之后不为0
	fmt.Println(b.IsZero())
	// 两数相加：100+150
	fmt.Println(b.Add(150))
}
```

## 结构体的匿名字段

结构体允许字段没有字段名，只有类型。这种字段就叫做匿名字段。

```go
package main

import "fmt"

// 定义一个匿名字段的结构体
type Person struct {
	string
	int
}

func main() {
	p1 := &Person{
		"江子牙",
		23,
	}
	// &{江子牙 23}
	fmt.Println(p1)
}
```

## 嵌套结构体

一个结构体中可以嵌套包含另一个结构体或结构体指针。

比如：

```go
package main

import "fmt"

// 定义一个学生的结构体
type Student struct {
	name string
	age int
	*Class  // 嵌套一个匿名结构体的指针
	School // 嵌套一个匿名结构体
}

// 定义一个班级结构体
type Class struct {
	className string
}

// 定义一个学校结构体
type School struct {
	schoolName string
}

func main() {
	s1 := &Student{
		"江子牙",
		23,
		&Class{
			"高三一班",
		},
		School{"明珠中学"},
	}
	// &{江子牙 23 0xc0000421c0 {明珠中学}}
	fmt.Println(s1)
	// 高三一班
	fmt.Println(s1.className)
	// 明珠中学
	fmt.Println(s1.schoolName)
}
```

学生结构体嵌套了班级和学校两个结构体，获取班级学校和班级名字的时候，会一直往下找。可以通过`.`的方式。

## 结构体的继承

go语言没有面向对象，更没有继承。只是可以实现同样的效果。

```go
package main

import "fmt"

// 定义一个儿子的结构体
type Child struct {
	name string
	*Father // 通过匿名结构体实现继承
}

// 定义一个父亲结构体
type Father struct {
	money int64
}

// 为父亲添加一个生小孩的方法
func (f *Father) Make(name string) {
	fmt.Printf("%s生小孩", name)
}

func main() {
	ch := &Child{
		name:"小明",
		Father:&Father{50000},
	}

	// 小明继承的财产为：50000
	fmt.Printf("%s继承的财产为：%d\n",ch.name, ch.money )

	// 小明生小孩
	ch.Make("小明")
}
```

小明结构体中并没有`money`字段和`Make`方法。这两个都是他父亲结构体中才有的，这样一来便实现了继承属性和方法。

## 结构体字段的可见性

结构体中字段大写开头表示可公开访问，小写表示私有（仅在定义当前结构体的包中可访问）。

## 结构体json序列化

```go
package main

import (
	"encoding/json"
	"fmt"
)

// 学生结构体
type Student struct {
	ID      int
	Gender  string
	Name    string
	IsLogin bool
}

func main() {
	var stu1 = Student{
		ID:      1,
		Gender:  "男",
		Name:    "豪杰",
		IsLogin: false,
	}

	// 序列化：把内存里的数据转换为字符串，以便用于网络传输和数据交换
	if v, err := json.Marshal(stu1); err != nil {
		fmt.Println("JSON格式化出错")
		fmt.Println(err)
	} else {
		// 字节类型的切片
		fmt.Println(v)
		// 转为字符串
		fmt.Println(string(v))
	}

	// 反序列化：把json字符串转换为当前编程语言可用的对象
	str := `{"ID":1,"Gender":"男","Name":"豪杰","IsLogin":false}`
	stu2 := new(Student)
	_ = json.Unmarshal([]byte(str), stu2)
	fmt.Println(stu2)
	fmt.Printf("stu2 类型：%T\tstu2值：%#v", stu2, stu2)
}

```

## 结构体的标签

`Tag`是结构体的元信息，可以在运行的时候通过反射的机制读取出来。 `Tag`在结构体字段的后方定义，由一对反引号包裹起来，具体的格式如下：

```go
`key1:"value1" key2:"value2"`
```

结构体标签由一个或多个键值对组成。键与值使用冒号分隔，值用双引号括起来。键值对之间使用一个空格分隔。 

**注意事项：** 为结构体编写`Tag`时，必须严格遵守键值对的规则。结构体标签的解析代码的容错能力很差，一旦格式写错，编译和运行时都不会提示任何错误，通过反射也无法正确取值。例如不要在key和value之间添加空格。

因为序列化和反序列化大多是与前端进行数据交互。不能给他返回一些大写的键。所以可以通过`tag`来解决。

```go
package main

import (
	"encoding/json"
	"fmt"
)

//Student 学生
type Student struct {
	ID     int    `json:"id"` //通过指定tag实现json序列化该字段时的key
	Gender string //json序列化是默认使用字段名作为key
	name   string //私有不能被json包访问
}

func main() {
	s1 := Student{
		ID:     1,
		Gender: "男",
		name:   "沙河娜扎",
	}
	data, err := json.Marshal(s1)
	if err != nil {
		fmt.Println("json marshal failed!")
	}

	//json str:{"id":1,"Gender":"男"}
	fmt.Printf("json str:%s\n", data)
}

```