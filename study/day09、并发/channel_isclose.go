package main

import "fmt"

func send(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}

func main() {
	ch := make(chan int, 100)

	go send(ch)

	// 利用的for循环接收
	//for  {
	//	ret, ok := <- ch
	//	if !ok {
	//		break
	//	}
	//	fmt.Println(ret)
	//}

	// 利用for range接收
	for ret := range ch {
		fmt.Println(ret)
	}
}
