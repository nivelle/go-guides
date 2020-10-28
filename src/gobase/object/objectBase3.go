package main

import "fmt"

//人
type Person4 struct {
	name string
	sex  string
	age  int
}

// 自定义类型
// go支持只提供类型而不写字段名的方式，也就是匿名字段，也称为嵌入字段
type myStr string

// 学生
type Student4 struct {
	Person4
	int
	myStr
}

func main() {
	s1 := Student4{Person4{"5lmh", "man", 18}, 1, "bj"}
	fmt.Println(s1)
}
