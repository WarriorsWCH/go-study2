package main

import "fmt"

func main() {
	fmt.Println("Hello world!")

	//	var关键字是Go最基本的定义变量方式，与C语言不同的是Go把变量类型放在变量名后面
	var x int
	x = 1
	fmt.Println(x)

	//	:=  简短声明
	y, z := 'a', 'b'
	fmt.Println(y, z)

	//常量
	const hello int = 12
	//	数据类型：
	//	bool
	//	rune
	//	int8  int16 int32 int64
	//	byte
	//	uint8  uint16 uint32 uint64
	//	float32 float64
	//	complex64   complex128 复数
	//	string
	//	array slice
	//	map

	var arr [10]int
	arr[0] = 1
	arr[6] = 10
	fmt.Printf("%v\n", arr)

	arr2 := [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("%v 类型%T\n", arr2,arr2)

	//	动态数组
	arr3 := make([]int, 3, 5)
	arr3 = append(arr3, 5, 5, 7)
	fmt.Printf("%v 类型%T\n", arr3,arr3)
	//	cap 查看数组容量
	fmt.Printf("%v\n", cap(arr3))

	//make用于内建类型（map、slice 和channel）的内存分配
	var m1 map[string]string = make(map[string]string, 10)
	m1["name"] = "jack"
	m1["age"] = "12"
	fmt.Printf("%s\n", m1)

	//循环
	sum := 0
	for i := 0; i < 10; i++ {
		sum += i
	}
	fmt.Println(sum)

	//for代替while循环
	j := 1
	for j < 10 {
		j++
		fmt.Println(j)
	}
	// for 循环中声明和操作多个变量
	for x1, x2 := 10, 1; x2 <= 10 && x1 <= 19; x2, x1 = x2+1, x1+1 {
		fmt.Printf("x1=%d,x2=%d\n", x1, x2)
	}

	//range的使用非常简单，对于遍历array，*array，string它返回两个值分别是数据的索引和值，
	//遍历map时返回的两个值分别是key和value，遍历channel时，则只有一个返回数据
	nums := []int{1, 2, 3}
	for i, num := range nums {
		fmt.Printf("index:%d,value:%d\n", i, num)
	}

	kvs := map[string]string{"name": "low", "age": "29"}
	for k, v := range kvs {
		fmt.Printf("key=%s,value=%s\n", k, v)
	}

}
