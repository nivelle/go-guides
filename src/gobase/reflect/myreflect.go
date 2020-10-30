package main

import (
	"fmt"
	"reflect"
)

func main() {

	var x float64 = 3.4

	t := reflect.TypeOf(x)
	fmt.Println("type:", t)

	v := reflect.ValueOf(x)
	fmt.Println("value1:", v)

	var y float64 = 5.902

	v2 := reflect.ValueOf(y) // v2 is reflect.Value
	fmt.Println("value2:", v2)

	//interface对象通过.(float64)类型断言获取float64类型的值
	var v3 float64 = v2.Interface().(float64)
	fmt.Println("value3:", v3)

	var z float64 = 6.01
	v4 := reflect.ValueOf(z)

	//reflect.Value提供了Elem()方法，可以获得指针向指向的value
	var v6 float64 = 7.98
	v7 := reflect.ValueOf(&v6)
	v7.Elem().SetFloat(7.99)
	fmt.Println("v7:", v7.Elem().Interface())

	defer func() {
		err := recover()
		fmt.Println(" 内部尝试捕获异常")
		if err != nil {
			fmt.Println("捕获了异常，异常内容:", err)
		} else {
			fmt.Println("无内鬼")
		}
	}()
	v4.SetFloat(6.02)
	var v5 = v4.Interface().(float64)
	fmt.Println("value4:", v5)

}
