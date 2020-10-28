package main

import "fmt"

/**
  intSeq函数返回另一个在intSeq函数体内定义的匿名函数。这个返回的函数使用闭包的方式隐藏变量i
 */
func intSeq() (func() int) {
	i := 0
	return func() int {
		i += 1
		return i
	}
}

func main()  {
	nextInt:=intSeq()
	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newNextInt:=intSeq()
	fmt.Println(newNextInt())
}



