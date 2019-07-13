package main

import (
	"fmt"
	"sync"
	"time"
)

// 声明结构体 实现优雅的等待
var wg sync.WaitGroup

// 启动goroutine

func hello() {
	fmt.Println("hello go1")
	time.Sleep(5 * time.Second)
	fmt.Println("hello go2")

	// 标记该任务已完成 最好用defer语句，保证goroutine出错或异常计数器减1。
	defer wg.Done()
}

func main() {
	defer fmt.Println("main stop")

	// 标记一个任务
	wg.Add(1)
	go hello()

	fmt.Println("hello main ")
	// 创建一个goroutine需要时间，等一下他
	// goroutine等待的时间太长，不等他了。
	//time.Sleep(time.Second)

	// 阻塞，一直等待所有的goroutine结束
	wg.Wait()

}
