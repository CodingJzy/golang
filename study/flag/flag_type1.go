package main

import (
	"flag"
	"fmt"
)

func main() {
	name := flag.String("name", "江子牙", "姓名")
	age := flag.Int("age", 23, "年龄")
	isLogin := flag.Bool("isLogin", true, "是都登录")

	// 此时三个变量都是对应类型的指针, 可以用*取值
	fmt.Printf("%p %v \n", name, *name)
	fmt.Printf("%p %v \n", age, *age)
	fmt.Printf("%p %v \n", isLogin, *isLogin)
}
