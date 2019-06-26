package main

import (
	"fmt"
)

func testBreak() {
	// break 语句可以结束for switch select 的代码块，还可以在break后面添加标签，表示退出某个标签对应的代码块。标签要求定义在对应的代码块前

Loop:
	for i := 0; i < 3; i++ {
		for j := 0; j < 5; j++ {
			switch j {
			case 2:
				fmt.Println(2, i, j)
				break Loop
			case 4:
				fmt.Println(4, i, j)
				break Loop
			}
		}
	}
}

func testContinue() {
	// continue 可以退出当前循环，继续下一次外层循环。仅限在for循环内部使用。和break一样。支持标签。

Loop:
	for i := 0; i < 2; i++ {
		for j := 0; j < 5; j++ {
			if j == 2 {
				fmt.Println(j, "退出当前循环， 继续下一次外层循环")
				continue Loop
			} else if j == 3 {
				fmt.Println(j, "退出当前循环，继续下一次外层循环")
				continue Loop
			}
			fmt.Println(j, "哈哈")
		}
	}
}

func main() {

	// 跳出指定循环--- 可以跳出多层循环
	// testBreak()

	// 继续下一次循环
	testContinue()

}
