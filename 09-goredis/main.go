package main

import (
	"fmt"

	"github.com/astaxie/goredis"
)

// Redis是一个开源的、使用C语言编写的、支持网络交互的、可基于内存也可持久化的Key-Value数据库。

// Redis 优势
// 性能极高 – Redis能读的速度是110000次/s,写的速度是81000次/s 。

// 丰富的数据类型 – Redis支持二进制案例的 Strings, Lists, Hashes, Sets 及 Ordered Sets 数据类型操作。

// 原子 – Redis的所有操作都是原子性的，同时Redis还支持对几个操作全并后的原子性执行。

// 丰富的特性 – Redis还支持 publish/subscribe, 通知, key 过期等等特性。

// Redis与其他key-value存储有什么不同？
// Redis有着更为复杂的数据结构并且提供对他们的原子性操作，这是一个不同于其他数据库的进化路径。Redis的数据类型都是基于基本数据结构的同时对程序员透明，无需进行额外的抽象。

// Redis运行在内存中但是可以持久化到磁盘，所以在对不同数据集进行高速读写时需要权衡内存，因为数据量不能大于硬件内存。在内存数据库方面的另一个优点是，相比在磁盘上相同的复杂的数据结构，在内存中操作起来非常简单，这样Redis可以做很多内部复杂性很强的事情。同时，在磁盘格式方面他们是紧凑的以追加的方式产生的，因为他们并不需要进行随机访问。

func main() {

	var client goredis.Client
	client.Addr = "127.0.0.1:6379"
	// set(key, value)：给数据库中名称为key的string赋予值value
	err := client.Set("test", []byte("hello redis"))
	checkErr(err)

	// get(key)：返回数据库中名称为key的string的value
	res, err := client.Get("test")
	checkErr(err)

	fmt.Println(string(res))

	f := make(map[string]interface{})
	f["name"] = "tom"
	f["age"] = 22
	f["sex"] = "fmale"

	// hmset(key, (fields))：向名称为key的hash中添加元素field
	err = client.Hmset("test_hash", f)
	checkErr(err)

	// sorted set commands
	_, err = client.Zadd("test_zadd", []byte("hello"), 100)
	checkErr(err)

	// https://github.com/astaxie/goredis/blob/master/redis.go文档

}

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}
