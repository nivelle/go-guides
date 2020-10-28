package main

import "fmt"

func ShowNoNameMethod(f func() int) string {
	var intValue = f()
	//匿名函数赋值给变量
	noMethod := func(s string) string {
		return "匿名函数返回值:" + s
	}("have no name")
	fmt.Println(noMethod)
	fmt.Println(intValue)
	return noMethod
}

// 函数返回值可以设置返回名字
func myInnerMethod() (returnStr string) {
	returnStr = "good boy"
	return returnStr
}

// 函数作为入参
func methodParams(fn func() string) string {
	return fn()
}

//定义函数类型
type FormatFunc func(s string, x, y int) string

func format(fn FormatFunc, s string, x, y int) string {
	return fn(s, x, y)
}

func main() {
	fmt.Println("no name method:")
	fn := myInnerMethod
	result := methodParams(fn)
	fmt.Println(result)

	//匿名函数作为参数
	var methodResult = func() int {
		fmt.Println("good girl")
		return 2
	}

	var endResult = ShowNoNameMethod(methodResult)
	fmt.Println(endResult)
	s1 := format(func(s string, x, y int) string {
		return fmt.Sprintf(s, x, y)
	}, "%d, %d", 10, 20)

	println("========", s1)
}
