package main

import "fmt"

func main() {
	jobs := make(chan int, 5)
	done := make(chan bool)

	go func() {
		for {
			j, more := <-jobs
			// 如果jobs已经关闭，并且通道中所有的值已经接收完毕，那么more的值将是false.
			// 当我们完成所有的任务时，将使用这个特性去进行通知
			if more {
				fmt.Println("received job", j)
			} else {
				fmt.Println("receive all jobs")
				done <- true
				return
			}
		}
	}()

	for j := 1; j <= 3; j++ {
		jobs <- j
		fmt.Println("sent job:", j)
	}
	close(jobs)
	fmt.Println("sent all jobs")
	//如果通道为空会阻塞
	<-done
	fmt.Println("sent all jobs after")

}
