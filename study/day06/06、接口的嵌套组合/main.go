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
	fmt.Println("接口的嵌套组合---将多个接口放在一个接口内")

	// 不仅结构体与结构体之间可以嵌套，接口与接口之间也可以通过嵌套创造出新的接口。
	// 接口与接口嵌套组合成了新的接口，只要接口的所有方法被实现，则这个接口中的所有嵌套接口的方法均可以被调用。

	// 实现人的结构体的技能类型
	var s Skill = new(Person)
	s.Say("Life Is Short Let's Go")
	s.Run(9999)

	// 实现哑巴(人)结构体的技能类型，无法说话，只能跑步
	var yaba Run = new(Person)
	yaba.Run(100000)

}
