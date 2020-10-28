package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
	"time"
)

func Logger() gin.HandlerFunc {

	return func(context *gin.Context) {
		t := time.Now()
		context.Set("example", "12345")
		fmt.Println("before===========")
		//请求之前
		context.Next()
		fmt.Println("执行逻辑")
		fmt.Println("after===========")
		//请求之后
		latency := time.Since(t)
		log.Print(latency)
		//发送状态
		status := context.Writer.Status()
		log.Println("记录发送状态:", status)
	}
}

func main() {
	router := gin.New()
	router.Use(Logger())

	router.GET("/test", func(context *gin.Context) {
		example := context.MustGet("example").(string)
		log.Println(example)
	})
	router.Run(":8080")
}
