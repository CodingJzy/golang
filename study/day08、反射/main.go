package main

import (
	"fmt"
	"reflect"
)

func main() {
	var a int

	// 取得变量a的类型对象,类型为reflect.Type()
	typeOfA := reflect.TypeOf(a)
	// 通过typeOfA类型对象的成员函数，可以分别获取变量的类型名(Name)：int、种类(Kind)：int
	fmt.Println(typeOfA.Name(), typeOfA.Kind())
}
