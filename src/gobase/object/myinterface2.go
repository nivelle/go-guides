package main

import "fmt"

type WashingMachine interface {
	wash()
	dry()
}

type dryer struct {
}

func (d dryer) dry() {
	fmt.Println("甩一甩")
}

type haier struct {
	//接口的方法可以通过在类型中嵌入其他类型或者结构体来实现。
	dryer
}

func (h haier) wash() {
	fmt.Println("洗刷刷")
}

func main() {
	var object1 = haier{}
	object1.dryer.dry()
	object1.wash()
}
