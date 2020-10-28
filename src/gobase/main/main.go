package main

import (
	"fmt"
)
import OB "gobase/object" //别名 //包名是从GOPATH/src/ 后开始计算的，使用/ ;绝对路径导入
import (
	. "gobase/object/extends"//这种格式相当于把 mymethod 包直接合并到当前程序中，在使用 mymethod 包内的方法是可以不用加前缀 mymethod.，直接引用
)

func main() {
	fmt.Println("hello my first go ")

	fmt.Print("object========")
	var student = OB.CreateStudent(1, "jessy", 18)
	fmt.Println(student)
	fmt.Println(student.ShowStudent("nivelle============="))
	fmt.Println(student)
	student.ModifyStudent()
	//修改对象
	fmt.Println(student)
	fmt.Println("object & reflect ========")
	OB.ReachObjectElement()
	OB.ReachObjectMethod()
	fmt.Println("method========")

	fmt.Println("extends========")
	var stu Student
	stu.Person = Person{1, "fuck"}
	stu.Score = 10
	stu.Print()
}
