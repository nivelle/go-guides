package main

import "fmt"

func main() {
	ShowDefer()

	deferClosure()
}

/**
    1. 关键字 defer 用于注册延迟调用。
    2. 这些调用直到 return 前才被执。因此，可以用来做资源清理。
    3. 多个defer语句，按先进后出的方式执行。
    4. defer语句中的变量，在defer声明时就决定了。
 */
func ShowDefer() {
	var variable = 1
	fmt.Println("defer before variable is:", variable)
	defer fmt.Println("defer 1")
	defer fmt.Println("defer 2")
	defer func() {
		variable = 2
		fmt.Println("defer variable is:", variable)
	}()
	fmt.Println("defer after")
	fmt.Println(variable)
}

func deferClosure() {
	var whatever [5]struct{}
	for i := range whatever {
		defer func() {
			defer fmt.Printf("%d", i)
			fmt.Println()
		}()
	}
	var whatever2 [5]struct{}
	for i := range whatever2 {
		defer fmt.Printf("%d", i)
		fmt.Println()

	}
}
