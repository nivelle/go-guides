package main

import (
	"reflect"
	"fmt"
)

type MyInt int // 不允许为基本类型（int）对象定义方法，type 为类型取别名

// 定义对象（类）的方法： func (对象类型参数) 方法名（参数列表）（返回值列表）「」
func (obj MyInt) add(b int) int {
	return int(obj) + b
}

type student struct {
	id   int
	name string
	age  int
}

//对象调用时，会自动将对象传给obj
func (obj student) ShowStudent(nameOld string) string {
	return obj.name + nameOld
}

//修改对象则需要传入 对象地址
func (obj *student) ModifyStudent() {
	obj.age = 100
}

func CreateStudent(id int, name string, age int) student {
	var stu = student{id, name, age}
	return stu
}

type user struct {
	name string
	age  int
}

type manager struct {
	user
	title string
}

func (m manager) ManagerMethod1() {
	fmt.Println("manager method1:", m.name)
}

func (mp *manager) ManagerMethod2() {
	fmt.Println("manager method2:", mp.name)
}

func (u user) UserMethod1() {
	fmt.Println("user method1:", u.name)
}

// *user 表示 ptr是指向user的指针
func (uPtr *user) UserMethod2() {
	fmt.Println("user method2:", uPtr.name)
}

func ReachObjectMethod() {
	var m manager
	t := reflect.TypeOf(&m)
	fmt.Println("manager &m的类型:", t)
	fmt.Println("manager &m的类型的元素类型:", t.Elem())
	s := []reflect.Type{t, t.Elem()}
	fmt.Println("构造一个数组，包括指针类型和实例基类型:", s)
	for _, t := range s {
		fmt.Println("reflect方法遍历:", t)
		// 这个方法数量相对与调用方,如果外部调用,但是方法是私有的则访问不到
		fmt.Println("reflect方法遍历,方法数量:", t.NumMethod())
		for i := 0; i < t.NumMethod(); i++ {
			fmt.Println(" ", t.Method(i))
		}
	}

}
func ReachObjectElement() {
	var m manager
	t := reflect.TypeOf(&m)
	fmt.Println("manager的指针类型:", t.Kind())
	fmt.Println("manager的指针类型的名字:", t)

	//获取指针的基类型
	if t.Kind() == reflect.Ptr {
		t = t.Elem()
		fmt.Println("指针类型的基类型:", t)
	}
	for i := 0; i < t.NumField(); i++ {
		f := t.Field(i)
		fmt.Println("属性名：", f.Name, "属性类型：", f.Type, "属性位移：", f.Offset)
		//匿名字段结构
		if f.Anonymous {
			for x := 0; x < f.Type.NumField(); x++ {
				af := f.Type.Field(x)
				fmt.Println(" ", af.Name, af.Type)
			}
		}
	}
	name, _ := t.FieldByName("title")
	fmt.Println("FieldByName :", name.Name, name.Type)

}
