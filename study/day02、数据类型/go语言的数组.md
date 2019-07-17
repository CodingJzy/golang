## go语言的数组

[TOC]

### 简介

数组是同一种数据类型元素的集合，固定大小的连续空间，和Python这种动态语言的列表有点区别。

数组长度在声明时已经确定，或者编译器确定了，后期可以修改数组成员，但是不能修改大小。

### 定义

格式：

```go
var 数组变量名 [元素数量]类型
```

声明一个变量名为`team`，长度为4，元素类型为`string`的数组：

```go
var team [4]string
```

注意：

数组的长度必须是常量。长度是数组类型的一部分，比如`var team1 [2]string`和`team`类型是不一样的。

### 初始化

#### 空数组

```go
var kong = []string{}
```

#### 方法1

使用初始化列表来设置数组元素的值：

```go
package main

import (
	"fmt"
)

func main() {
	// 声明并初始化
	var team  = [4]string{"Linux", "Python","Java","Go"}
	
	// [Linux Python Java Go]
	fmt.Println(team)
}

```

#### 方法2

让编译器确定数组的长度：

```
package main

import (
	"fmt"
)

func main() {
	// 让编译器确定数组的长度
	var team = [...]string{"Linux", "Python","Java","Go"}

	// [Linux Python Java Go]
	fmt.Println(team)
}

```

#### 方法3

还可以使用指定索引值的方式来初始化数组：

```go
package main

import (
	"fmt"
)

func main() {
	// 指定索引值的方式来初始化数
	var team = [...]string{1: "Linux", 2: "Python",4: "Java",3: "Go"}

	// [Linux Python Java Go]
	fmt.Println(team)
}

```

### 数组遍历

#### for range

```go
package main

import (
	"fmt"
)

func main() {
	var team = [...]string{0: "Linux", 1: "Python",3: "Java",2: "Go"}

	// 数组遍历：for range
	for index, value := range team {
		fmt.Printf("索引：%d\t索引值：%v\n", index, value)
	}
}

```

#### for

```go
package main

import "fmt"

func main() {
	var team = [...]string{0: "Linux", 1: "Python", 3: "Java", 2: "Go"}

	for i := 0; i < len(team); i++ {
		fmt.Printf("索引：%d\t索引值：%v\n", i, team[i])
	}
}

```

两种结果一样：

```go
索引：0	索引值：Linux
索引：1	索引值：Python
索引：2	索引值：Go
索引：3	索引值：Java
```

