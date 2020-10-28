package main

import (
	"net/rpc"
	"fmt"
	"net/http"
	"gobase/myrpc"
)

func main() {
	var ms = new(myrpc.MathService)
	rpc.Register(ms)
	//将Rpc绑定到HTTP协议上
	rpc.HandleHTTP()
	fmt.Println("启动服务...")
	err := http.ListenAndServe(":1234", nil)
	if err != nil {
		fmt.Println(err.Error())
	}
	fmt.Println("服务已停止!")
}
