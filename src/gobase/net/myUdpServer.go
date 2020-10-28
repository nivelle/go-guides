package main

import (
	"net"
	"fmt"
)

func main() {
	listen, err := net.ListenUDP("udp", &net.UDPAddr{
		IP:   net.IPv4(0, 0, 0, 0),
		Port: 30000,
	})

	if err != nil {
		fmt.Println("listen failed,err:", err)
		return
	}

	defer listen.Close()

	for {
		var data [1024]byte
		//接收数据
		n, addr, err := listen.ReadFromUDP(data[:])
		if err != nil {
			fmt.Println("read udp failed, err:", err)
			continue
		}
		fmt.Printf("udp 服务端 收到数据, data:%v addr:%v count:%v\n", string(data[:n]), addr, n)
		_, err = listen.WriteToUDP(data[:n], addr)
		if err != nil {
			fmt.Println("write to udp failed, err:", err)
			continue
		}
		fmt.Println("服务端发送数据成功,data:", string(data[:n]))

	}
}
