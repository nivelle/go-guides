package main

import "fmt"

type User1 struct {
	id   int
	name string
}

// 匿名字段 包含了 user的字段
type Manager1 struct {
	User1
	title string
}

func (self *User1) ToString() string {
	return fmt.Sprintf("User: %p, %v", self, self)
}

func (self *Manager1) ToString() string {
	return fmt.Sprintf("Manager: %p, %v", self, self)
}

func main() {
	m := Manager1{User1{1, "Tom"}, "Administrator"}
	fmt.Println(m.ToString())
	fmt.Println(m.User1.ToString())
}
