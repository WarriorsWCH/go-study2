

package main 

import(
	"fmt"
	"net"
	"strings"
)

var onlineConns = make(map[string]net.Conn)
var messageQueue = make(chan string, 1000)
var quiteChan = make(chan bool)

func checkErr(err error) {
	if err != nil {
		panic(err)
	}
}

// 专门负责接收信息的协程
func ProcessInfo(connect net.Conn){

	buf := make([]byte, 1024)
	// defer connect.Close()
	// 协程退出时，说明客户端断开连接，所以要将当前连接从onlineConns删除掉
	defer func(conn net.Conn) {
		delete(onlineConns, fmt.Sprintf("%s",conn.RemoteAddr()))
		conn.Close()
	}(connect)

	for {
		numOfBytes, err := connect.Read(buf)
		
		if err != nil {
			break
		}

		// 如果接收字节数不为0 说明有消息发过来
		if numOfBytes != 0 {
			// 获取链接方的IP
			// remoteAddr := connect.RemoteAddr()
			// fmt.Println(remoteAddr,"接受信息：",string(buf[0:numOfBytes]))

			message := string(buf[0:numOfBytes])
			messageQueue <- message
		}
	}
}
// 处理消息的协程
func ConsumeMessage(){
	for {
		select{
			case message := <- messageQueue:
				// fmt.Println("解析")
				// 解析消息
				doProcessMessage(message)
			case <-quiteChan:
				break
		}
	}
}
// 解析消息
func doProcessMessage(message string) {
	// 示例：127.0.0.1:53959#你好#哈哈
	contents := strings.Split(message, "#")
	// 127.0.0.1   你好  哈哈
	if len(contents) > 1{
		addr := contents[0]
		// 处理空格
		addr = strings.Trim(addr, " ")
		sendMessage := strings.Join(contents[1:],"#")//你好#哈哈

		if connect, ok := onlineConns[addr]; ok {
			_, err := connect.Write([]byte(sendMessage))
			checkErr(err)
		}
	}
}

func main(){

	// 开启监听socket 监听在本地地址8080端口
	listen_socket, err := net.Listen("tcp", "127.0.0.1:8080")
	checkErr(err)

	defer listen_socket.Close()

	fmt.Println("server is waiting.....")

	go ConsumeMessage()

	for {
		// 接受（accept）客户端的连接请求
		// 等待刻画段连接请求，在没有客户端连接请求到之前，程序会一直阻塞在这个函数里
		connect, err := listen_socket.Accept()
		// 已经接收客户端请求连接，Aceepet函数创建并返回一个新的套接字，用于与客户端通信
		// 如果不再接受其他客户端的链接请求，可以关闭监听套接字了
		checkErr(err)

		// 存储 映射
		onlineConns[fmt.Sprintf("%s",connect.RemoteAddr())] = connect

		for item := range onlineConns{
			fmt.Println(item)
		}
		// 如果有客户端链接 则打开一个协程处理
		go ProcessInfo(connect)
	}
}


















