package main

import "fmt"

type Test struct {
	param1Int int
	param2Str string
}

// 无参数、无返回值
func (receiver Test) method0() {
	fmt.Println("method0:", receiver)
}

// 单参数、无返回值
func (receiver Test) method1(i int) {
	receiver.param1Int = i
	fmt.Println("method1:", receiver)
}

// 多参数、无返回值
func (receiver Test) method2(x string, y int) {
	receiver.param1Int = y
	receiver.param2Str = x
	fmt.Println("method2:", receiver)

}

// 无参数、单返回值
func (receiver Test) method3() (i int) {
	fmt.Println("method3:", receiver)
	return receiver.param1Int
}

// 多参数、多返回值
func (receiver Test) method4(x string, y int) (z int, err error) {
	receiver.param1Int = y
	receiver.param2Str = x
	fmt.Println("method4:", receiver)
	return receiver.param1Int, nil
}

// 无参数、无返回值
func (receiverPtr *Test) method5() {
	fmt.Println("method5 param1:", receiverPtr.param1Int)
	fmt.Println("method5 param3:", receiverPtr.param2Str)
	fmt.Println("method5 :", *receiverPtr)

}

// 单参数、无返回值
func (receiverPtr *Test) method6(i int) {
	fmt.Println("method6 before :", receiverPtr)
	receiverPtr.param1Int = i
	fmt.Println("method6 after :", receiverPtr)
	return
}

func (receiverPtr *Test) method7(x int, y string) {
	fmt.Println("method7 before :", receiverPtr)
	receiverPtr.param1Int = x
	receiverPtr.param2Str = y
	fmt.Println("method7 after :", receiverPtr)
}

func (receiverPtr *Test) method8() (i int) {
	fmt.Println("method8 before :", receiverPtr)
	return receiverPtr.param1Int
}

func (receiverPtr *Test) method9(x int, y string) (z int, err error) {
	fmt.Println("method9 before :", receiverPtr)
	receiverPtr.param1Int = x
	receiverPtr.param2Str = y
	fmt.Println("method9 after :", receiverPtr)
	return receiverPtr.param1Int, nil
}

func main() {
	//对于普通函数，接收者为值类型时，不能将指针类型的数据直接传递，反之亦然
	var test = Test{1, "nivelle"}
	test.method4("jessy", 2)
	fmt.Println("原来的对象不会改变:", test)

	test.method9(3, "jessy")
	fmt.Println("原来的对象被改变:", test)

	var testPtr = &test
	testPtr.method9(4, "jessy")
	fmt.Println("原来的对象被改变:", test)

}
