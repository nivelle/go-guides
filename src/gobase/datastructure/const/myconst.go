package main

import "fmt"

const (
	const1 = iota
	const2 //在常量中如果不指定类型和初始化值，则与上一行非空常量右值相同
	const3 = 100
	const4
	const5 = iota
)

func MyConst() {
	const name int = 100    //指定类型
	const name1 = "nivelle" //类型推断

	fmt.Println(name)
	fmt.Println(name1)

	fmt.Println("iota 为go语言的枚举:")
	fmt.Println(const1)
	fmt.Println(const2)
	fmt.Println(const3)
	fmt.Println(const4)
	fmt.Println(const5)

}

func main()  {
	MyConst()
}
