package main

import (
	"sync"
	"math/rand"
	"sync/atomic"
	"runtime"
	"time"
	"fmt"
)

func main() {
	var state = make(map[int]int)
	var mutex = &sync.Mutex{}
	var ops int64 = 0
	for r := 0; r < 100; r++ {
		go func() {
			total := 0
			for {
				key := rand.Intn(5)
				mutex.Lock()
				total += state[key]
				mutex.Unlock()
				atomic.AddInt64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}
	for w := 0; w < 10; w++ {
		go func() {
			for {
				key := rand.Intn(5)
				val := rand.Intn(100)
				//互斥锁
				mutex.Lock()
				state[key] = val
				mutex.Unlock()
				//原子操作
				atomic.AddInt64(&ops, 1)
				runtime.Gosched()
			}
		}()
	}
	time.Sleep(time.Second)
	opsFinal := atomic.LoadInt64(&ops)
	fmt.Println("ops:", opsFinal)
	mutex.Lock()
	fmt.Println("state:", state)
	mutex.Unlock()
}
