package object

import "fmt"

type Person2 struct {
	name string
	sex  string
	age  int
}

type Student2 struct {
	Person2
	id   int
	addr string
}

func main() {
	// 初始化
	s1 := Student2{Person2{"5lmh", "man", 20}, 1, "bj"}
	fmt.Println(s1)

	s2 := Student2{Person2: Person2{"5lmh", "man", 20}, id: 2, addr: "tj"}
	fmt.Println(s2)

	s3 := Student2{Person2: Person2{name: "5lmh"}}
	fmt.Println(s3)
}
