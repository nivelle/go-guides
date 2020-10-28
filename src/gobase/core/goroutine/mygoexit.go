package main

import (
	"fmt"
	"runtime"
)

func main() {

	go func() {
		defer fmt.Println("A.defer")
		fmt.Println("A1")
		func() {
			defer fmt.Println("B.defer")
			// 结束协程
			runtime.Goexit()
			defer fmt.Println("C.defer")
			fmt.Println("B2")
		}()
		fmt.Println("A2")
	}()
	fmt.Println("A3")

}
