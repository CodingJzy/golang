package main

import "fmt"

// Student
type Student struct {
	Id    int64
	Name  string
	Age   int
	Sex   string
	Class string
}

// NewStudent return a new Student
//func NewStudent(id int64, name string, age int, sex, class string) *Student {
//	return &Student{
//		Id:    id,
//		Name:  name,
//		Age:   age,
//		Sex:   sex,
//		Class: class,
//	}
//}

func NewStudent(id int64, name string, age int, sex, class string) *Student {
	return &Student{
		Id:    id,
		Name:  name,
		Age:   age,
		Sex:   sex,
		Class: class,
	}
}

func InputStudent() *Student {
	var (
		id    int64
		name  string
		age   int
		sex   string
		class string
	)
	fmt.Println("请按提示录入信息")
	fmt.Printf("请输入学号：")
	fmt.Scanln(&id)
	fmt.Printf("请输入姓名：")
	fmt.Scanln(&name)
	fmt.Printf("请输入年龄：")
	fmt.Scanln(&age)
	fmt.Printf("请输入性别：")
	fmt.Scanln(&sex)
	fmt.Printf("请输入班级：")
	fmt.Scanln(&class)
	stu := NewStudent(id, name, age, sex, class)
	return stu
}
