## Go语言标准库flag

go语言内置的`flag`的包实现了命令行参数的解析，`flag`包使得开发命令行工具更为简单。

### os.Args

简单的获取命令行参数，可以通过`os.Args`获取。

```go
package main

import (
	"fmt"
	"os"
)

func main() {

	// os.args 是一个字符串切片类型
	fmt.Printf("%v\n", os.Args)

	// 遍历每个命令行参数
	for index, arg := range os.Args {
		fmt.Printf("第%d个 参数值：%v\n", index, arg)
	}
}

```

运行：

```linux
go run args.go arg1 arg2 arg3
```

执行结果：

```text
[C:\Users\JIANGZ~1\AppData\Local\Temp\go-build073370358\b001\exe\args.exe arg1 arg2 arg3]
第0个 参数值：C:\Users\JIANGZ~1\AppData\Local\Temp\go-build073370358\b001\exe\args.exe
第1个 参数值：arg1
第2个 参数值：arg2
第3个 参数值：arg3
```

### 基本使用

#### 导入

```go
import "flag"
```

#### 参数类型

支持的命令行参数有`bool`、`int`、`int64`、`uint`、`uint64`、`float`、`float64`、`string`等

#### 定义命令行参数

##### flag.Type()

格式：`flag.Type(flag名, 默认值, 帮助信息)`

比如说我们要定义姓名、年龄、婚否三个命令行参数。

```go
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

```

运行和输出：

```linux
>go run flag_type1.go
0xc000046200 江子牙
0xc00005a058 23
0xc00005a070 true
```

##### flag.TypeVar()

格式：`flag.TypeVar(Type指针, flag名, 默认值, 帮助信息)`

比如说我们要定义姓名、年龄、婚否三个命令行参数。

```go
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
```

运行和输出：

```go
>go run flag_type2.go
"江子牙"
23
false
```

#### 解析命令行参数

通过两种方法定义好命令行flag参数后，需要通过flag.Parse()来对命令进行参数解析。  

支持的格式有以下几种：

- `-flag xxx`：使用空格，1个符号
- `--flag xxx`：使用空格，2个符号
- `-flag=xxx`：使用等号，一个符号
- `--flag=xxx`：使用等号，一个符号

其中布尔类型的必须使用等号的方式指定。

另外还有一些其他方法：

- `flag.NFlag()`：返回命令行参数个数
- `flag.Args()`：返回命令行参数之后的其他参数
- `flag.NArg()`：返回命令行参数之后的其他参数个数

完整示例：

```go
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

```

运行和输出：

```text
>go run  flag_parse.go --name="江子牙" --age 21 -isLogin=false hehe haha
江子牙 21 false
3
[hehe haha]
2
```











