package main

import (
	"time"
	"fmt"
)

func main() {
	//定时器等待2秒,然后通知一个通道C
	timer1 := time.NewTimer(time.Second * 2)
	fmt.Println(timer1)
	//直到定时器的通道C明确的发送了定时器失效的值之前,将一直阻塞
	<-timer1.C
	fmt.Println("time 1 is ok")

	timer2 := time.NewTimer(time.Second * 1)
	go func() {
		<-timer2.C
		//主程阻塞了，然后<-timer2.C也在阻塞，提前获取了定时器执行完毕的信号
		fmt.Println("timer 2 is ok")
	}()
	time.Sleep(time.Second * 5)
	//停止执行器，如果已经执行完毕，则返回false
	stop2 := timer2.Stop()
	if stop2 {
		fmt.Println("timer 2 is stopped")
	} else {
		fmt.Println("timer2 is done")
	}
}
