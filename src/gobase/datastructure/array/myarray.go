package main

import "fmt"

func showArray() {

	var arr [5]int //元素初始值都会0（int）
	arr[0] = 111
	arr[1] = 222
	fmt.Println(arr)
	fmt.Printf("%T \n 数组:", arr)

	var arr2 = [5]int{11, 22, 33, 44}

	fmt.Println(arr2)

	fmt.Println(len(arr2))
	for i := 0; i < len(arr2); i++ {
		fmt.Println(arr2[i])
	}

	for index, value := range arr2 {
		fmt.Println(index, value)
	}
	fmt.Printf("数组地址默认是第一个元素的位置：%p \n", &arr2)
	fmt.Printf("数组第二个元素位置比第一个大8：%p \n", &arr2[1])

	var arr3 = arr2
	fmt.Printf("数组赋值 是深拷贝:%p \n", &arr3)

	//在数组字面值中，如果在数组的长度位置出现的是“...”省略号，则表示数组的长度是根据初始值的长度确定的
	var arr4 = [...]int{1, 2, 3, 4}
	fmt.Println("数组4:", arr4)
	fmt.Println("数组4长度:", len(arr4))

	/**
	  指定 index 2 , 3 ,4 的值 分别为 1，2，20，其余默认为0 最后一个index是最大索引
	 */
	var arr5 = [5] int{2: 1, 3: 2, 4: 20}
	fmt.Println("数组5:", arr5)
	fmt.Println("数组5长度:", len(arr5))

	var arr6 = [...]int{10: -1}
	fmt.Println("数组6:", arr6)

	//
	fns := [](func(x int) int){
		func(x int) int { return x + 1 },
		func(x int) int { return x + 2 },
	}
	fmt.Println(fns[0], fns[1])

}

func main() {
	showArray()
}
