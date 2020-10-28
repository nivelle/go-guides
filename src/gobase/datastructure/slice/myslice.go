package main

import "fmt"

func OperatorSlice() {

	/**
	 切片与数组最大的不同就是切片不是固定长度的
	 */
	var slice1 []int
	if slice1 == nil {
		fmt.Println("切片1 是空")
	} else {
		fmt.Println("切片1 不是空")
	}

	slice2 := []int{1,2,3}
	fmt.Println("切片2:",slice2)
	fmt.Println("切片2 len=", len(slice2))
	fmt.Println("切片2 cap=", cap(slice2))
	/**
	  第三个参数和第二参数len 相等,都是2
	 */
	var slice3 = make([]int,2)
	fmt.Println(slice3)
	fmt.Println("切片3:",slice3)
	fmt.Println("切片3 len=", len(slice3))
	fmt.Println("切片3 cap=", cap(slice3))

    /**
      第二个参数=len
      第三个参数=cap
     */
	var slice4  []int = make([]int,3,4)
	fmt.Println(slice4)
	fmt.Println("切片4:",slice4)
	fmt.Println("切片4 len=", len(slice4))
	fmt.Println("切片4 cap=", cap(slice4))


	slice5 := []int {1,2,3}
	fmt.Println("切片5:",slice5)
	fmt.Println("切片5 len=", len(slice5))
	fmt.Println("切片5 cap=", cap(slice5))

	arr := [5]int{1,2,3,4,5}
	var slice6 []int
	slice6 = arr[1:4]
	fmt.Println("切片6:",slice6)
	fmt.Println("切片6 len=", len(slice6))
	fmt.Println("切片6 cap=", cap(slice6))

	/**
	  初始化的同时 指定 len:cap
	 */
	slice7:= []int {10,20,30,5:30}
	fmt.Println("切片7:",slice7)
	fmt.Println("切片7 len=", len(slice7))
	fmt.Println("切片7 cap=", cap(slice7))

	slice8 := [...]int{9, 8, 7, 6, 5, 4, 3, 2, 1, 0}
	fmt.Println("切片8:",slice8)
	fmt.Println("切片8 len=", len(slice8))
	fmt.Println("切片8 cap=", cap(slice8))

}

func main()  {
	OperatorSlice()
}
