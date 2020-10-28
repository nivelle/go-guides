package main

import "fmt"

// 函数 pings参数只允许 接收数据
func ping(pings chan<- string, msg string) {
	pings <- msg
	//msg2:=<-pings
	fmt.Println(msg)
}

//pings 接收数据，pongs 发送数据
func pong(pings <-chan string, pongs chan<- string) {
	msg := <-pings
	fmt.Println("msg:=<-pings:", msg)
	pongs <- msg
	fmt.Println("pongs<-msg:", msg)
}

func main() {
	pings := make(chan string, 1)
	pongs := make(chan string, 1)
	ping(pings, "pass message")
	pong(pings, pongs)
	fmt.Print(<-pongs)
}
