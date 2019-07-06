package main

import (
	"fmt"
	"os"
	"time"
)

func ShowMain() {
	fmt.Printf(`
-----------学生信息管理系统-------------
1、添加学生
2、修改学生
3、删除学生
4、学生列表
5、退   出
`)
}

func main() {
	var id int64
	manager := NewStudentManager()
	for {
		ShowMain()
		// 用户输入指令
		var input int
		fmt.Print("请输入指令：")
		_, _ = fmt.Scanln(&input)
		fmt.Printf("输入的选项为：%d\n", input)
		switch input {
		case 1:
			fmt.Println("添加学生...")
			stu := InputStudent()
			manager.AddStudent(stu)
		case 2:
			fmt.Println("修改学生")
			manager.ListStudent()
			fmt.Print("请输入要更新的学生ID：")
			_, _ = fmt.Scanln(&id)
			manager.UpdateStudent(id)
		case 3:
			fmt.Println("删除学生...")
			manager.ListStudent()
			fmt.Print("请输入要删除的学生ID：")
			_, _ = fmt.Scanln(&id)
			manager.DeleteStudent(id)
		case 4:
			fmt.Println("学生列表...")
			manager.ListStudent()

		case 5:
			fmt.Println("即将退出...")
			time.Sleep(3 * time.Second)
			os.Exit(0)
		default:
			fmt.Printf("指令不存在")
		}
	}
}
