package main

import (
	"fmt"
	"reflect"
)

func ShowType() int32 {
	var myInt int32 = 10
	fmt.Print(myInt)
	return myInt
}

func typeConvert() {
	myInt := 10
	myDouble := 11.2
	myString := "nivelle"
	fmt.Println(myInt)
	fmt.Println(myDouble)
	fmt.Println(myString)
	var myByte byte = 'a'
	fmt.Printf("%s", myByte)
	var myBool bool
	fmt.Println("%s", myBool)
	var myString1 string = "哈哈嘿嘿"
	fmt.Println(myString1)
	string1 := "haha"
	string2 := "哈哈"
	fmt.Println(string1 + string2)
	a:=100
	// reflect.TypeOf(&a).Elem() 获取指针类型的元素类型
	va,vp := reflect.ValueOf(a),reflect.TypeOf(&a).Elem()
	fmt.Println("a的值:",va,"a指针的类型:",vp)
	fmt.Println("va canAddr:",va.CanAddr(),"va canSet:",va.CanSet())
}

type Rect struct {
	width  float64
	height float64
}

func (r *Rect) size() float64 {
	var resultValue = r.width * r.height
	return resultValue
}

func ShowObjectType() {
	var rObject  = Rect{100, 100}
	fmt.Println("对象获取值:", rObject)
	var r *Rect = &rObject
	fmt.Println("*是指针运算符,可以表示一个变量是指针类型，指针的值:", r)
	fmt.Println("*也表示一个指针变量所指向的存储单元,也就是这个地址所存储的值:", *r)
	fmt.Println("查看指针变量的地址:", &r)
}

func main()  {
	ShowType()
	typeConvert()
	ShowObjectType()
}
