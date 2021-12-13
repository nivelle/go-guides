package object

import "fmt"

type Person1 struct {
	Id int
	Name string
}

type Student1 struct {
	Person1
	Score int //属性外部访问 则 属性也需要大写
}

func (p *Person1) Print()  {
	fmt.Printf("id：%d\n", p.Id)
	fmt.Printf("姓名：%s\n", p.Name)

}

func (stu *Student1) Print() {
	fmt.Printf("id：%d\n", stu.Id)
	fmt.Printf("姓名：%s\n", stu.Name)
	fmt.Printf("score：%d\n", stu.Score)
}

