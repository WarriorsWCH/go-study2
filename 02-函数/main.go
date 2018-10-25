package main

import "fmt"

// 参数值类型 传递的是值
func swap(a int, b int) (int, int) {
	return b, a
}
// 传递的是指针
func sdd(c *int) *int {
	*c = *c + 1
	return c
}
// 结构体
type Human struct {
	name string
	age  int
}
// 结构的字段还是结构体
type Student struct {
	Human
	id int
}

func main() {
	fmt.Println("Hello go!")

	a := 1
	b := 2
	a, b = swap(a, b)
	fmt.Printf("%d,%d\n", a, b)

	c := 3
	sdd(&c)
	fmt.Printf("%d\n", c)
	// defer是Go语言提供的关键字，常用来释放资源，会在函数返回之前进行调用
	for i := 1; i < 5; i++ {
		//调用顺序类似于栈，越后面的defer表达式越先被调用
		defer fmt.Println(i)
	}
	fmt.Println("before defer")

	defer func() {
		fmt.Println("defer 1")
	}()

	f := func() {
		fmt.Println("defer 2")
	}
	defer f()

	h := Human{"lee", 20}
	fmt.Printf("%v\n", h)

	s := Student{Human{"jack", 40}, 11111}
	fmt.Printf("%v\n", s)

}
