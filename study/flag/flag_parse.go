package main

import (
	"flag"
	"fmt"
)

func main() {
	// 定义命令行参数
	name := flag.String("name", "江子牙", "姓名")
	age := flag.Int("age", 23, "年龄")
	isLogin := flag.Bool("isLogin", true, "是都登录")

	// 解析命令行参数
	flag.Parse()

	// 输出
	fmt.Println(*name, *age, *isLogin)

	// 返回命令行参数个数
	fmt.Println(flag.NFlag())

	// 返回命令行参数之后的其他参数
	fmt.Println(flag.Args())

	// 返回命令行参数之后的其他参数个数
	fmt.Println(flag.NArg())
}
