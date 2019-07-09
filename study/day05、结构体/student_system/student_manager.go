package main

import (
	"fmt"
	"strings"
)

// StudentManager struct
type StudentManager struct {
	items map[int64]*Student
}

// new manager
func NewStudentManager() *StudentManager {
	return &StudentManager{
		items: map[int64]*Student{},
	}
}

// AddStudent
func (s StudentManager) AddStudent(student *Student) {
	if _, ok := s.items[student.Id]; !ok {
		s.items[student.Id] = student
		return
	}
	fmt.Printf("学生%s已存在", student.Name)
}

// DeleteStudent
func (s StudentManager) DeleteStudent(id int64) {
	if _, ok := s.items[id]; ok {
		delete(s.items, id)
		fmt.Println("删除成功")
		return
	}
	fmt.Printf("该学生不存在")
}

// updateStudent
func (s StudentManager) UpdateStudent(id int64) {
	if v, ok := s.items[id]; ok {
		var updateField string
		fmt.Print(`
1、学号
2、姓名
3、年龄
4、性别
5、班级
请输入修改字段(,隔开)：`)
		_, _ = fmt.Scan(&updateField)
		fields := strings.Split(updateField, ",")
		var updateAfterId int64
		// 字符串切割
		for _, field := range fields {
			updateBeforeId := v.Id
			if updateAfterId != 0 {
				v = s.items[updateAfterId]
			}
			if field == "1" {
				// 删除原来的学生
				delete(s.items, updateBeforeId)
				fmt.Printf("请输入ID：")
				_, _ = fmt.Scanln(&updateAfterId)
				// 新增一个学生
				stu := NewStudent(updateAfterId, v.Name, v.Age, v.Sex, v.Class)
				s.items[updateAfterId] = stu
			}
			if field == "2" {
				fmt.Printf("请输入姓名：")
				_, _ = fmt.Scanln(&v.Name)
			}
			if field == "3" {
				fmt.Printf("请输入年龄：")
				_, _ = fmt.Scanln(&v.Age)
			}
			if field == "4" {
				fmt.Println("修改后", v.Sex)
				fmt.Printf("请输入性别：")
				_, _ = fmt.Scanln(&v.Sex)
			}
			if field == "5" {
				fmt.Printf("请输入班级：")
				_, _ = fmt.Scanln(&v.Class)
			}
		}
		fmt.Println("修改成功")
		return
	}
	fmt.Printf("该学生不存在")
}

// listStudent
func (s StudentManager) ListStudent() {
	fmt.Println("学生信息：")
	for _, v := range s.items {
		fmt.Printf("学号：%d\t姓名：%s\t年龄：%d\t性别：%s\t班级：%s\n", v.Id, v.Name, v.Age, v.Sex, v.Class)
	}
}
