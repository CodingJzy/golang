package main

import (
	"container/list"
	"fmt"
	"sync"
)

func testArray() {
	// 数组：固定大小的连续空间，同一种数据类型的集合。不能混用，和Python这种动态语言有点区别。

	// 声明数组：var 数组变量名 [元素数量]类型
	// 数组长度在声明时已经确定，或者编译器确定了。后期不能修改大小。
	var team [4]string
	team[0] = "Linux"
	team[1] = "Python"
	team[2] = "Java"
	team[3] = "Go"
	fmt.Println(team)

	// 声明并初始化一个空数组
	team1 := []string{}
	fmt.Println(team1)

	// 声明并初始化数组
	t1 := [4]string{"a", "b", "c", "d"}
	// 让编译器确定数组长度
	t2 := [...]int{1, 2, 3, 4}
	fmt.Println(t1)
	fmt.Println(t2)

	// 通过索引初始化
	t3 := [4]int{1: 1, 2: 2}
	fmt.Println(t3)

	// 数组的遍历  for range, k 为索引, v 为索引值
	for k, v := range t1 {
		fmt.Println(k, v)
	}

	// 练习1：求数组内的和
	array1 := [...]int{1, 3, 5, 7, 8}
	var sum int
	for _, v := range array1 {
		sum += v
	}
	fmt.Println("-------练习1：求和--------", sum)

}

func testSlice() {
	// 这和python的切片类似, 顾头不顾尾。是一种引用类型，他的内部结构包含地址、长度(len：当前切片内的元素个数)、容量(cap：底层数组的容下的最大元素数)
    // 切片必须初始化或者append()才能使用
	var t3 = [...]float64{1.1, 1.2, 1.3}

	// 从数组得到切片
	fmt.Println(t3[1:])

	t4 := [30]int{}
	for i := 0; i < 30; i++ {
		t4[i] = i + 1
	}
	fmt.Println(t4)
	fmt.Println("原有切片", t4[:])
	fmt.Println("中间切", t4[1:11])
	fmt.Println("从头开始", t4[:11])
	fmt.Println("到尾结束", t4[20:])

	// 切片重置
	fmt.Println("切片重置", t4[0:0])

	// 声明切片：var 切片名 元素类型
	// 声明一个字符串切片
	var slice1 []string
	// 声明一个整型切片
	var slice2 []int
	// 声明一个空切片
	var slice3 = []bool{}

	// 打印三个切片
	// []  []  []
	// 虽然值都是一样的，都是没有一个元素。但是空切片已经分配了内存，所以空切片不为nil。声明未使用的切片的默认值是nil，所以为false
	fmt.Println(slice1, slice2, slice3)
	// 判断切片为空的结果
	fmt.Println(slice1 == nil)
	fmt.Println(slice2 == nil)
	fmt.Println(slice3 == nil)

	// 使用make()函数构造切片
	// 格式：make([]type, size, cap)
	// 元素类型、长度、预分配长度(设置后，不影响size，只是提前分配空间，降低多次分配空间造成的性能问题)

	slice4 := make([]int, 5)
	slice5 := make([]int, 5, 10)
	// slice5 虽然内部的存储空间已经分配了10个，但是实际上使用了2个元素，容量并不会影响切片的长度。依然都是2
	fmt.Println(slice4)
	fmt.Println(slice5)
	fmt.Println(len(slice5))
	fmt.Println(len(slice5))

	// append()：为切片添加元素
	// 当切片容量不够时，可以通过append()扩容，最简单的扩容规律按两倍数扩充：1,2,4...
	// 扩容规律还有其他情况，这里暂不讨论
	var slice6 []int
	for i := 0; i < 20; i++ {
		slice6 = append(slice6, i)
		// 打印扩容情况
		fmt.Printf("长度：%d     cap：%d      指针：%p\n", len(slice6), cap(slice6), slice6)
	}

	// append多个元素
	slice6 = append(slice6, 21, 22, 23)
	fmt.Println(slice6)

	// 添加切片
	slice7 := []int{25, 26, 27, 28, 29, 30}
	slice6 = append(slice6, slice7...)
	fmt.Println(slice6)

	// 拷贝切片：切片是引用类型
	slice8 := []int{100, 200, 300}
	slice9 := slice8

	slice10 := make([]int, 3, 3)
	copy(slice10, slice8)
	slice10[2] = 250
	slice9[2] = 200
	fmt.Println(slice8)
	fmt.Println(slice9)
	fmt.Println(slice10)

	// go 语言本身没有对删除切片提供专用的语法或者接口，需要使用切片本身的特性来删除元素
	// 本质：以被删除的元素为分界点，将前后两个部分的内存重新连接起来
	// 删除某个元素：ｃ
	slice11 := []string{"a", "b", "c", "d"}
	slice11 = append(slice11[:2], slice11[3:]...)
	fmt.Println(slice11)

}

func testMap() {
	// 定义：map [keyType] valueType
	// 其实和python中的字典类似 key-value结构
	dic1 := make(map[string]int)
	dic1["age"] = 23
	dic1["money"] = 5201314
	fmt.Println(dic1)

	v1 := dic1["age"]
	v2 := dic1["name"]
	v2, ok := dic1["name"]

	fmt.Println(v1)
	fmt.Println(v2)
	fmt.Println(ok)

	// 另一种方式
	dic2 := map[string]string{
		"one":   "Python",
		"two":   "Go",
		"three": "Java",
		"four":  "Linux",
	}
	fmt.Println(dic2)

	// map 遍历 ， 无序的
	for k, v := range dic2 {
		fmt.Println(k, v)
	}

	for _, v := range dic2 {
		fmt.Println(v)
	}

	for k := range dic2 {
		fmt.Println(k)
	}

	// 要想结果排序，可以声明一个切片变量，遍历map的时候，把k append到切片中，利用内置函数sort对切片进行排序

	// 删除 delete(dic, key) 删除map中的键值对
	delete(dic2, "three")
	fmt.Println(dic2)

	// 清空map中的所有元素, go 语言没有提供任何清空所有元素的函数、方法。唯一方法就是重新make一个新的map 因为 go 语言中的并行垃圾回收效率比写一个清空函数高效多了。

	// 能够并发环境中使用的mao---sync.Map
	// 在并发的情况下。同时读写map会出现竞态问题。在非并非情况下。map性能更高。
	var sm sync.Map

	// 赋值
	sm.Store(1, "Python")
	sm.Store(2, "Go")
	sm.Store(3, "Linux")
	sm.Store(4, "Java")

	// 取值
	smV1, _ := sm.Load(1)
	fmt.Println(smV1)

	// 删除
	sm.Delete(1)

	// 遍历
	sm.Range(func(k, v interface{}) bool {
		fmt.Println(k, v)
		return true
	})
}

func testList() {
	// 在Go语言中, 和Python的列表也是有相同点的，都可以同时存储不同的数据类型
	// 通过container/list包来实现，内部实现的原理的双链表，能高效的进行任意位置的元素插入和删除操作

	// 列表的初始化
	// 通过container/list包的New方法初始化list
	l1 := list.New()
	// 通过声明初始化list
	var l2 list.List

	// 在列表中插入元素
	// 从尾部开始插入
	l1.PushBack("first")
	// 从头部开始插入
	l1.PushFront(67)
	// 添加的元素为列表
	l3 := list.New()
	l4 := list.New()
	l3.PushBack(1)
	l3.PushBack(2)
	l1.PushBackList(l3)
	l1.PushFrontList(l4)

	// 在列表中删除元素
	l2.PushBack(1)
	l2.PushBack(2)
	// 尾部添加后保存元素句柄
	ele := l2.PushBack(4)
	// 在元素4后面加入5
	l2.InsertAfter(5, ele)
	// 在元素4前面加入3
	l2.InsertBefore(3, ele)
	// 删除
	l2.Remove(ele)

	// 列表的遍历
	// 遍历双链表需要配合front()函数获取头元素，遍历时，只要不为空就一直继续下去。每一次遍历调用next
	for i := l2.Front(); i != nil; i = i.Next() {
		fmt.Println(i.Value)
	}

	for j := l1.Front(); j != nil; j = j.Next() {
		fmt.Println(j.Value)
	}

}

func main() {
	fmt.Println("容器：存储和组织数据的方式")

	// 数组
	// testArray()

	// 切片
	testSlice()

	// map
	// testMap()

	// list
	// testList()

}
