package main

import (
	"reflect"
	"fmt"
)

func showReflect() {

	a := 10
	t := reflect.TypeOf(a)
	fmt.Println("变量a的类型:", t)
	fmt.Println("类型名字:", t.Name())
	fmt.Println("类型的类型:", t.Kind())

	v := reflect.ValueOf(a)
	fmt.Println("变量a的值:", v)

	//通过反射创建指定参数和类型的数组
	arrayOfValue := reflect.ArrayOf(3, reflect.TypeOf(0.00))
	fmt.Println("arrayOf:", arrayOfValue)

	//基类型和指针类型
	x := 100
	tx, tp := reflect.TypeOf(x), reflect.TypeOf(&x)
	fmt.Println("反射值的类型和指针类型", tx, tp)
	fmt.Println(tx, tp, tx == tp)
	fmt.Println("txKind=", tx.Kind(), ";tpKind=", tp.Kind())
	fmt.Println("txName=", tx.Name(), ";tpName=", tp.Name())

	//切片的基础类型，elem来查询
	var mySlice = []int{}
	fmt.Println("切片的基类型：", reflect.TypeOf(mySlice).Elem())

	myMap := map[string]int{}
	fmt.Println("map 的基类型：", reflect.TypeOf(myMap).Elem())
	myMap1 := map[int]string{}
	fmt.Println("map 的基类型,指的是value的类型：", reflect.TypeOf(myMap1).Elem())

}

func main() {
	showReflect()
}
