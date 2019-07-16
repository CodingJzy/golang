## go语言的接口

### 简介

- 接口是双方规定的一种合作协议，接口实现者不需要关心接口会被怎样使用，调用者不需要关心接口的实现细节。接口是一种类型。也是一种抽象结构。不会暴露所含数据的格式、类型及结构。比如只要一台洗衣机有洗衣服的功能和甩干的功能，我们就称为洗衣机，不关心属性(数据)，只关心行为(方法)。
- 接口和其他动态语言的鸭子模型有密切关系。比如说`python`、`javascript`。任何类型，只要实现了该接口中的方法集，那么就属于这个类型。
- 每个接口由数个方法组成。

### 接口的定义

格式：

```go
type 接口类型名 interface {
    方法名1( 参数列表1 ) 返回值列表1
    方法名2( 参数列表2 ) 返回值列表2
    …
}
```

- 接口类型名：使用`type`将接口定义为自定义的类型名
- 方法名：当方法名首字母是大写时，且这个接口类型名首字母也是大写时，这个方法可以被接口所在的包之外的代码访问。
- 参数列表、返回值列表：参数列表和返回值列表的参数变量名可以被忽略

示例：

```go
// 可以吃
type Eat interface {
    eat()
}
```

看到这个接口时，只知道他有一个`eat()`方法。并不知道谁实现了这些接口，也不知道是怎么实现的这些这个方法的。

### 实现接口的条件

- 接口的方法与实现接口的类型方法格式一致
- 接口的方法与实现接口的类型方法数目保持一致。即接口中的所有方法均被实现。

```go
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

// 为小鸡结构体实现GaGa
func (c Chicken) GaGa() {
	fmt.Println("我是小鸡，但我也会嘎嘎叫")
}

// 为小鸡结构体实现YouYong
func (c Chicken) YouYong() {
	fmt.Println("我是小鸡，但我也会游泳")
}

func main() {
    
	// 实例化一个小鸡
	c := new(Chicken)

	// 初始化一个小鸡小鸭合体类型
	var duckChicken Duck = c

	// 小鸡小鸭开始嘎嘎
	duckChicken.GaGa()
	// 小鸡小鸭开始游泳
	duckChicken.YouYong()

}
```

可以看出来，`duckChicken`可以直接调用`GaGa`和`YouYong`。他并不知道这两个方法内部怎么实现的。

### 值类型接收者和指针型接收者实现接口

```go
package main

import (
	"fmt"
)

// 定义一个汪汪叫的接口
type WangWager interface {
	Wang()
}

// 定义一个狗结构体
type Dog struct {
	name string
}

func (d Dog) Wang()  {
	fmt.Println("汪汪叫")
}

func main() {
	d1 := Dog{"小狗1"}
	d2 := &Dog{
		"小狗2",
	}
	var w1 WangWager = d1
	w1.Wang()

	var w2 WangWager = d2
	w2.Wang()

}
```

可以发现，使用值类型实现接口后，不管dog结构体实例化是指针型还是值类型。都可以赋值给该接口变量。因为go语言内部有对指针类型变量求值的语法糖。

但是指针型实现接口后，只能指针型变量赋值给接口变量。

### 类型与接口的关系

类型与接口有一对多和多对多的关系。

#### 一(类型)对多(接口)：

```go
package main

import (
	"fmt"
)

// 定义一个Writer接口
type Writer interface {
	Writer(p []byte) (n int, err error)
}

// 定义一个Closer接口
type Closer interface {
	Closer() error
}

// 定义一个socket结构体类型
type Socket struct {
}

// 为socket实现一个Writer()方法
func (s *Socket) Writer(p []byte) (n int, err error) {
	fmt.Println("开始写入")
	return n, err
}

// 为socket实现一个Closer()方法
func (s *Socket) Closer() (err error) {
	fmt.Println("开始关闭")
	return err
}

// 定义一个函数，负责写
func useWriter(w Writer) {
	// 执行w接口的写方法
    _, _ := w.Writer(nil)
}
// 定义一个函数，负责关闭
func useCloser(c Closer) {
	// 执行c接口的关闭方法
    _ := c.Closer()
}

func main() {
	fmt.Println("理解类型与接口的关系")

	// 类型和接口之间有一对多和多对一的关系。

	// 一个类型可以实现多个接口
	// 实例化socker结构体
	s := new(Socket)
	useWriter(s)
	useCloser(s)
}
```

可以看出两个函数完全独立，完全不知道对方的存在，也不知道使用自己的接口是socket使用的

#### 多(类型)对一(接口)

接口的方法不一定要一个结构体类型完全实现，接口的方法可以通过结构体嵌入实现。

```go
package main

import (
	"fmt"
)

// 定义一个服务接口，实现了服务开启和日志输出的方法
type Service interface {
	Start()     // 启动服务
	Log(string) // 日志输出
}

// 定义一个游戏服务结构体
type GameService struct {
	Logger // 内嵌logger结构体
}

// 为游戏结构体实现游戏服务的启动method
func (g *GameService) Start() {
	fmt.Println("游戏服务启动成功")
}

// 定义一个日志器结构体
type Logger struct {
}

// 为日志服务器结构体实现日志输出的method
func (l *Logger) Log(s string) {
	fmt.Println(s)
}

func main() {

	// 多个类型可以实现相同的接口
	var ser Service = new(GameService)
	ser.Start()
	ser.Log("日志输出")
}
```

可以看出，服务接口下一个服务启动功能和日志输出功能，但是游戏服务类型并没有实现日志输出功能，而是间接通过内嵌日志类型来实现，日志类型实现了日志输出功能，所有游戏服务类型实现的接口可以直接使用日志输出功能。、

### 接口的嵌套组合

不仅类型与类型之间可以嵌套，接口与接口之间也可以嵌套。

```go
package main

import (
	"fmt"
)

// 定义一个Say接口
type Say interface {
	Say(s string)
}

// 定义一个Run接口
type Run interface {
	Run(n int)
}

// 定义一个Skill接口
type Skill interface {
	// 嵌套了两个接口
	Say
	Run
}

// 定义一个Person结构体
type Person struct {
}

// 为Person结构体实现Say方法
func (p *Person) Say(s string) {
	fmt.Println("说话：", s)
}

// 为Person结构体实现Run方法
func (p *Person) Run(n int) {
	fmt.Println("步数：", n)
}

func main() {
	// 实现人的结构体的技能类型
	var s Skill = new(Person)
	s.Say("Life Is Short Let's Go")
	s.Run(9999)

	// 实现哑巴(人)结构体的技能类型，无法说话，只能跑步
	var yaBa Run = new(Person)
	yaBa.Run(100000)

}

```

### 空接口

空接口是接口类型的特殊形式，空接口没有任何方法。因此任何类型都无须实现，从实现的角度来看。任何值都满足这个接口的需求。因此空接口类型可以保存任何值，也可以从中取值。

#### 保存值

```go
package main

import (
	"fmt"
)

func main() {
	// 将各种数据类型的值保存到空接口
	var any interface{}

	any = 1
	fmt.Println(any)
	any = 0.99
	fmt.Println(any)
	any = "Hello Gp"
	fmt.Println(any)
	any = true
	fmt.Println(any)
	any = []string{}
	fmt.Println(any)
	any = [...]string{}
	fmt.Println(any)
	any = map[string]int{}
	fmt.Println(any)
}

```



#### 空接口的应用

##### 空接口作为函数的参数

```go
package main

import "fmt"

func show(v interface{})  {
	fmt.Println(v)
}

func main() {
	show("江子牙")
	show(520)
	show(true)
}
```

##### 空接口作为map的value

```go
package main

import "fmt"

func main() {
	a := map[string]interface{}{
		"name":    "江子牙",
		"age":     21,
		"isLogin": true,
	}
	fmt.Println(a)
}
```

### 接口和类型之间的转换

go 语言中使用接口断言(type assertions) 将接口转换成另一外一个接口，也可以将接口转另外的类型。使用非常频繁。

#### 类型断言

格式：

```go
t := i.(T)
```

- i：代表接口变量
- T：代表转换的目标类型
- t：转换后的变量

如果i没有完全实现T接口的方法，这个语句会触发宕机，触发宕机不是很友好，因为有另一种保守写法。

```go
t, ok := i.(T)
```

这种写法的好处就是如果发送接口未实现时，将会返回一个布尔值`false`，即`ok`的值，而且t的类型为0。正常实现时，`ok`为`true`。

#### 接口转化为其他接口

例子：

```go
package main

import (
	"fmt"
)

// 定义飞行动物接口
type Flyer interface {
	Fly(s string)
}

// 定义爬行动物接口
type Walker interface {
	Walk(s string)
}

// 定义小鸟类结构体
type Bird struct {
}

// 定义小猪类结构体
type Pig struct {
}

// 为小鸟类实现飞的技能和走路的技能
func (b *Bird) Fly(s string) {
	fmt.Println("小鸟飞行：", s)
}
func (b *Bird) Walk(s string) {
	fmt.Println("小鸟走路：", s)
}

// 为小猪实现走路的技能
func (p *Pig) Walk(s string) {
	fmt.Println("死猪跑不动吗：", s)
}

func main() {
	// 先创建一个字典来存放接口的信息。
	animals := map[string]interface{}{
		"bird": new(Bird),
		"pig":  new(Pig),
	}
	fmt.Println(animals)
	// 循环map
	for name, obj := range animals {
		// 类型断言：判断obj是爬行动物还是飞行动物
		f, isFly := obj.(Flyer)
		w, isWalk := obj.(Walker)
		fmt.Printf("name：%s\tisFlyer：%v\tisWalker：%v\n", name, isFly, isWalk)

		// 类型判断
		if isFly {
			f.Fly("100米")
		}
		if isWalk {
			w.Walk("1000步")
		}
	}
}
```

#### 接口转化为类型

```go
package main

import (
	"fmt"
)

// 定义爬行动物接口
type Walker interface {
	Walk(s string)
}

// 定义小猪类结构体
type Pig struct {
}

// 为小猪实现走路的技能
func (p *Pig) Walk(s string) {
	fmt.Println("死猪跑不动吗：", s)
}

func main() {
	// 将接口转为其他类型：可以实现将接口转为普通的指针类型
	// 实例化出一个小猪结构体
	p1 := new(Pig)
	// 声明一个类型为小猪爬行类w接口
	var w Walker = p1
	fmt.Println(w)
	fmt.Printf("%T\n", w)
	fmt.Printf("%T\n", p1)
	// 将w接口转为*Pig类型
	p2 := w.(*Pig)
	fmt.Printf("p1 = %p\np2 = %p", p1, p2)
}
```

执行结果：

```go
&{}
*main.Pig
*main.Pig
p1 = 0x5861b0
p2 = 0x5861b0
```

### 判断接口中变量的类型

#### 判断基本类型

```go
package main

import (
	"fmt"
)

func printType(i interface{}) {
	switch i.(type) {
	case int:
		fmt.Println("int类型")
	case string:
		fmt.Println("string类型")
	case bool:
		fmt.Println("bool类型")
	default:
		fmt.Println("不知名类型")
	}
}

func main() {
	// 使用类型分支判断基本类型
	printType("str")
	printType(1)
	printType(true)
	printType([]string{})
}
```

执行结果：

```go
string类型
int类型
bool类型
不知名类型
```

#### 判断接口类型

```go
package main

import (
	"fmt"
)

// 刷脸的接口
type useFace interface {
	Face(string)
}

// 刷人民币值为100的接口
type useOneHundred interface {
	OneHundred(string)
}

// 支付宝方式结构体
type Alipy struct {
}

// 现金方式方式结构体
type Cash struct {
}

// 为支付宝提供人脸支付的方法
func (a Alipy) Face(s string) {
	fmt.Println(s)
}

// 为现金支付提供支付100元的方法
func (c Cash) OneHundred(s string) {
	fmt.Println(s)
}

func printPay(payMethod interface{}) {
	var pay = payMethod
	// 判断接口类型
	switch payMethod.(type) {
	case useFace:
		face := pay.(*Alipy)
		face.Face("刷脸")
	case useOneHundred:
		cash := pay.(*Cash)
		cash.OneHundred("支付100毛爷爷")
	}
}

func main() {

	//使用类型分支判断接口类型
	printPay(new(Alipy))
	printPay(new(Cash))
}
```

执行结果：

```go
刷脸
支付100毛爷爷
```



