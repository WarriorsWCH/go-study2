
package main 

import (
	"fmt"
	"encoding/json"
	"crypto/md5"
)

type Student struct {
	Name string `json:"student_name"`
	Age int
}

func main() {

	// 对数组型的json编码
	x := [5]int{1,2,3,4,5}
	s, err := json.Marshal(x)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(s))

	// 对map类型进行json编码
	m := make(map[string]interface{})
	m["name"] = "jack"
	m["age"] = 200
	s2, err2 := json.Marshal(m)
	if err2 != nil {
		panic(err2)
	}

	fmt.Println(string(s2))


	// 对结构体进行编码
	student := Student{"tom",33}
	s3, err3 := json.Marshal(student)
	if err3 != nil {
		panic(err3)
	}

	fmt.Println(string(s3))

	// json解码
	var s4 interface{}
	json.Unmarshal([]byte(s3), &s4)
	fmt.Printf("解码后：%v\n",s4)

	myMD5 := md5.New()
	myMD5.Write([]byte("abcd1234哈哈"))
	r := myMD5.Sum([]byte(""))
	fmt.Printf("\n%x\n",r)
}















