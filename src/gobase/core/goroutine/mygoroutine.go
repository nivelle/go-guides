package main

import "fmt"

func goroutine1() {
	fmt.Println("goroutine1")
}

func goroutine2() {
	fmt.Println("goroutine2")
}

func ShowGoroutine() {
	go goroutine2()
	fmt.Println("after goroutine2")
	go goroutine1()
}

func main()  {
	ShowGoroutine()
}
