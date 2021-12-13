package object

import (
	"fmt"
)

type Animal interface {
	ShowHair()
	SayWords(message string) string
}

type Origin interface {
	Run()
}

type Cat struct {
	hair string //implicit assignment of unexported field 'hair' in myinterface.Cat literal , 给外部包用的变量，必须首字母大写
	eyes string
}

//实现接口的方式就是让 对象有满足接口的方法
func (cat Cat) ShowHair() {
	fmt.Printf("猫咪的头发颜色。%s ,%s\n", cat.hair, cat.eyes)
}

func (cat *Cat) SayWords(message string) string {
	var eyes = cat.eyes
	return message + ":" + eyes
}

func (cat Cat) Run() {
	fmt.Println("这只眼睛是:", cat.eyes, ",它可以跑！")
}

//接口类型变量：接口类型变量能够存储所有实现了该接口的实例
func main() {
	cat := Cat{"Black", "Gree"}
	fmt.Println(cat.hair, cat.eyes)
	cat.ShowHair()
	cat1 := Cat{"Red", "Red"}
	result := cat1.SayWords("fuck")
	fmt.Println(result)

	cat1.Run()

	//定义一个接口类型变量
	var x Animal
	myCat1 := Cat{hair: "green", eyes: "big"}
	x = &myCat1
	x.ShowHair()

	var y Animal
	myCat2 := &Cat{hair: "green", eyes: "big"}
	y = myCat2
	sayMessage := y.SayWords("nivelle")
	fmt.Println(sayMessage)

	var z Origin
	myCat3 := Cat{hair: "purple", eyes: "purple"}
	z = myCat3
	z.Run()

}
