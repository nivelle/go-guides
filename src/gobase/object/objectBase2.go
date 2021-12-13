package object

import "fmt"

//人
type Person3 struct {
	name string
	sex  string
	age  int
}

type Student3 struct {
	Person3
	id   int
	addr string
	//同名字段
	name string
}

func main() {
	var s Student3
	// 给自己字段赋值了
	s.name = "5lmh"
	fmt.Println(s)

	// 若给父类同名字段赋值，如下
	s.Person3.name = "枯藤"
	fmt.Println(s)
}
