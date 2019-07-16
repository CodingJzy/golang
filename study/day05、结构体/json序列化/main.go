package main

import (
	"encoding/json"
	"fmt"
)

// 学生结构体
type Student struct {
	ID      int    `json:"id"`
	Gender  string `json:"gender"`
	Name    string `json:"name"`
	IsLogin bool   `json:"is_login"`
}

func main() {
	var stu1 = Student{
		ID:      1,
		Gender:  "男",
		Name:    "豪杰",
		IsLogin: false,
	}

	// 序列化：把内存里的数据转换为字符串，以便用于网络传输和数据交换
	if v, err := json.Marshal(stu1); err != nil {
		fmt.Println("JSON格式化出错")
		fmt.Println(err)
	} else {
		fmt.Println(v)
		fmt.Println(string(v))
	}

	// 反序列化：把json字符串转换为当前编程语言可用的对象
	str := `{"id":1,"gender":"男","name":"豪杰","is_login":false}`
	stu2 := new(Student)
	_ = json.Unmarshal([]byte(str), stu2)
	fmt.Println(stu2)
	fmt.Printf("stu2 类型：%T\tstu2值：%#v", stu2, stu2)
}
