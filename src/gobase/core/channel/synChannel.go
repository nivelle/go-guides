package main

import (
	"fmt"
	"time"
)

func worker(done chan bool) {
	fmt.Println("working...")
	time.Sleep(time.Second)
	fmt.Println("done")
	done <- true
}

func main() {
	done := make(chan bool, 1)
	fmt.Println("主函数开始运行")
	go worker(done)
	fmt.Println("主函数阻塞，等待通道有数据")
	<-done
	fmt.Println("主函数解除阻塞，从通道拿到了信号")
}
