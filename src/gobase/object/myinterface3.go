package main

import "fmt"

type Sayer interface {
	say()
}

type Mover interface {
	move()
}

type animal interface {
	Sayer
	Mover
}

type cat struct {
	name string
}

func (c cat) say() {
	fmt.Println("喵喵喵")
}

func (c cat) move() {
	fmt.Println("走走走")
}

func main() {
	//接口与接口间可以通过嵌套创造出新的接口。
	var x animal
	x = cat{name: "明明"}
	x.move()
	x.say()
}
