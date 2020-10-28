package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup

func hello(i int) {
	//goroutine结束就登记: -1
	defer wg.Done()
	fmt.Println("Hello Goroutine!", i)
}
func main() {
	fmt.Println("waitGroup waiting....")
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go hello(i)
	}
	wg.Wait()
	fmt.Println("waitGroup process....")

}
