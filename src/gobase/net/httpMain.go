package main

import (
	"net/http"
	"fmt"
	"encoding/json"
	"os"
)

func Accept() {
	http.HandleFunc("/go", myHandler)
	http.ListenAndServe("127.0.0.1:8080",nil)
}

func myHandler(response http.ResponseWriter, request *http.Request) {
	fmt.Println(request.RemoteAddr,"连接成功")
	fmt.Println("method:",request.Method)
	fmt.Println("url:",request.URL.Path)
	fmt.Println("header:",request.Header)
	fmt.Println("body:",request.Body)
	enc := json.NewEncoder(os.Stdout)
	d := map[string]int{"aoole": 5, "lettuce": 7}
	enc.Encode(d)
	response.Write([]byte("nivelle is great"))
}

func main()  {
	Accept()
	fmt.Println("服务启动了。。。")
}