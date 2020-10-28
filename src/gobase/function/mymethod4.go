package main

import "fmt"

// 参数 类型是 地址（引用传递）
func swap(x, y *int) {
	var temp int
	temp = *x
	*x = *y
	*y = temp
}

func swap2(x, y int) {
	var temp int
	temp = x
	x = y
	y = temp

}

func main() {
	// 无论是值传递，还是引用传递，传递给函数的都是变量的副本，不过，值传递是值的拷贝。引用传递是地址的拷贝，一般来说，地址拷贝更为高效。而值拷贝取决于拷贝的对象大小，对象越大，则性能越低
	// map、slice、chan、指针、interface默认以引用的方式传递
	var a, b = 1, 2
	swap(&a, &b)
	fmt.Println(a, b)

	var c, d = 1, 2
	swap2(c,d)
	fmt.Println(c, d)

}
