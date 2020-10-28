package main

import (
	"fmt"
	"runtime"
)

func main() {

	go func(s string) {
		for i := 0; i < 2; i++ {
			fmt.Println(s)
		}
	}("匿名函数默认参数:world")
	for i := 0; i < 2; i++ {
		//让出CPU时间片，重新等待安排任务
		runtime.Gosched()
		fmt.Println("hello")
	}
}
