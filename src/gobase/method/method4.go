package main

import "fmt"

/**
    • 类型 T 方法集包含全部 receiver T 方法。
    • 类型 *T 方法集包含全部 receiver T + *T 方法。
    • 如类型 S 包含匿名字段 T,则 S 和 *S 方法集包含 T 方法。
    • 如类型 S 包含匿名字段 *T,则 S 和 *S 方法集包含 T + *T 方法。
    • 不管嵌入 T 或 *T，*S 方法集总是包含 T + *T 方法。
 */
type T struct {
	int
}

func (t T) test1() {
	fmt.Println(t, "类型 -->test1方法->方法集包含全部 receive T 方法。")
}

func (t *T) test2() {
	fmt.Println(t, "类型 -->test2方法->方法集包含全部 receive T 方法 和 *T 方法。")
}

type S struct {
	T
}

func main() {
	t1 := T{1}
	t1.test1()
	t1.test2()
	fmt.Printf("t1 is : %v \n", t1)
	fmt.Println("===============")
	t2 := &t1
	t2.test1()
	t2.test2()
	fmt.Printf("t2 is : %v \n", t2)

	fmt.Println("===============")

	s1 := S{T: T{2}}
	s1.test1()
	s1.test2()

	s2 := &s1
	s2.test1()
	s2.test2()

}
