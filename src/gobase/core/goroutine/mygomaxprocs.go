package main

import (
	"fmt"
	"runtime"
	"time"
)

func a() {
	for i := 1; i < 10; i++ {
		fmt.Println("A:", i)
	}
}

func b() {
	for i := 1; i < 10; i++ {
		fmt.Println("B:", i)
	}
}
func main() {
	//goroutine 和 OS线程是多对多的关系，即m:n ; GOMAXPROCS 就是 n
	// n:1 串行 n:>1 并行
	runtime.GOMAXPROCS(2)
	go b()
	go a()
	time.Sleep(time.Second)
}
