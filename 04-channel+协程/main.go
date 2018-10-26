package main

import (
	"fmt"
	"runtime"
	"strconv"
	"time"

	"github.com/astaxie/beego"
)

// 与传统的系统级线程和进程相比，协程的大优势在于其“轻量级”，
// 可以轻松创建上百万个而不会导致系统资源衰竭，而线程和进程通常多也不能超过1万个。这也是协程也叫轻量级线程的原因。

func Read(ch chan int) {
	value := <-ch
	fmt.Println("value:", value)
}
func Write(ch chan int) {
	ch <- 10
}

func test(ch chan int) {
	ch <- 1
	fmt.Println("ch 1")
	ch <- 1
	fmt.Println("ch 2")
	ch <- 1
	fmt.Println("come to end goroutine 1")
}
func main() {
	beego.Run()
	fmt.Println("hello go!")

	// channel是go语言级别提供的groutine间的通讯方式
	ch := make(chan int)
	go Read(ch)
	go Write(ch)

	time.Sleep(time.Second)
	fmt.Println("end of code")

	ch2 := make(chan int, 0) //等价于ch2 := make(chan int)都是不带缓冲的channel
	ch2 = make(chan int, 2)  //带缓冲的channel
	go test(ch2)

	time.Sleep(time.Second * 2)
	fmt.Println("running end")
	// 缓冲满之后阻塞，除非有goroutine对其进行操作
	<-ch2
	time.Sleep(time.Second)

	testSelect()

	testRoutine()
}

func testSelect() {
	ch := make(chan int)
	timeout := make(chan int, 1)

	go func() {
		time.Sleep(time.Second)
		timeout <- 1
	}()
	// Go语言直接在语言级别支持select关键字，用于处理异步IO问题
	select {
	case <-ch:
		fmt.Println("read ch")
	case <-timeout:
		fmt.Println("read timeout")
	}
	fmt.Println("select end of code")
}

func testRoutine() {
	ch := make(chan int)

	// 协程1
	go func() {
		for i := 0; i < 100; i++ {
			if i == 10 {
				// 主动出让cpu使用的话，需要导入runtime包
				runtime.Gosched()
				<-ch
			}
			fmt.Println("routine 1:", strconv.Itoa(i))
		}
	}()

	// 协程2
	go func() {
		for i := 100; i < 200; i++ {
			fmt.Println("routine 2:", strconv.Itoa(i))
		}
		ch <- 1
	}()

	time.Sleep(time.Second)
}
