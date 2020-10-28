package main

import (
	"container/ring"
	"fmt"
)

func main() {
	myRing := ring.New(2)
	fmt.Println(myRing)
	for i := 0; i < 3; i++ {
		myRing.Value = i
		myRing = myRing.Next()
	}
	printRing := func(v interface{}) {
		fmt.Print(v, " ")
	}
	myRing.Do(printRing)
	fmt.Println()
}
