package main

import "fmt"

func recv(ch chan bool) {
	res := <-ch
	fmt.Println(res)
}

func main() {
	ch := make(chan bool, 1)
	ch <- false
	fmt.Println("长度：", len(ch), " 容量：", cap(ch))
	go recv(ch)
	ch <- true
	fmt.Println("main stop")
}
