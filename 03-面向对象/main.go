package main

import "fmt"

type Integer struct {
	value int
}

//类型和作用在它上面定义的方法必须在同一个包里定义，这就是为什么不能在 int、float这些的类型上定义方法。
func (a Integer) compare(b Integer) bool {
	return a.value < b.value
}

type Point struct {
	px float32
	py float32
}

func (point *Point) setXY(px, py float32) {
	point.px = px
	point.py = py
}
func (point *Point) getXY() (float32, float32) {
	return point.px, point.py
}

// 方法定义成指针和不是指针的区别
func (point Point) setXY2(px, py float32) {
	fmt.Println("setXY2",point)
	point.px = px
	point.py = py
}
func (point Point) getXY2() (float32, float32) {
	return point.px, point.py
}


type Person struct {
	name string
	age  int
}

func (p Person) getAttr() (string, int) {
	return p.name, p.age
}
// 继承 Student继承Person
type Student struct {
	Person
	id int
}

func (s Student) getId() int {
	return s.id
}
// 在Go语言中，一个类只需要实现了接口要求的所有函数，我们就说这个类实现了该接口。    
type Amimal interface{
	Fly() bool
	Run() bool
}


type Bird struct{

}
func (bird Bird) Fly() bool{
	fmt.Println("fly。。。。")
	return true
}
func (bird Bird) Run() bool{
	fmt.Println("run。。。。")
	return true
}
func main() {
	fmt.Println("hello go!")

	a := Integer{1}
	b := Integer{2}

	fmt.Printf("%v\n", a.compare(b))

	//new()指针变量
	point := new(Point)
	point.setXY(1.222, 3.333)
	px, py := point.getXY()
	fmt.Println(px, py)
	point.setXY2(2.222, 6.333)
	px, py = point.getXY2()
	fmt.Println(px, py)


	student := new(Student)
	student.name = "jack"
	student.age = 10
	// student可以调用person的方法
	name, age := student.getAttr()
	fmt.Println(name, age)

	var animal Amimal
	bird := new(Bird)

	animal = bird
	animal.Fly()
	animal.Run()

}
