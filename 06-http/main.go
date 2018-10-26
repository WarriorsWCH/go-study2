

package main 

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"strings"
)

func main() {

	resp, err := http.Post(
					"https://www.baidu.com",
					"application/x-www-form-urlencoded",
					strings.NewReader("id=1"))
	if err != nil {
		panic(err)
	}

	defer resp.Body.Close()

	fmt.Println(resp)
	body, err := ioutil.ReadAll(resp.Body)
	fmt.Println(string(body))

	server()
}

func server() {
	http.HandleFunc("/hello", func(w http.ResponseWriter, r *http.Request){
			w.Write([]byte("hello world"))
		})

	// 使用 net/http 包提供的 http.ListenAndServe() 方法，可以在指定的地址进行监听， 开启一个HTTP
	http.ListenAndServe("127.0.0.1:8080", nil)
}






















