package main

import "fmt"

//channel是一种类型，一种引用类型
func MyChannelShow(input string) {
	ch1 := make(chan string)
	go func() { ch1 <- "my name is " + input }()
	message := <-ch1
	fmt.Println("message is:	", "`", message, "`")
}

func cacheChannel(input string) {
	// 允许缓存两个值
	message := make(chan string, 2)
	message <- "one:" + input
	message <- "two:" + input
	close(message)
	for chanMessage := range message {
		fmt.Println("通道内的数据：", chanMessage)
	}
	for {
		if data, ok := <-message; ok {
			fmt.Println("通道内的数据,断言遍历:", data)
		}
	}
}

/**
1. channel不像文件一样需要经常去关闭，只有当你确实没有任何发送数据了，或者你想显式的结束range循环之类的，才去关闭channel；
2. 关闭channel后，无法向channel 再发送数据(引发 panic 错误后导致接收立即返回零值)；
3. 关闭channel后，可以继续从channel接收数据；
4. 对于nil channel，无论收发都会被阻塞。
 */
func main() {
	var input string
	fmt.Println("阻塞等待输入。。。")
	fmt.Scanln(&input)
	fmt.Println("channel======")
	MyChannelShow(input)
	/**
	默认通道是 无缓冲 的，这意味着只有在对应的接收（<- chan）通道准备好接收时，才允许进行发送（chan <-）。
	可缓存通道允许在没有对应接收方的情况下，缓存限定数量的值
	 */
	cacheChannel(input)

}
