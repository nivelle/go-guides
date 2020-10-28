package main

import (
	"fmt"
)

func show(a interface{}) {
	//空接口作为函数参数:可以接收任意类型的函数参数
	fmt.Printf("type:%T value:%v\n", a, a)
}

func main() {
	//空接口是指没有定义任何方法的接口。因此任何类型都实现了空接口。
	//空接口类型的变量可以存储任意类型的变量
	var x interface{}
	s := "123"
	x = s
	fmt.Printf("type:%T value:%v \n", x, x)

	i := 100
	x = i
	fmt.Printf("type:%T value:%v \n", x, x)

	b := true
	x = b
	fmt.Printf("type:%T value:%v \n", x, x)

	fmt.Println()
	show("fuck")
	show(10000000)

	var studentInfo = make(map[string]interface{})
	studentInfo["name"] = "fuck"
	studentInfo["age"] = 18
	studentInfo["married"] = false

	fmt.Println(studentInfo)

}
