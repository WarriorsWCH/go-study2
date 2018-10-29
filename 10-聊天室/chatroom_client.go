

package main 

import(
	"fmt"
	"net"
	"bufio"
	"os"
	"strings"
)

func checkErr(err error) {
	if err != nil {
		fmt.Printf("Error:%s", err.Error())
	}
}

func MessageSend(connect net.Conn){
	var input string
	for {
		reader :=  bufio.NewReader(os.Stdin)
		data, _, _ := reader.ReadLine()
		input = string(data)

		if strings.ToUpper(input) == "EXIT" {
			connect.Close()
			break
		}

		_, err := connect.Write([]byte(input))
		if err != nil {
			connect.Close()
			fmt.Println("failed")
			break
		}
	}
}


func main(){
	connect, err := net.Dial("tcp", "127.0.0.1:8080")
	checkErr(err)

	defer connect.Close()

	// connect.Write([]byte("hello socekt"))
	// 开启协程 发送消息
	go MessageSend(connect)

	buf := make([]byte, 1024)
	for {
		// fmt.Printf("main")
		numOfBytes, err := connect.Read(buf)
		if err != nil {
			fmt.Printf("你一退出，欢迎下次光临")
			// Error:read tcp 127.0.0.1:53488->127.0.0.1:8080: use of closed network connection
			// break
			os.Exit(0)
		}
		fmt.Println("reciver server message:",string(buf[0:numOfBytes]))
	}

	fmt.Println("it is over")
}






