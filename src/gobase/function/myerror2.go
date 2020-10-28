package main

import (
	"fmt"
)

func test1() {

	defer func() {
		if err := recover(); err != nil {
			println(err.(string))
			fmt.Println(err)
		}
	}()
	panic("我是一个:panic error!")
}

func test2() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	ch := make(chan int, 10)
	close(ch)
	ch <- 1
}

//延迟调用中引发的错误,可被后续延迟调用捕获,但仅最后一个被执行的错误可被捕获
func test3() {
	defer func() {
		fmt.Println(recover())
	}()
	defer func() {
		panic("defer panic1")
	}()
	defer func() {
		panic("defer panic2")
	}()

	panic("test panic")

}

func main() {
	test1()
	test2()
	test3()
}
