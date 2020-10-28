package main

import (
	"fmt"
	"time"
)

//从 jobs通道接收任务，并将结果放到results通道返回
func worker(id int, jobs <-chan int, results chan<- int) {
	for j := range jobs {
		fmt.Println("woker", id, "processing job", j)
		time.Sleep(time.Second)
		results <- j * 2
	}
}

func main() {
	jobs := make(chan int, 10)
	results := make(chan int, 10)

	//启动3个worker,初始是阻塞的，因为通道没值
	for w := 1; w <= 3; w++ {
		go worker(w, jobs, results)
	}

	for j := 1; j <= 9; j++ {
		//将job放到通道
		jobs <- j
	}
	// close 表示这些就是所有的任务了
	close(jobs)
	for a := 1; a <= 9; a++ {
		fmt.Println("result:", <-results)
	}
}
