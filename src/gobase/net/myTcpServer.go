package main

import (
	"net"
	"bufio"
	"fmt"
)

func process(conn net.Conn) {
	defer conn.Close()

	for {
		reader := bufio.NewReader(conn)
		var buf [128]byte
		n, err := reader.Read(buf[:]) //读取数据
		if err != nil {
			fmt.Println("read from client failed, err:", err)
			break
		}
		receiveStr := string(buf[:n])
		fmt.Println("收到client端发来的数据：", receiveStr)
		conn.Write([]byte("响应给客户端:" + receiveStr )) // 发送数据
	}
}

func main() {
	listen, err := net.Listen("tcp", "127.0.0.1:20000")
	if err != nil {
		fmt.Println("listen failded ,err:", err)
		return
	}
	for {
		conn, err := listen.Accept()
		if err != nil {
			fmt.Println("accept failded,err:", err)
			continue
		}
		// 启动一个goroutine处理连接
		go process(conn)
	}
}
