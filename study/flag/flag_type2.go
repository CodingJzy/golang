package main

import (
	"flag"
	"fmt"
)

func main() {

	var (
		name    string
		age     int
		isLogin bool
	)

	flag.StringVar(&name, "name", "江子牙", "姓名")
	flag.IntVar(&age, "age", 23, "年龄")
	flag.BoolVar(&isLogin, "isLogin", false, "是否登录")

	// 此时三个变量均有有默认值
	fmt.Printf("%#v\n", name)
	fmt.Printf("%#v\n", age)
	fmt.Printf("%#v\n", isLogin)

}
