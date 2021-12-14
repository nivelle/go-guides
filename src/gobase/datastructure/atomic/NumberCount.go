package main

import (
	"sync/atomic"
	"runtime"
	"time"
	"fmt"
)

func main() {
	var ops uint64 = 0
	for i := 0; i < 1; i++ {
		go func() {
			for {
				atomic.AddUint64(&ops, 1)
				//untime.Gosched()用于让出CPU时间片
				fmt.Print("before ...")
				runtime.Gosched()
				fmt.Print("after ...")

			}
		}()
	}
	time.Sleep(time.Second)
	opsFinal := atomic.LoadUint64(&ops)
	fmt.Print("ops:", opsFinal)
}
