package main

import (
	"fmt"
	"os"
)

func main() {

	// os.args 是一个字符串切片类型
	// 运行：go run args.go arg1
	// 结果：[C:\Users\JIANGZ~1\AppData\Local\Temp\go-build591088138\b001\exe\args.exe arg1]
	fmt.Printf("%v\n", os.Args)

	// 遍历每个命令行参数
	for index, arg := range os.Args {
		fmt.Printf("第%d个 参数值：%v\n", index, arg)
	}
}
