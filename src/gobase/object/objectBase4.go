package main

import "fmt"

//人
type Person struct {
	name string
	sex  string
	age  int
}

// 学生
type Student struct {
	*Person
	id   int
	addr string
}

func main() {
	s1 := Student{&Person{"5lmh", "man", 18}, 1, "bj"}
	fmt.Println(s1)
	fmt.Println(s1.name)
	fmt.Println(s1.Person.name)

}
