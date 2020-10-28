package main

import "fmt"

type User5 struct {
	id   int
	name string
}

func (user *User5) Test1(i int) {
	fmt.Printf("%p,%v \n", user, user)
}

func (user User5) Test2(i int) {
	fmt.Printf("%p,%v \n", user, user)
}

func main() {
	u1 := User5{1, "Tom"}
	//instance.method(args...)
	//隐式传递 receiver[方法调用者]
	u1.Test1(2)

	//<type>.func(instance, args...)
	//显式传递 receiver【方法调用者】
	//调用者类型.方法名(方法调用者,参数)
	(*User5).Test1(&u1, 3)

     fmt.Println("==================")
	u2 := User5{1, "Tom"}
	u2.Test2(5)


}
