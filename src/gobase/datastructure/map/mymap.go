package main

import "fmt"

func MyMap() {
	fmt.Println("==================")

	var m map[int]string
	fmt.Printf("%p \n", m)

	//直接申明初始化 map
	var map1 = map[int]string{123: "nivelle", 124: "jessy"}
	fmt.Print(map1)
	fmt.Print(map1[124])

	fmt.Println()

	//默认长度是0（可以进行读写）
	map2 := make(map[string]string)
	map2["name"] = "nivelle"
	map2["name2"] = "nivelle2"

	fmt.Println(map2)

	for k, v := range map1 {
		fmt.Println(k, v)
	}
	fmt.Println(len(map1))

	value, isExist := map2["name"]
	if isExist {
		fmt.Println(value)
	} else {
		fmt.Println("key 不存在")
	}

	delete(map2, "name")

	fmt.Println("删除name后")
	fmt.Println(map2)
}

func main()  {
	MyMap()
}
