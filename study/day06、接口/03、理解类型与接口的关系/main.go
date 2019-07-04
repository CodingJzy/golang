package main

import (
	"fmt"
)

// 定义一个Write接口
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
	w.Writer(nil)
}

// 定义一个函数，负责关闭
func useCloser(c Closer) {
	// 执行c接口的关闭方法
	c.Closer()
}

// 定义一个服务接口，有服务开启和日志输出的方法
type Service interface {
	Start()     // 启动服务
	Log(string) // 日志输出
}

// 定义一个游戏服务结构体
type GameService struct {
	Looger // 内嵌logger结构体
}

// 为游戏结构体实现游戏服务的启动method
func (g *GameService) Start() {
	fmt.Println("游戏服务启动成功")
}

// 定义一个日志器结构体
type Looger struct {
}

// 为日志服务器结构体实现日志输出的method
func (l *Looger) Log(s string) {
	fmt.Println(s)
}

func main() {
	fmt.Println("理解类型与接口的关系")

	// 类型和接口之间有一对多和多对一的关系。

	// 一个类型可以实现多个接口
	// 实例化socker结构体
	s := new(Socket)
	useWriter(s)
	useCloser(s)
	// 可以看出两个函数完全独立。完全不知道对象的存在，也不知道使用自己的接口是socket使用的

	// 多个类型可以实现相同的接口
	var ser Service = new(GameService)
	ser.Start()
	ser.Log("日志输出")
}
