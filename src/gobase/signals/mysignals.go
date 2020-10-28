package main

import (
	"os"
	"os/signal"
	"syscall"
	"fmt"
)

func main() {
	sings := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	//注册通道 接收 ： syscall.SIGINT, syscall.SIGTERM 信号
	signal.Notify(sings, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		//阻塞 等通道有数据
		sig := <-sings
		fmt.Println()
		fmt.Println(sig)
		done <- true
	}()

	fmt.Println("awaiting signal")
	//阻塞 等待子协程 的执行完毕
	<-done
	fmt.Println("exiting")

}
