package main

import "fmt"

// channel

func main() {

	// 定义一个ch1变量，是一个channel类型，这个channel内部传递的数据是int类型
	// channel是引用类型
	var ch1 chan int
	var ch2 chan string

	fmt.Println("ch1：", ch1)
	fmt.Println("ch2：", ch2)

	ch3 := make(chan string)
	fmt.Println("ch3：", ch3)

	// 通道的操作：发送、接受、关闭
	// 发送和接收  <-
	ch3 <- "1"     // 把10发送的通道中
	recv1 := <-ch3 // 从通道中取出10，保存到变量中
	//
	fmt.Println("recv1：", recv1)
	//ch3 <- "2"
	//recv2 := <- ch3
	//fmt.Println("recv2：", recv2)

	// 关闭通道
	//close(ch3)
	// 1、关闭之后可以再取值，只是为对应类型的nli值而已
	//fmt.Println(<-ch3)
	// 2、关闭之后不可以再次发送值，会panic
	// 3、关闭之后不能再关闭
}
