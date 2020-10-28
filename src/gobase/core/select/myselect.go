package main

import (
	"fmt"
	"time"
)

func test1(ch chan string) {
	time.Sleep(time.Second * 5)
	ch <- "test1"
}
func test2(ch chan string) {
	time.Sleep(time.Second * 2)
	ch <- "test2"
}

func write(ch chan string) {
	for {
		select {
		// 写数据
		case ch <- "hello":
			fmt.Println("write hello")
		default:
			fmt.Println("channel full")
		}
		time.Sleep(time.Millisecond * 500)
	}
}

func main() {

	var c1, c2, c3 chan int //类型为int的通道
	var i1, i2 int
	i2 = 13
	select {
	case i1 = <-c1:
		fmt.Println("received ", i1, " from c1")
	case c2 <- i2:
		fmt.Println("send ", i2, " to c2")
	case i3, ok := <-c3:
		if ok {
			fmt.Println("receive", i3, "from c3")
		} else
		{
			fmt.Println("c3 is closed")
		}
	default:
		fmt.Println("no mesg")
	}

	// 2个管道
	output1 := make(chan string)
	output2 := make(chan string)
	// 跑2个子协程，写数据
	go test1(output1)
	go test2(output2)
	// 用select监控
	select {
	case s1 := <-output1:
		fmt.Println("s1=", s1)
	case s2 := <-output2:
		fmt.Println("s2=", s2)
	}

	// 创建2个管道
	intChan := make(chan int, 1)
	stringChan := make(chan string, 1)
	go func() {
		//time.Sleep(2 * time.Second)
		intChan <- 1
		time.Sleep(time.Second)
	}()
	go func() {
		stringChan <- "hello"
	}()
	select {
	case value := <-intChan:
		fmt.Println("int:", value)
	case value := <-stringChan:
		fmt.Println("string:", value)
	}
	fmt.Println("main结束")
	// 创建管道
	output3 := make(chan string, 10)
	// 子协程写数据
	go write(output3)
	// 取数据
	for s := range output3 {
		fmt.Println("res:", s)
		time.Sleep(time.Second)
	}
}

/**
  1. 每个case都必须是一个通信

  2. 所有channel表达式都会被求值

  3. 所有被发送的表达式都会被求值

  4. 如果任意某个通信可以进行,它就执行;其他被忽略;

  5. 如果有多个case都可以运行,Select会随机公平地选出一个执行,其他不会执行;

  6. 如果有default 子句，则执行该语句，如果没有default子句，select将阻塞，直到某个通信可以运行；go 不会重新对 channel 或值进行求值
 */
