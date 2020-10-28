package main

import "fmt"

/**
  函数定义基础
 */
func Add(a int, b int) int {
	return a + b
}

func Sub(a int, b int) int {
	return a - b
}

func Values() (int, int) {
	return 3, 7
}

func Sum(nums ...int) int {
	fmt.Println(nums, " ")
	total := 0
	for _, num := range nums {
		total += num
	}
	println("sum result：", total,"；参数个数：",len(nums))

	return total
}

func test(x, y int, s string) (int, string) {
	n := x + y
	return n, fmt.Sprintf(s, n)
}

func main() {
	result1, result2 := Values()
	fmt.Println(result1)
	fmt.Println(result2)
	_, result3 := Values()
	fmt.Println(result3)
	total := Sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	fmt.Println("变参函数求和:", total)
	x, y := test(1, 2, "jessy")
	fmt.Println(x)
	fmt.Println(y)
}
