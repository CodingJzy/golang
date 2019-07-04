package main

import (
	"fmt"
	"reflect"
)

type dummy struct {
	a int
	b string
	float64
	bool
	//嵌入字段
	next *dummy
}

type S struct {
}
type S1 struct {
	name string
}

func (s S1) GetName(name string) {
	fmt.Println(name, "----------------")
}

func testValue() {
	a := 1024

	// 获取变量a的反射值对象
	valueOfA := reflect.ValueOf(a)

	// 获取interface{} 类型的值，通过类型断言转换
	getA1 := valueOfA.Interface().(int)

	// 获取64位的值，强制类型转换为int类型
	getA2 := int(valueOfA.Int())

	// 打印
	fmt.Println(getA1, getA2)
}

func testStructValue() {

	// 值包装结构体
	d := reflect.ValueOf(dummy{
		next: &dummy{},
	})

	// 获取字段数量
	fmt.Println("字段数量(NumField)", d.NumField())

	// 获取索引为1的字段(b字段)
	stringField := d.Field(2)

	// 输出字段类型
	fmt.Println(stringField.Type())

	// 根据名字查找字段
	fmt.Println(d.FieldByName("b").Type())

	// 多层查找：next字段的3索引的值
	fmt.Println(d.FieldByIndex([]int{4, 3}))
}

func testPanDuan() {
	/*
		isNil()：判断是否为nil，常判断指针是否为空
		isValid()：判断是否有效，常判断返回值是否有效
	*/

	// 声明一个int的空指针。默认值为nil
	var a *int
	fmt.Println("判断a是否为nli：", reflect.ValueOf(a).IsNil())

	// nli值
	fmt.Println("判断nli的有效性：", reflect.ValueOf(nil).IsValid())

	// 实例化一个结构体
	s := S{}
	s1 := S1{}
	s1.GetName("getName")

	fmt.Println("判断结构体s字段name的有效性：", reflect.ValueOf(s).FieldByName("name").IsValid())
	fmt.Println("判断结构体s1字段name的有效性：", reflect.ValueOf(s1).FieldByName("name").IsValid())
	fmt.Println("判断结构体s的method的有效性：", reflect.ValueOf(s).MethodByName("GetName").IsValid())
	fmt.Println("判断结构体s1的method的有效性：", reflect.ValueOf(s1).MethodByName("GetName").IsValid())

	// 实例化一个map
	mp := map[string]string{}
	mp1 := map[string]string{
		"name": "name",
	}
	// 尝试从map中查找下一个不存在的键
	fmt.Println("判断mp中是否存在键name：", reflect.ValueOf(mp).MapIndex(reflect.ValueOf("name")).IsValid())
	fmt.Println("判断mp1中是否存在键name：", reflect.ValueOf(mp1).MapIndex(reflect.ValueOf("name")).IsValid())
}

func testChangeValue() {
	/*
		判定及获取元素的相关方法：
			Elem()：取值指向的元素值，类似于*操作。
			Addr()：对可寻址的值返回其地址，类似于&操作。
			CanAddr()：表示值是否可寻址
			CanSet()：返回值能否被修改
	*/

}

func main() {
	// 反射不仅可以获取值的类型信息，还可以动态获取或设置变量的值，使用reflect.Value

	// 使用反射值对象获取被包装的值
	//testValue()

	// 使用反射访问结构体的成员字段的值
	//testStructValue()

	// 反射对象的空和有效性判断
	//testPanDuan()

	// 利用反反射值对象修改变量的值
	testChangeValue()
}
