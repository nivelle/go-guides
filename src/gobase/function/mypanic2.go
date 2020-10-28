package main

import (
	"fmt"
)

//捕获函数 recover 只有在延迟调用内直接调用才会终止错误，否则总是返回 nil。任何未捕获的错误都会沿调用堆栈向外传递。
func test21() {
	defer func() { fmt.Println("捕获了异常:", recover()) }() //有效
	defer recover()                                     //无效
	defer fmt.Println(recover())                        //无效
	defer func() {
		func() {
			println("defer inner")
			recover() //无效
		}()
	}()
	panic("test panic")
}

func except() {
	fmt.Println("except 我捕获了异常:")
	fmt.Println(recover())
}

func test22() {
	defer except()
	panic("test2 panic")
}

//如果需要保护代码 段，可将代码块重构成匿名函数，如此可确保后续代码被执
func test23(x, y int) {
	var z int
	func() {
		defer func() {
			if err := recover(); err != nil{
				fmt.Println(err)
				z = 0
			}
		}()
		panic("test3 panic")
		z = x / y
		return
	}()
	fmt.Printf("x /y = %d \n", z)
}
func main() {
	//test1()
	//test2()
	test23(2, 2)
}
