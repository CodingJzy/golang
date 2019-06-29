package main

import (
	"fmt"
)

func main() {
	n1 := 100
	n2 := 200
	fmt.Println("相加", n1+n2)
	fmt.Println("相减", n1-n2)
	fmt.Println("相乘", n1*n2)
	fmt.Println("相除", n1/n2)
	n1++
	fmt.Println("自增", n1)
	n2--
	fmt.Println("自减", n2)
}
