package main

import (
	"fmt"
	"sync"
	"time"
)

func read(i int) {
	// 读锁定
	m.RLock()
	// 开始读
	fmt.Println(i, "start read")
	// 读取中
	fmt.Println(i, "reading")
	// 睡2秒
	time.Sleep(2 * time.Second)
	// 读取完毕
	fmt.Println(i, "read complete")
	// 读解锁
	m.RUnlock()

}

func write(i int) {

	// 写入中
	m.Lock()
	// 开始写
	fmt.Println(i, "start write")
	fmt.Println(i, "writing")
	time.Sleep(2 * time.Second)
	// 写完成
	fmt.Println(i, "write complete")
	// 写解锁
	m.Unlock()

}

func testRead() {
	go read(1)
	go read(2)
	go read(3)
}

func testWrite() {
	go write(1)
	go write(2)
	go write(3)
}

var m *sync.RWMutex

func main() {
	/*
		基本遵循两大原则
			1、可以随便读，多个goroutine同时读
			2、写的时候，啥也不能干，不能读也不能写

		RWMutex提供了四个方法：

			func (*RWMutex) Lock       // 写锁定

			func (*RWMutex) Unlock     // 写解锁

			func (*RWMutex) RLock 	    // 读锁定

			func (*RWMutex) RUnlock     // 读解锁
	*/
	m = new(sync.RWMutex)
	// 同时读：可以看出其中某个子线程还没结束，另一个线程已经在读。
	testRead()

	// 同时写：其中一个线程写入完毕之后其余线程才能写
	//testWrite()

	// 让主线程等待子线程的结束
	time.Sleep(6 * time.Second)

}
