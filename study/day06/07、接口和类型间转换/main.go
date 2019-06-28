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
	fmt.Println("在接口和类型间转换")

	// go 语言中使用接口断言(type assertions) 将接口转换成另一外一个接口，也可以将接口转另外的类型。使用非常频繁。

	// 类型断言的格式
	/*
		t := i.(T)
			i：代表接口变量
			T：代表转换的目标类型
			t：转换后的变量
		如果i没有完全实现T接口的方法，这个语句会触发宕机，触发宕机不是很友好，因为有另一种保守写法。

		t, ok := i.(T)
		这种写法的好处就是如果发送接口未实现时，将会返回一个布尔值false，即ok的值，而且t的类型为0。正常实现时，ok为true。
	*/

	// 将接口转换为其他接口
	// 先创建一个字典来存放接口的信息。

	animals := map[string]interface{}{
		"bird": new(Bird),
		"pig":  new(Pig),
	}
	fmt.Println(animals)
	// 循环map
	for name, obj := range animals {
		// 类型断言
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

	// 将接口转为其他类型：可以实心将接口转为普通的指针类型
	// 实例化出一个小猪结构体
	p1 := new(Pig)
	// 声明一个类型为小猪爬行类w接口
	var w Walker = p1
	fmt.Println(w)
	fmt.Printf("%T\n", w)
	fmt.Printf("%T\n", p1)
	p2 := w.(*Pig)
	fmt.Printf("p1 = %p\np2 = %p", p1, p2)

}
