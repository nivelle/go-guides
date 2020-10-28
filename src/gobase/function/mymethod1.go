package main

//1. 包名一般要和所在的目录同名，也可以不同，包名中不能包含- 等特殊符号;
//2. 一个文件夹下的所有源码文件只能属于同一个包，同样属于同一个包的源码文件不能放在多个文件夹下

import "fmt"

func innerMethod() {
	var myInt int32 = 10
	fmt.Print(myInt)
}

func OutMethod() {
	var myInt2 int32 = 100
	fmt.Println("这是mymethod-》OutMethod")
	fmt.Print(myInt2)
	fmt.Println("外部方法调用内部方法")
	forMethod()
	forMethod2()
	forMethod3()
}

func forMethod() {
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}
}

var jIndex  = 1
func forMethod2()  {
	for {
		if jIndex<=5{
			fmt.Println(jIndex)
		}else {
			break
		}
		jIndex++
	}
}

func forMethod3()  {
	for (jIndex<6){
		fmt.Println(jIndex)
	}
}
func ForMethod4()  {
	slice :=[]int {0,1,2,3}
	for index,value:=range slice {
		fmt.Println(index,value)
	}
}

func main()  {
	ForMethod4()

	OutMethod()


}
